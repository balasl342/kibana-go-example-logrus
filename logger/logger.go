package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

type ElasticsearchHook struct {
	client *elasticsearch.Client
}

// Create Elastic client
func NewElasticsearchHook(url, apiKey string) (*ElasticsearchHook, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{url},
		APIKey:    apiKey,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchHook{client: client}, nil
}

// Custom Fire method
func (hook *ElasticsearchHook) Fire(entry *logrus.Entry) error {
	data, err := json.Marshal(entry.Data)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "logs",
		DocumentID: fmt.Sprintf("%d", time.Now().UnixNano()),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), hook.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	return nil
}

// Custom Levels support
func (hook *ElasticsearchHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
