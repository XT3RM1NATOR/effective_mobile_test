package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type EnvMode string

const (
	Dev  EnvMode = "dev"
	Prod EnvMode = "prod"
)

type Config struct {
	Server   Server
	Postgres Postgres
	API      API
}

type (
	Postgres struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
	}

	Server struct {
		Port string `env:"SERVER_PORT"`
	}

	API struct {
		EnrichAPI string `env:"ENRICH_API_URL"`
	}
)

var instance Config

func Load() *Config {
	if err := godotenv.Load("../../.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	if err := env.Parse(&instance); err != nil {
		panic(err)
	}
	return &instance
}
