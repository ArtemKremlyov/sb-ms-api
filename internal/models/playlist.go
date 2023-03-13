package models

import "time"

type Playlist struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name" gorm:"not null" validate:"min=2,max=30,required"`
	Songs     []Song `json:"songs" gorm:"many2many:playlist_songs"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
