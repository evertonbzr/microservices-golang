package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/microservices-golang/internal/config"
	"github.com/evertonbzr/microservices-golang/internal/db"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	db := db.InitDB(config.POSTGRES_URL)
	slog.Info("Database connected", "url", config.POSTGRES_URL)

}
