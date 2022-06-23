package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func LoadDB() error {
	config := GetDataSource()
	connStr := config.URL
	var err error

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
