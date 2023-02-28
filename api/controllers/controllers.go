package controllers

import (
	"gitlab.com/sb-cloud/player-ms-api/api/controllers/playlist"
	"gitlab.com/sb-cloud/player-ms-api/internal/services"
)

type Controllers struct {
	PlaylistController *playlist.Controller
}

func InitControllers(s *services.Services) *Controllers {
	return &Controllers{
		PlaylistController: playlist.NewPlaylistController(s.PlaylistService),
	}
}
