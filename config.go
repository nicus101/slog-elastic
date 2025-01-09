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

// ConnectEsLog establishes a connection to Elasticsearch using the configured credentials
// and initializes the ESIndex client for the specified index. It returns an error if
// the connection cannot be established.
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

// LoadFromEnv loads configuration from environment variables, optionally reading from a .env file if present.
// It validates required fields and returns an error if any required field is missing or if there are issues
// loading the environment variables. The function will not return an error if the .env file is missing,
// but will return errors for other file-related issues.
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
