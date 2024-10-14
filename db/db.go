package db

import (
	"github.com/sherwin-77/book-crud-golang/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	config := config.GetConfiguration()

	database, err := gorm.Open(sqlite.Open(config.Database.Path), &gorm.Config{
		TranslateError: true,
	})

	DB = database

	return err
}
