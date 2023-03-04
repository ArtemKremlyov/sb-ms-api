package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab.com/sb-cloud/player-ms-api/internal/config"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
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

	// Replace `&Product{}, &User{}` with the models of your application.
	if err := db.AutoMigrate(&models.Playlist{}, &models.Song{}); err != nil {
		log.Fatalln(err)
	}
}
