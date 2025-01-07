package slogelastic

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/joho/godotenv"
)

type ContextAttrFunc func(context.Context) []slog.Attr

type Config struct {
	Addresses string `env:"ES_LOG_ADDRESSES"`
	User      string `env:"ES_LOG_USER"`
	Pass      string `env:"ES_LOG_PASS"`
	Index     string `env:"ES_LOG_INDEX"`

	ESIndex      *index.Index
	MinLevel     slog.Level
	ContextFuncs []ContextAttrFunc
}

func (cfg *Config) ConnectEsLog() error {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.Addresses,
		},
		Username: cfg.User,
		Password: cfg.Pass,
	}

	es, err := elasticsearch.NewTypedClient(esCfg)
	if err != nil {
		return fmt.Errorf(
			"connecting to elasticsearch: %w", err,
		)
	}

	cfg.ESIndex = es.Index(cfg.Index)
	return nil
}

func (cfg *Config) LoadFromEnv() error {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("loading .env: %w", err)
	}

	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("parsing environment: %w", err)
	}

	// Add validation
	if cfg.Addresses == "" {
		return fmt.Errorf("ES_LOG_ADDRESSES is required")
	}
	if cfg.Index == "" {
		return fmt.Errorf("ES_LOG_INDEX is required")
	}

	return nil
}
