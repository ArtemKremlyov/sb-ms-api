package main

import (
	"gitlab.com/sb-cloud/player-ms-api/api/controllers"
	"gitlab.com/sb-cloud/player-ms-api/api/router"
	"gitlab.com/sb-cloud/player-ms-api/internal/config"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
	repositories "gitlab.com/sb-cloud/player-ms-api/internal/repositories"
	"gitlab.com/sb-cloud/player-ms-api/internal/services"
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

	if err := db.AutoMigrate(&models.Playlist{}, &models.Song{}); err != nil {
		log.Fatalln(err)
	}

	repos := repositories.InitRepositories(db)
	serv := services.InitServices(repos)
	controlls := controllers.InitControllers(serv)

	router := router.NewRouter(controlls)
	err = router.Run(cfg.ServerHTTPPort)
	if err != nil {
		return
	}
}
