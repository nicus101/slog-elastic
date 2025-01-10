package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/MatusOllah/slogcolor"
	slogelastic "github.com/nicus101/slog-elastic"
	slogmulti "github.com/samber/slog-multi"
)

func initLogs() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var slogEsCfg slogelastic.Config
	if err := slogEsCfg.LoadFromEnv(); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	if err := slogEsCfg.ConnectEsLog(); err != nil {
		log.Fatal("Cannot connect elastic:", err)
	}

	slogEsCfg.ContextFuncs = append(
		slogEsCfg.ContextFuncs,
		func(ctx context.Context) []slog.Attr {
			kot := ctx.Value("kot")
			return []slog.Attr{
				slog.Any("kot", kot),
			}
		})

	slogHandler := slogEsCfg.NewElasticHandler().WithAttrs([]slog.Attr{slog.Bool("slog-sim", true)})

	fanout := slogmulti.Fanout(
		slogHandler,
		slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions),
	)
	slog.SetDefault(slog.New(fanout))
}
