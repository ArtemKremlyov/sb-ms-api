package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/sb-cloud/player-ms-api/api/controllers"
	"gitlab.com/sb-cloud/player-ms-api/api/router"
	"gitlab.com/sb-cloud/player-ms-api/internal/config"
	repositories "gitlab.com/sb-cloud/player-ms-api/internal/repositories"
	"gitlab.com/sb-cloud/player-ms-api/internal/services"
	"log"
	"net/url"
)

func main() {
	cfg, _ := config.New()

	dsn := url.URL{
		User:     url.UserPassword(cfg.PostgresUser, cfg.PostgresPassword),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", cfg.PostgresHost, cfg.PostgresPort),
		Path:     cfg.PostgresDB,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repos := repositories.InitRepositories(db)
	serv := services.InitServices(repos)
	controlls := controllers.InitControllers(serv)

	router := router.NewRouter(controlls)
	err = router.Run(cfg.ServerHTTPPort)
	if err != nil {
		return
	}
}
