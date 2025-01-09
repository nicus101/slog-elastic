package slogelastic

import (
	"context"
	"log/slog"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

// createBaseDocument creates the initial log document with basic record information
// including timestamp, log level, and message
func createBaseDocument(rec slog.Record) map[string]any {
	return map[string]any{
		"time":    rec.Time,
		"level":   rec.Level.String(),
		"message": rec.Message,
	}
}

// buildPrefix creates a dot-separated string from the groups slice,
// adding a trailing dot if groups exist
func buildPrefix(groups []string) string {
	if len(groups) == 0 {
		return ""
	}
	return strings.Join(groups, ".") + "."
}

// collectRecordAttributes extracts all attributes from the slog.Record
// into a slice of slog.Attr
func collectRecordAttributes(rec slog.Record) []slog.Attr {
	var attrs []slog.Attr
	rec.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr)
		return true
	})
	return attrs
}

// collectContextAttributes gathers attributes from all context functions
// and filters out empty keys
func collectContextAttributes(ctx context.Context, contextFuncs []ContextAttrFunc) []slog.Attr {
	var attrs []slog.Attr
	for _, fn := range contextFuncs {
		fnAttrs := fn(ctx)
		for _, attr := range fnAttrs {
			if attr.Key != "" {
				attrs = append(attrs, attr)
			}
		}
	}
	return attrs
}

// addAttributesToDocument adds the given attributes to the document map
// with proper group prefixing and "attribute." prefix
func addAttributesToDocument(document map[string]any, attrs []slog.Attr, prefix string) {
	for _, attr := range attrs {
		key := prefix + "attribute." + attr.Key
		val := attr.Value.Any()
		document[key] = val
	}
}

// indexDocument sends the document to Elasticsearch using the provided index client
// and returns any error that occurs during indexing
func indexDocument(esIndex *index.Index, document map[string]any) error {
	_, err := esIndex.Document(document).Do(context.TODO())
	return err
}
