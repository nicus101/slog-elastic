package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	slogelastic "github.com/nicus101/slog-elastic"
)

func initLogs() {
	//slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "__secret__",
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal("connect elastic:", err)
	}
	esIndex := es.Index("test-logs-2")

	slog.SetDefault(slog.New(slogelastic.Config{
		ESIndex: esIndex,
	}.NewElasticHandler()))
}
