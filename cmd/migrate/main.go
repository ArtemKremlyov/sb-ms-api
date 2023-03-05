package main

import (
	"gitlab.com/sb-cloud/player-ms-api/internal/config"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg, _ := config.New()

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Playlist{}, &models.Song{}, &models.PlaylistSong{}); err != nil {
		log.Fatalln(err)
	}
}
