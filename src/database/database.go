package config

import (
	"log"

	"github.com/felipemaxplay/first-go-api/src/config"
	"github.com/felipemaxplay/first-go-api/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB() *gorm.DB {
	config := config.GetDataSource()
	connStr := config.URL

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Player{})

	return db
}

func CloseDB(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	dbSql.Close()
}
