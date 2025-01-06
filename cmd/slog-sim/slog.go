package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"

	slogelastic "github.com/nicus101/slog-elastic"
)

func initLogs() {
	//slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var slogEsCfg slogelastic.Config
	if err := slogEsCfg.LoadFromEnv(); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	if err := slogEsCfg.ConnectEsLog(); err != nil {
		log.Fatal("Cannot connect elastic:", err)
	}

	slog.SetDefault(slog.New(slogEsCfg.NewElasticHandler()))
}
