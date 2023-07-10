package handler

import (
	"github.com/vinidu4rte/url-shortener-golang/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeHandler() {
	db = config.GetDatabase()
}
