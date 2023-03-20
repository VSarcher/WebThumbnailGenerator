package models

import (
	"fmt"

	"github.com/VSarcher/WebThumbnailGenerator/database"
	"gorm.io/gorm"
)

type ImageInfo struct {
	gorm.Model
	Name  string `json:"name"`
	Url   string `json:"url"`
	Image string `json:"-"`
}

// func NewImageInfo(db *gorm.DB)

func (img ImageInfo) SaveImage() error {
	data := database.DB.Save(img)
	fmt.Println(data)
	return nil
}
