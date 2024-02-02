package config

import (
	"errors"
	"os"
)

type Config struct {
	Token string
}

func LoadConfig() (*Config, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("token is not set")
	}

	return &Config{Token: token}, nil
}
