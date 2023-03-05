package repositories

import (
	playlist "gitlab.com/sb-cloud/player-ms-api/internal/repositories/playlist"
	"gitlab.com/sb-cloud/player-ms-api/internal/repositories/song"
	"gorm.io/gorm"
)

type Repositories struct {
	PlaylistRepository *playlist.PlaylistRepository
	SongRepository     *song.SongRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		PlaylistRepository: playlist.NewPlaylistRepository(db),
		SongRepository:     song.NewSongRepository(db),
	}
}
