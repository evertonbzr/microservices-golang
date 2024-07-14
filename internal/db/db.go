package db

import (
	"log"
	"time"

	"github.com/evertonbzr/microservices-golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", "error", err)
	}

	configSQLDriver, err := db.DB()

	if err != nil {
		log.Fatal("Error getting SQL driver from gorm", "error", err)
	}

	configSQLDriver.SetMaxIdleConns(300)
	configSQLDriver.SetMaxOpenConns(380)
	configSQLDriver.SetConnMaxIdleTime(30 * time.Minute)
	configSQLDriver.SetConnMaxLifetime(time.Hour)

	if config.IsDevelopment() {
		db = db.Debug()
	}

	return db
}
