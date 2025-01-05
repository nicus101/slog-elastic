package slogelastic

import (
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type Config struct {
	ESIndex *index.Index
}

func (cfg Config) NewElasticHandler() slog.Handler {

	return &Handler{
		esIndex: cfg.ESIndex,
	}
}
