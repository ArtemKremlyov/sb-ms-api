package pg

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(db *gorm.DB) *Postgres {
	return &Postgres{
		db,
	}
}

func (p *Postgres) PlaylistGetAll() ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := p.db.Find(&playlists).Error
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func (p *Postgres) PlaylistGetByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	err := p.db.First(&playlist, id).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (p *Postgres) PlaylistAddSong(playlist *models.Playlist, song *models.Song) error {
	return p.db.Model(playlist).Association("Songs").Append(song)
}

func (p *Postgres) PlaylistCreate(playlist *models.Playlist) error {
	return p.db.Create(playlist).Error
}

func (p *Postgres) PlaylistUpdate(playlist *models.Playlist) error {
	return p.db.Save(playlist).Error
}

func (p *Postgres) PlaylistDelete(id uint) error {
	return p.db.Delete(&models.Playlist{ID: id}).Error
}

func (r *Postgres) SongGetAll() ([]models.Song, error) {
	var songs []models.Song
	err := r.db.Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *Postgres) SongGetByID(id uint) (*models.Song, error) {
	var song models.Song
	err := r.db.First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *Postgres) SongCreate(song *models.Song) error {
	return r.db.Create(song).Error
}

func (r *Postgres) SongUpdate(song *models.Song) error {
	return r.db.Save(song).Error
}

func (r *Postgres) SongDelete(id uint) error {
	return r.db.Delete(&models.Song{ID: id}).Error
}
