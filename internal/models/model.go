package models

import (
	"gorm.io/gorm"
)

type ImageInfo struct {
	gorm.Model
	Name  string `json:"name"`
	Url   string `json:"url"`
	Image string `json:"-"`
}
