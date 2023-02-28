package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresUser     string `envconfig:POSTGRES_USER`
	PostgresPassword string `envconfig:POSTGRES_PASSWORD`
	PostgresPort     string `envconfig:POSTGRES_PORT`
	PostgresHost     string `envconfig:POSTGRES_HOST`
	PostgresDB       string `envconfig:POSTGRES_DB`
	ServerHTTPPort   string `envconfig:"SERVER_HTTP_PORT" default:"8080"`
}

func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var c Config
	err = envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
