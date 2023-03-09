package models

import (
	"time"
)

type Song struct {
	ID          uint          `gorm:"primarykey"`
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration" gorm:"not null"`
	Playing     bool          `json:"playing" gorm:""`
	Playlists   []Playlist    `json:"playlists" gorm:"many2many:playlist_songs"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
