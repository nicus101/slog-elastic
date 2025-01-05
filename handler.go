package slogelastic

import (
	"context"
	"fmt"
	"log/slog"
)

type Handler struct{}

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
	// TODO: in context coud be some values
	// TODO: handle attributes

	fmt.Printf("### %v\n", document)

	return nil
}
