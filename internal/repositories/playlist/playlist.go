package repository

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/gorm"
)

type PlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{db: db}
}

func (r *PlaylistRepository) GetAll() ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := r.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func (r *PlaylistRepository) GetByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	err := r.db.First(&playlist, id).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *PlaylistRepository) AddSong(playlist *models.Playlist, song *models.Song) error {
	return r.db.Model(playlist).Association("Songs").Append(song)
}

func (r *PlaylistRepository) Create(playlist *models.Playlist) error {
	return r.db.Create(playlist).Error
}

func (r *PlaylistRepository) Update(playlist *models.Playlist) error {
	return r.db.Save(playlist).Error
}

func (r *PlaylistRepository) Delete(id uint) error {
	return r.db.Delete(&models.Playlist{ID: id}).Error
}
