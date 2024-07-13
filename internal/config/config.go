package config

import (
	"log"
	"log/slog"

	"github.com/spf13/viper"
)

var (
	PORT         string
	ENV          string
	POSTGRES_URL string
	REDIS_URL    string
)

func Load(env string) {
	slog.Info("Loading config...", "env", env)

	if env == "" {
		viper.Set("ENV", "development")
		viper.Set("PORT", "3001")
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error loading .env file", "error", err)
		}

	} else {
		viper.AutomaticEnv()
	}

	ENV = viper.GetString("ENV")
	PORT = viper.GetString("PORT")
	POSTGRES_URL = viper.GetString("POSTGRES_URL")
	REDIS_URL = viper.GetString("REDIS_URL")
}

func IsDevelopment() bool {
	return ENV == "development"
}

func IsProduction() bool {
	return ENV == "production"
}

func IsTest() bool {
	return ENV == "test"
}
