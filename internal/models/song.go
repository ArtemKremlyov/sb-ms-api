package models

import (
	"gorm.io/gorm"
	"time"
)

type Song struct {
	gorm.Model
	ID          uint          `gorm:"primarykey"`
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration" gorm:"not null"`
	Playing     bool          `json:"playing" gorm:""`
	Playlists   []Playlist    `json:"playlists" gorm:"many2many:playlist_songs"`
}
