package victoriametrics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	InstantQuery(ctx context.Context, query string, ts time.Time) ([]QueryResult, error)
	RangeQuery(ctx context.Context, query string, start, end time.Time, step time.Duration) ([]QueryResult, error)
	WriteMetrics(ctx context.Context, metrics string) error
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) Client {
	if baseURL == "" {
		baseURL = "http://localhost:8428"
	}
	return &client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// Полная структура ответа VictoriaMetrics
type apiResponse struct {
	Status    string       `json:"status"`
	Data      responseData `json:"data"`
	Error     string       `json:"error,omitempty"`
	ErrorType string       `json:"errorType,omitempty"`
}

type responseData struct {
	ResultType string          `json:"resultType"`
	Result     json.RawMessage `json:"result"`
}

type QueryResult struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value,omitempty"`  // Для instant query
	Values [][]interface{}   `json:"values,omitempty"` // Для range query
}

func (c *client) InstantQuery(ctx context.Context, query string, ts time.Time) ([]QueryResult, error) {
	params := url.Values{}
	params.Add("query", query)
	if !ts.IsZero() {
		params.Add("time", ts.Format(time.RFC3339))
	}

	u := fmt.Sprintf("%s/api/v1/query?%s", c.baseURL, params.Encode())
	return c.executeQuery(ctx, u)
}

func (c *client) RangeQuery(ctx context.Context, query string, start, end time.Time, step time.Duration) ([]QueryResult, error) {
	params := url.Values{}
	params.Add("query", query)
	params.Add("start", start.Format(time.RFC3339))
	params.Add("end", end.Format(time.RFC3339))
	params.Add("step", step.String())

	u := fmt.Sprintf("%s/api/v1/query_range?%s", c.baseURL, params.Encode())
	return c.executeQuery(ctx, u)
}

func (c *client) executeQuery(ctx context.Context, url string) ([]QueryResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	if apiResp.Status == "error" {
		return nil, fmt.Errorf("api error: %s (type: %s)", apiResp.Error, apiResp.ErrorType)
	}

	var results []QueryResult
	if err := json.Unmarshal(apiResp.Data.Result, &results); err != nil {
		return nil, fmt.Errorf("unmarshal results failed: %w", err)
	}

	return results, nil
}

func (c *client) WriteMetrics(ctx context.Context, metrics string) error {
	u := fmt.Sprintf("%s/api/v1/import/prometheus", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", u, bytes.NewBufferString(metrics))
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
