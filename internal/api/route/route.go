package route

import (
	"github.com/evertonbzr/microservices-golang/internal/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RouteParams struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewRoute(
	db *gorm.DB,
	cache *redis.Client,
) *RouteParams {
	return &RouteParams{
		DB:    db,
		Cache: cache,
	}
}

func (r *RouteParams) Init(router *chi.Mux) {
	userHandler := handler.NewUserHandler(r.DB, r.Cache)

	router.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.ListUsers())
		r.Get("/ping", userHandler.Ping())
	})
}
