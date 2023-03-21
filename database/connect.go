package database

import (
	"github.com/VSarcher/WebThumbnailGenerator/internal/models"
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
	db.AutoMigrate(models.ImageInfo{})
	return nil
}
