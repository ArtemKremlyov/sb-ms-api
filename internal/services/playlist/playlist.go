package playlist

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
)

//go:generate mockgen -source playlist.go -destination=./mocks/mocks.go
type Repository interface {
	GetAll() ([]models.Playlist, error)
	GetByID(id uint) (*models.Playlist, error)
	AddSong(playlist *models.Playlist, song *models.Song) error
	Create(playlist *models.Playlist) error
	Update(playlist *models.Playlist) error
	Delete(id uint) error
}

type Service struct {
	repo Repository
}

func NewPlaylistService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllPlaylists() ([]models.Playlist, error) {
	return s.repo.GetAll()
}

func (s *Service) GetPlaylistByID(id uint) (*models.Playlist, error) {
	return s.repo.GetByID(id)
}

func (s *Service) AddSong(playlist *models.Playlist, song *models.Song) error {
	return s.repo.AddSong(playlist, song)
}

func (s *Service) CreatePlaylist(playlist *models.Playlist) error {
	return s.repo.Create(playlist)
}

func (s *Service) UpdatePlaylist(playlist *models.Playlist) error {
	return s.repo.Update(playlist)
}

func (s *Service) DeletePlaylist(id uint) error {
	return s.repo.Delete(id)
}
