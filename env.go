package slogelastic

import (
	"errors"
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func (cfg *Config) LoadFromEnv() error {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf(
			"loading .env: %w", err,
		)
	}

	err = env.Parse(cfg)
	if err != nil {
		return fmt.Errorf(
			"parsing enviroment: %w",
			err,
		)
	}

	return nil
}
