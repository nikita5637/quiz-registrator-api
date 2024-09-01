package elasticsearch

import (
	"bytes"
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
)

// Logger ...
type Logger struct {
	elasticSearchClient *elasticsearch.Client
	index               string
}

// Config ...
type Config struct {
	ElasticSearchClient *elasticsearch.Client
	Index               string
}

// New ...
func New(cfg Config) *Logger {
	return &Logger{
		elasticSearchClient: cfg.ElasticSearchClient,
		index:               cfg.Index,
	}
}

// Write ...
func (l *Logger) Write(b []byte) (int, error) {
	req := esapi.IndexRequest{
		Index:        l.index,
		DocumentType: "position",
		DocumentID:   uuid.NewString(),
		Body:         bytes.NewBuffer(b),
		Refresh:      "false",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := req.Do(ctx, l.elasticSearchClient)
	if err == nil {
		defer res.Body.Close()
	}

	return len(b), nil
}
