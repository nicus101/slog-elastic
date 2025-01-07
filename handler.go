package slogelastic

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

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

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := *h
	h2.groups = append(h2.groups, name)
	return &h2
}

func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	document := map[string]any{
		"time":    rec.Time,
		"level":   rec.Level.String(),
		"message": rec.Message,
	}

	// Handle groups when processing attributes
	prefix := ""
	if len(h.groups) > 0 {
		prefix = strings.Join(h.groups, ".") + "."
	}

	var attrs []slog.Attr
	rec.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr)
		return true
	})

	for _, fn := range h.contextFuncs {
		fnAttrs := fn(ctx)
		for _, attr := range fnAttrs {
			if attr.Key != "" {
				attrs = append(attrs, attr)
			}
		}
	}

	for _, attr := range attrs {
		key := prefix + "attribute." + attr.Key
		val := attr.Value.Any()
		document[key] = val
	}

	_, err := h.esIndex.Document(document).Do(context.TODO())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Err indexing: ", err)
	}

	return nil
}

// Add option pattern for configuration
type Option func(*Handler)

func WithErrorHandler(fn func(error)) Option {
	return func(h *Handler) {
		h.errorHandler = fn
	}
}

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
