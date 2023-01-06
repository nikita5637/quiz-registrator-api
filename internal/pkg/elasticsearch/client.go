package elasticsearch

import (
	"bytes"
	"context"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/google/uuid"
)

// Client ...
type Client struct {
	elasticAddress      string
	elasticIndex        string
	elasticSearchClient *elasticsearch.Client
}

// Config ...
type Config struct {
	ElasticAddress string
	ElasticIndex   string
}

// New ...
func New(cfg Config) (*Client, error) {
	elasticSearchClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ElasticAddress},
	})
	if err != nil {
		return nil, err
	}

	if _, err := elasticSearchClient.Ping(); err != nil {
		return nil, err
	}

	return &Client{
		elasticAddress:      cfg.ElasticAddress,
		elasticSearchClient: elasticSearchClient,
		elasticIndex:        cfg.ElasticIndex,
	}, nil
}

// Write ...
func (c *Client) Write(b []byte) (int, error) {
	req := esapi.IndexRequest{
		Index:        c.elasticIndex,
		DocumentType: "position",
		DocumentID:   uuid.NewString(),
		Body:         bytes.NewBuffer(b),
		Refresh:      "false",
	}

	res, err := req.Do(context.Background(), c.elasticSearchClient)
	if err == nil {
		defer res.Body.Close()
	}

	return len(b), nil
}
