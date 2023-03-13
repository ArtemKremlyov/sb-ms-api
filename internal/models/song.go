package models

import (
	"time"
)

type Song struct {
	ID          uint          `gorm:"primarykey"`
	Name        string        `json:"name" gorm:"not null" validate:"min=2,max=30,required"`
	Description string        `json:"description" validate:"min=2,max=512,required"`
	Duration    time.Duration `json:"duration" gorm:"not null" validate:"min=1,max=2560,required"`
	Playing     bool          `json:"playing" gorm:""`
	Playlists   []Playlist    `json:"playlists" gorm:"many2many:playlist_songs"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
