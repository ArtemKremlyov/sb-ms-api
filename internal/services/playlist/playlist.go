package playlist

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	repository "gitlab.com/sb-cloud/player-ms-api/internal/repositories/playlist"
)

type Service struct {
	repo *repository.PlaylistRepository
}

func NewPlaylistService(repo *repository.PlaylistRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllPlaylists() ([]models.Playlist, error) {
	return s.repo.GetAll()
}

func (s *Service) GetPlaylistByID(id int64) (*models.Playlist, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreatePlaylist(playlist *models.Playlist) error {
	return s.repo.Create(playlist)
}

func (s *Service) UpdatePlaylist(playlist *models.Playlist) error {
	return s.repo.Update(playlist)
}

func (s *Service) DeletePlaylist(id int64) error {
	return s.repo.Delete(id)
}
