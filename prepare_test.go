package slogelastic

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

func AlreadyConnected() *index.Index {
	client, _ := elasticsearch.NewTypedClient(elasticsearch.Config{
		Transport: &TestRoundTripper{},
	})
	return client.Index("not-relevant")
}

type TestRoundTripper struct{}

func (t *TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var document map[string]any
	json.NewDecoder(req.Body).Decode(&document)
	document["time"] = time.Time{}
	json.NewEncoder(os.Stdout).Encode(document)

	// Create a new response with status code 200
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}

	return resp, nil
}
