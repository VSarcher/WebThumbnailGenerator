package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/VSarcher/WebThumbnailGenerator/internal/models"
)

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open("image.db"), &gorm.Config{})
	if err != nil {
		// panic("failed to connect database")
		return err
	}

	db.AutoMigrate(&models.ImageInfo{})
	//initial model
	db.Create(&models.ImageInfo{Url: "google.com", Name: "Google"})
	return nil
}
