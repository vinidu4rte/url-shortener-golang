package config

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	database, err := InitializeDatabase()

	if err != nil {
		fmt.Printf("ERROR! Database connection: %v", err)
		return err
	}

	db = database
	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
