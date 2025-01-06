package slogelastic

import (
	"fmt"
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type Config struct {
	Addresses string `env:"ES_LOG_ADDRESSES"`
	User      string `env:"ES_LOG_USER"`
	Pass      string `env:"ES_LOG_PASS"`
	Index     string `env:"ES_LOG_INDEX"`

	ESIndex *index.Index
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

func (cfg Config) NewElasticHandler() slog.Handler {

	return &Handler{
		esIndex: cfg.ESIndex,
	}
}
