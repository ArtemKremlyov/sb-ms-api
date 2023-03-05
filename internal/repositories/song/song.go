package song

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetAll() ([]models.Song, error) {
	var songs []models.Song
	err := r.db.Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongRepository) GetByID(id uint) (*models.Song, error) {
	var song models.Song
	err := r.db.First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *SongRepository) Create(song *models.Song) error {
	return r.db.Create(song).Error
}

func (r *SongRepository) Update(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *SongRepository) Delete(id uint) error {
	return r.db.Delete(&models.Song{ID: id}).Error
}
