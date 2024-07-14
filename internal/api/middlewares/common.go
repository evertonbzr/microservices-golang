package middleware

import (
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
)

func CommonMiddleware(r *chi.Mux) {
	r.Use(m.RequestID)
	r.Use(m.RealIP)
	r.Use(m.Heartbeat("/healthz"))
	r.Use(m.Logger) // <--<< Logger should come before Recoverer
	r.Use(m.Recoverer)
}
