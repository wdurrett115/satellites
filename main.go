package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

	return mux
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	srv := &http.Server{
		Addr:         *addr,
		Handler:      routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", *addr)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
