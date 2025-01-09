package slogelastic

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type Handler struct {
	esIndex      *index.Index
	minLevel     slog.Level
	contextFuncs []ContextAttrFunc
	groups       []string
	errorHandler func(error)
}

var _ slog.Handler = &Handler{}

// Enabled checks if the given log level meets the minimum level requirement
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

// WithAttrs returns the handler itself, maintaining compatibility with slog.Handler interface
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup creates a new handler with the given group name appended to the groups slice
func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := *h
	h2.groups = append(h2.groups, name)
	return &h2
}

// Handle processes a log record, converting it to a document and sending it to Elasticsearch
func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	document := createBaseDocument(rec)
	prefix := buildPrefix(h.groups)

	recordAttrs := collectRecordAttributes(rec)
	contextAttrs := collectContextAttributes(ctx, h.contextFuncs)

	allAttrs := append(recordAttrs, contextAttrs...)
	addAttributesToDocument(document, allAttrs, prefix)

	if err := indexDocument(h.esIndex, document); err != nil {
		h.errorHandler(err)
	}

	return nil
}

// Add option pattern for configuration
type Option func(*Handler)

// WithErrorHandler returns an Option that sets a custom error handler function
func WithErrorHandler(fn func(error)) Option {
	return func(h *Handler) {
		h.errorHandler = fn
	}
}

// NewElasticHandler creates a new Handler with the given configuration and options
func (cfg Config) NewElasticHandler(opts ...Option) slog.Handler {
	h := &Handler{
		esIndex:      cfg.ESIndex,
		contextFuncs: cfg.ContextFuncs,
		minLevel:     cfg.MinLevel,
		errorHandler: func(err error) {
			fmt.Fprintln(os.Stderr, "Elasticsearch logging error:", err)
		},
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}
