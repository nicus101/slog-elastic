package main

import (
	"log/slog"

	slogelastic "github.com/nicus101/slog-elastic"
)

func initLogs() {
	//slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	slog.SetDefault(slog.New(&slogelastic.Handler{}))
}
