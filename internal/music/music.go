package music

import (
	"fmt"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/gorm"
)

type PostgresImpl interface {
	PlaylistGetAll() ([]models.Playlist, error)
	PlaylistGetByID(id uint) (*models.Playlist, error)
	PlaylistAddSong(playlist *models.Playlist, song *models.Song) error
	PlaylistDeleteSong(playlist *models.Playlist, song *models.Song) error
	PlaylistCreate(playlist *models.Playlist) error
	PlaylistUpdate(playlist *models.Playlist) error
	PlaylistDelete(id uint) error

	SongsGetByPlaylistId(playlist *models.Playlist) ([]models.Song, error)
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
	playlists, err := s.p.PlaylistGetAll()
	if err != nil {
		return nil, fmt.Errorf("playlistGetAll error: %v", err)
	}

	return playlists, nil
}

func (s *Service) GetPlaylistByID(id uint) (*models.Playlist, error) {
	playlist, err := s.p.PlaylistGetByID(id)
	if err != nil {
		return nil, fmt.Errorf("getPlaylistById error: %v", err)
	}

	return playlist, nil
}

func (s *Service) GetPlaylistSongs(playlist *models.Playlist) ([]models.Song, error) {
	songs, err := s.p.SongsGetByPlaylistId(playlist)
	if err != nil {
		return nil, fmt.Errorf("getPlaylistSongs error: %v", err)
	}

	return songs, nil
}

func (s *Service) AddSong(playlist *models.Playlist, song *models.Song) error {
	err := s.p.PlaylistAddSong(playlist, song)
	if err != nil {
		return fmt.Errorf("addSong error: %v", err)
	}

	return nil
}

func (s *Service) PlaylistDeleteSong(playlist *models.Playlist, song *models.Song) error {
	err := s.p.PlaylistDeleteSong(playlist, song)
	if err != nil {
		return fmt.Errorf("deleteSong error: %v", err)
	}

	return nil
}

func (s *Service) CreatePlaylist(playlist *models.Playlist) error {
	err := s.p.PlaylistCreate(playlist)
	if err != nil {
		return fmt.Errorf("createPlaylist error: %v", err)
	}

	return nil
}

func (s *Service) UpdatePlaylist(playlist *models.Playlist) error {
	err := s.p.PlaylistUpdate(playlist)
	if err != nil {
		return fmt.Errorf("updatePlaylist error: %v", err)
	}

	return nil
}

func (s *Service) DeletePlaylist(id uint) error {
	err := s.p.PlaylistDelete(id)
	if err != nil {
		return fmt.Errorf("playlistGetAll error: %v", err)
	}

	return nil
}

// Songs
func (s *Service) GetAllSongs() ([]models.Song, error) {
	songs, err := s.p.SongGetAll()
	if err != nil {
		return nil, fmt.Errorf("getAllSongs error: %v", err)
	}

	return songs, nil
}

func (s *Service) GetSongByID(id uint) (*models.Song, error) {
	song, err := s.p.SongGetByID(id)
	if err != nil {
		return nil, fmt.Errorf("getSongById error: %v", err)
	}

	return song, err
}

func (s *Service) GetNextSong(playlist *models.Playlist, songId uint) (*models.Song, error) {
	songs, err := s.p.SongsGetByPlaylistId(playlist)
	if err != nil {
		return nil, fmt.Errorf("songsGetByPlaylistId error: %v", err)
	}

	var currentIndex int
	for i, song := range songs {
		if song.ID == songId {
			currentIndex = i
			break
		}
	}

	if len(songs) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var nextSong models.Song
	if currentIndex+1 < len(songs) {
		nextSong = songs[currentIndex+1]
	} else {
		nextSong = songs[0]
	}

	return &nextSong, err
}

func (s *Service) GetPrevSong(playlist *models.Playlist, songId uint) (*models.Song, error) {
	songs, err := s.p.SongsGetByPlaylistId(playlist)
	if err != nil {
		return nil, fmt.Errorf("songsGetByPlaylistId error: %v", err)
	}

	var currentIndex int
	for i, song := range songs {
		if song.ID == songId {
			currentIndex = i
			break
		}
	}

	if len(songs) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var prevSong models.Song
	if currentIndex-1 >= 0 {
		prevSong = songs[currentIndex-1]
	} else {
		prevSong = songs[len(songs)-1]
	}

	return &prevSong, err
}

func (s *Service) CreateSong(song *models.Song) error {
	err := s.p.SongCreate(song)
	if err != nil {
		return fmt.Errorf("createSong error: %v", err)
	}

	return nil
}

func (s *Service) UpdateSong(song *models.Song) error {
	err := s.p.SongUpdate(song)
	if err != nil {
		return fmt.Errorf("updateSong error: %v", err)
	}

	return nil
}

func (s *Service) DeleteSong(id uint) error {
	err := s.p.SongDelete(id)
	if err != nil {
		return fmt.Errorf("deleteSong: %v", err)
	}

	return nil
}
