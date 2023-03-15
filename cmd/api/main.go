package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/sb-cloud/player-ms-api/internal/api/grpc"
	"gitlab.com/sb-cloud/player-ms-api/internal/api/rest"
	"gitlab.com/sb-cloud/player-ms-api/internal/music"
	"gitlab.com/sb-cloud/player-ms-api/internal/pg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.com/sb-cloud/player-ms-api/internal/config"
	"gitlab.com/sb-cloud/player-ms-api/internal/models"
)

func main() {
	if err := run(); err != nil && !errors.Is(err, context.Canceled) {
		log.Println(err)
		log.Fatal("stopping")
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, _ := config.New()

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Playlist{}, &models.Song{}); err != nil {
		return err
	}

	p := pg.NewPostgres(db)
	mus := music.NewMusicService(p)
	restApi := rest.NewREST(mus)

	routes := restApi.Register(gin.New())
	srv := &http.Server{
		Addr:    ":" + cfg.ServerHTTPPort,
		Handler: routes,
	}

	grpc := grpc.New(mus)

	errc := make(chan error, 2)

	go start(ctx, errc, 30, func() error {
		return srv.ListenAndServe()
	}, func() error {
		return srv.Shutdown(ctx)
	})

	go start(ctx, errc, 30, func() error {
		return grpc.Start(":" + cfg.GRPCPort)
	}, func() error {
		grpc.Shutdown()
		return nil
	})

	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			stop()
		case err := <-errc:
			stop()
			<-errc
			return err
		}
	}
}

func start(ctx context.Context, errc chan<- error, timeout time.Duration, start, stop func() error) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
		errc <- err
	}()

	ec := make(chan error, 1)
	go func() { ec <- start() }()
	select {
	case err = <-ec:
		return
	case <-ctx.Done():
	}

	go func() { ec <- stop() }()
	timer := time.NewTimer(timeout)
	select {
	case err = <-ec:
		return
	case <-timer.C:
		log.Println(err)
	}
}
