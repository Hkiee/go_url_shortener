package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `env:"ENV" env-required:"true"`
	StoragePath string `env:"STORAGE_PATH" env-required:"true"`
}

type HttpServer struct {
	Addr        string        `env:"HTTP_ADDR" env-required:"true"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" env-required:"true"`
}

func MustLoad() *Config {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(".env", cfg); err != nil {
		log.Fatalf("cannot read config from .env: %v", err)
	}

	return cfg
}
