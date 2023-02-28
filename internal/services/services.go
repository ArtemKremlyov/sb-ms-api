package services

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/repositories"
	"gitlab.com/sb-cloud/player-ms-api/internal/services/playlist"
)

type Services struct {
	PlaylistService *playlist.Service
}

func InitServices(r *repositories.Repositories) *Services {
	return &Services{
		PlaylistService: playlist.NewPlaylistService(r.PlaylistRepository),
	}
}
