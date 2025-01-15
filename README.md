![Go Report Card](https://goreportcard.com/badge/nicus101/slog-elastic)
[![GoDoc](https://godoc.org/github.com/nicus101/slog-elastic?status.svg)](https://pkg.go.dev/github.com/nicus101/slog-elastic)
[![Coverage](https://img.shields.io/codecov/c/github/nicus101/slog-elastic)](https://codecov.io/gh/nicus101/slog-elastic)
[![License](https://img.shields.io/github/license/nicus101/slog-elastic)](./LICENSE)

# Slog-Elastic - implementation of slog.Handler for elasticsearch

Golang have structural logging for some time,
named [slog](https://pkg.go.dev/log/slog).
It even has rich and
[Awesome Slog](https://github.com/go-slog/awesome-slog)
ecosystem.
This package is very young, your feedback is welcome.

## Why?

Why use *ElasticSearch* directly?
There is [Logstash](https://www.elastic.co/logstash) after all.
But sometimes You...

 - Develop locally, and still want persistency.
 - Have IoT solution, and don't want big guns.
 - Didn't get access to logstash from Your operations department.

This package was made to solve the above.
Having somewhat decent implementation would be better than
cobbling ad-hoc solutions in every project.

## Features

 - Implements slog.Handler
 - Extracts values from *Context* using ContextAttrFunc
 - Utility function to connect database
 - Utility function to load `ES_LOG_xx` from .env or environment.
 - Works with [zinc](https://zincsearch.com/)
 - Works with [elasticsearch](https://www.elastic.co/elasticsearch/)
 - Ability  to overwrite error handler

## Usage

```go
// initialize by config
slogEsCfg := slogelastic.Config{
	Address:  "https://example.com",
	Index:    "some-log-index",
	User:     "john",
	Pass:     "secret",
	MinLevel: slog.LevelDebug,
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
```

## Planned features and to-do

 - Ability for client code to filter/rename attributes.
 - Safeguard to propagate all logs before application shutdown.
 - Bulk inserts asynchronously.
