package song

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gitlab.com/sb-cloud/player-ms-api/internal/repositories/song"
)

type Service struct {
	repo *song.SongRepository
}

func NewPlaylistService(repo *song.SongRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllSongs() ([]models.Song, error) {
	return s.repo.GetAll()
}

func (s *Service) GetSongByID(id uint) (*models.Song, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateSong(song *models.Song) error {
	return s.repo.Create(song)
}

func (s *Service) UpdateSong(song *models.Song) error {
	return s.repo.Update(song)
}

func (s *Service) DeleteSong(id uint) error {
	return s.repo.Delete(id)
}
