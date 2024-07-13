package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/microservices-golang/internal/api"
	"github.com/evertonbzr/microservices-golang/internal/cache"
	"github.com/evertonbzr/microservices-golang/internal/config"
	"github.com/evertonbzr/microservices-golang/internal/db"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	db := db.InitDB(config.POSTGRES_URL)
	slog.Info("Database connected", "env", config.ENV)

	cache := cache.InitRedis(config.REDIS_URL)
	slog.Info("Redis connected", "env", config.ENV)

	apiCfg := &api.APIConfig{
		DB:    db,
		Cache: cache,
		Port:  config.PORT,
	}

	api.Start(apiCfg)
}
