package slogelastic_test

import (
	"log"
	"log/slog"

	slogelastic "github.com/nicus101/slog-elastic"
)

func Example() {
	// initialize by config
	slogEsCfg := slogelastic.Config{
		Addresses: "https://example.com",
		Index:     "some-log-index",
		User:      "john",
		Pass:      "secret",
		MinLevel:  slog.LevelDebug,
	}

	// load from .env or enviroment ES_LOG_xxx
	err := slogEsCfg.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// connecting to ElasticSearch and selecting index
	err = slogEsCfg.ConnectEsLog()
	if err != nil {
		log.Fatal(err)
	}

	// or use arleady established connection
	slogEsCfg.ESIndex = slogelastic.AlreadyConnected()

	// finalize configuration and build slog.Handler
	esHandler := slogEsCfg.NewElasticHandler()

	// To see output in terminal, we recomend slogmulti.Fanout from Samber
	slog.SetDefault(slog.New(esHandler))

	// now we can use persistent logging in rest of application
	slog.Info("Registered banana", "bannanaId", 42)
	slog.Warn("Invalid monkeyId", "monkeyId", "mojo")
	slog.Error("BannanaDB connection failed", "error", "unknown protocol banana://")

	// Output: {"bannanaId":42,"level":"INFO","message":"Registered banana","time":"0001-01-01T00:00:00Z"}
	// {"level":"WARN","message":"Invalid monkeyId","monkeyId":"mojo","time":"0001-01-01T00:00:00Z"}
	// {"error":"unknown protocol banana://","level":"ERROR","message":"BannanaDB connection failed","time":"0001-01-01T00:00:00Z"}
}
