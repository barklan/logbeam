package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

const prefix = "LOGBEAM_"

type Config struct {
	Username       string `env:"USER" envDefault:"logbeam"`
	Password       string `env:"PASSWORD" envDefault:"logbeam"`
	RetentionHours uint64 `env:"RETENTION_HOURS" envDefault:"48"`
}

func Read() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg, env.Options{
		Prefix: prefix,
	}); err != nil {
		return nil, fmt.Errorf("failed to parse config from env vars: %w", err)
	}

	return &cfg, nil
}
