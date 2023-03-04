package song

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetAll() ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := r.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func (r *SongRepository) GetByID(id int64) (*models.Playlist, error) {
	var playlist models.Playlist
	err := r.db.First(&playlist, id).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *SongRepository) Create(playlist *models.Playlist) error {
	return r.db.Create(playlist).Error
}

func (r *SongRepository) Update(playlist *models.Playlist) error {
	return r.db.Save(playlist).Error
}

func (r *SongRepository) Delete(id int64) error {
	return r.db.Delete(&models.Playlist{ID: id}).Error
}
