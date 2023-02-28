package repositories

import (
	"github.com/jinzhu/gorm"
	repository "gitlab.com/sb-cloud/player-ms-api/internal/repositories/playlist"
)

type Repositories struct {
	PlaylistRepository *repository.PlaylistRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		PlaylistRepository: repository.NewPlaylistRepository(db),
	}
}
