package slogelastic

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type Handler struct {
	esIndex *index.Index
}

var _ slog.Handler = &Handler{}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	// TODO: implement it
	return true
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return h
}

func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	document := map[string]any{
		"time":    rec.Time,
		"level":   rec.Level.String(),
		"message": rec.Message,
	}
	// TODO: in context could be some values

	for attr := range rec.Attrs {
		key := "attribute." + attr.Key
		val := attr.Value.Any()

		document[key] = val
	}

	_, err := h.esIndex.Document(document).Do(context.TODO())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Err indexing: ", err)
	}

	return nil
}
