package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB_NAME 	string `env:"DB_NAME"`
	DB_HOST 	string `env:"DB_HOST"`
	DB_USER 	string `env:"DB_USER"`
	DB_PORT 	string `env:"DB_PORT"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s:%s?sslmode=disable", c.DB_USER,c.DB_PASSWORD,c.DB_HOST,c.DB_PORT,c.DB_NAME)
}

func loadEnv()(*Config, error) {
	cfg, err := env.ParseAs[Config]();
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %w", err)
	}
	return &cfg, nil
}
