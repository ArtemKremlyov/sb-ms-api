package music

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
)

type PostgresImpl interface {
	PlaylistGetAll() ([]models.Playlist, error)
	PlaylistGetByID(id uint) (*models.Playlist, error)
	PlaylistAddSong(playlist *models.Playlist, song *models.Song) error
	PlaylistCreate(playlist *models.Playlist) error
	PlaylistUpdate(playlist *models.Playlist) error
	PlaylistDelete(id uint) error

	SongGetAll() ([]models.Song, error)
	SongGetByID(id uint) (*models.Song, error)
	SongCreate(song *models.Song) error
	SongUpdate(song *models.Song) error
	SongDelete(id uint) error
}

type Service struct {
	p PostgresImpl
}

func NewMusicService(pg PostgresImpl) *Service {
	return &Service{
		p: pg,
	}
}

// Playlists
func (s *Service) GetAllPlaylists() ([]models.Playlist, error) {
	return s.p.PlaylistGetAll()
}

func (s *Service) GetPlaylistByID(id uint) (*models.Playlist, error) {
	return s.p.PlaylistGetByID(id)
}

func (s *Service) AddSong(playlist *models.Playlist, song *models.Song) error {
	return s.p.PlaylistAddSong(playlist, song)
}

func (s *Service) CreatePlaylist(playlist *models.Playlist) error {
	return s.p.PlaylistCreate(playlist)
}

func (s *Service) UpdatePlaylist(playlist *models.Playlist) error {
	return s.p.PlaylistUpdate(playlist)
}

func (s *Service) DeletePlaylist(id uint) error {
	return s.p.PlaylistDelete(id)
}

// Songs
func (s *Service) GetAllSongs() ([]models.Song, error) {
	return s.p.SongGetAll()
}

func (s *Service) GetSongByID(id uint) (*models.Song, error) {
	return s.p.SongGetByID(id)
}

func (s *Service) CreateSong(song *models.Song) error {
	return s.p.SongCreate(song)
}

func (s *Service) UpdateSong(song *models.Song) error {
	return s.p.SongUpdate(song)
}

func (s *Service) DeleteSong(id uint) error {
	return s.p.SongDelete(id)
}
