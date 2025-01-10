package slogelastic

import (
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddAttributesToDocument(t *testing.T) {
	tests := []struct {
		name   string
		attrs  []slog.Attr
		prefix string
		want   map[string]any
	}{
		{
			name: "basic attributes",
			attrs: []slog.Attr{
				slog.String("user", "john"),
				slog.Int("age", 25),
			},
			prefix: "",
			want: map[string]any{
				"user": "john",
				"age":  int64(25),
			},
		},
		{
			name: "with group prefix",
			attrs: []slog.Attr{
				slog.String("user", "john"),
			},
			prefix: "auth.",
			want: map[string]any{
				"auth.user": "john",
			},
		},
		{
			name: "with various types",
			attrs: []slog.Attr{
				slog.String("str", "value"),
				slog.String("int64", "42"),
				slog.Float64("float64", 3.14),
				slog.Bool("bool", true),
				slog.Time("time", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			prefix: "",
			want: map[string]any{
				"str":     "value",
				"int64":   "42",
				"float64": 3.14,
				"bool":    true,
				"time":    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}, {
			name: "with group",
			attrs: []slog.Attr{
				slog.Group("kotki", "lizi", "szyrklet", "wader", "czarny"),
			},
			prefix: "",
			want: map[string]any{
				"kotki.lizi":  "szyrklet",
				"kotki.wader": "czarny",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := make(map[string]any)
			addAttributesToDocument(document, tt.attrs, tt.prefix)
			assert.Equal(t, tt.want, document)
		})
	}
}
