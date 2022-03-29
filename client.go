package stanfordcorenlp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Client interface {
	Do(ctx context.Context, text string, annotators AnnotatorType) ([]byte, error)
}

type client struct {
	url        string
	httpClient *http.Client
}

var _ Client = (*client)(nil)

// NewClient is a factory method to create a client for Stanford CoreNLP server.
func NewClient(ctx context.Context, url string) *client {
	return &client{
		httpClient: &http.Client{},
		url:        url,
	}
}

// A properties provide specifications for what annotators to run and how to customize the annotators.
type properties struct {
	Annotators   AnnotatorType `json:"annotators,omitempty"`
	OutputFormat string        `json:"outputFormat,omitempty"`
}

// Do sends HTTP request to Stanford CoreNLP API and returns response.
func (c *client) Do(ctx context.Context, text string, annotators AnnotatorType) ([]byte, error) {
	p, err := json.Marshal(properties{
		Annotators:   annotators,
		OutputFormat: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal property: %w", err)
	}

	q, err := query.Values(&struct {
		Properties string `url:"properties"`
	}{
		Properties: string(p),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encode query string: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBufferString(text))
	if err != nil {
		return nil, fmt.Errorf("failed to setup http request: %w", err)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non successfule status code: %d", resp.StatusCode)
	}

	return respBody, nil
}
