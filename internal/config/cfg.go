package config

import "github.com/caarlos0/env/v6"

type Config struct {
	PostgresDSN    string `env:"POSTGRES_DSN,notEmpty"`
	GRPCPort       string `env:"GRPCPort" envDefault:"9090"`
	ServerHTTPPort string `env:"SERVER_HTTP_PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	cfg := &Config{}

	// Заполним структуру env переменными.
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
