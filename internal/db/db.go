package db

import (
	"log"

	"github.com/evertonbzr/microservices-golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", "error", err)
	}

	if config.IsDevelopment() {
		db = db.Debug()
	}

	return db
}
