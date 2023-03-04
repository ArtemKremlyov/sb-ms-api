package services

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/repositories"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/playlist"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/song"
)

type Services struct {
	PlaylistService *playlist.Service
	SongService     *song.Service
}

func InitServices(r *repositories.Repositories) *Services {
	return &Services{
		PlaylistService: playlist.NewPlaylistService(r.PlaylistRepository),
		SongService:     song.NewPlaylistService(r.SongRepository),
	}
}
