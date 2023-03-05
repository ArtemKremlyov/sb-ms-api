package models

import "gorm.io/gorm"

type PlaylistSong struct {
	gorm.Model
	ID         uint `gorm:"primarykey"`
	PlaylistID uint
	SongID     uint
}
