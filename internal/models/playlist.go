package models

import (
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name" gorm:"not null"`
	Songs []Song `json:"songs" gorm:"many2many:playlist_songs"`
}
