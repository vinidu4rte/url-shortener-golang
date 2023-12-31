package config

import (
	"github.com/vinidu4rte/url-shortener-golang/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=urlshortener port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Redirect{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
