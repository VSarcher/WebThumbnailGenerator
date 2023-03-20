package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("image.db"), &gorm.Config{})

	DB = db
	if err != nil {
		return err
	}

	return nil
}
