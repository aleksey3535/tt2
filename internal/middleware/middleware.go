package middleware

import (
	"log/slog"
	"net/http"
)

type Middleware struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) UseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		m.log.Info("Request received", "method", r.Method, "url", r.URL.String())
		next.ServeHTTP(w, r)
	})
}