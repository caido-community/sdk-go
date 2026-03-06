package caido

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HealthInfo contains the Caido instance health status.
type HealthInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Ready   bool   `json:"ready"`
}

// Health checks the instance health endpoint.
func (c *Client) Health(ctx context.Context) (*HealthInfo, error) {
	url := c.baseURL + "/health"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, opErr("health", "create request", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, opErr("health", "request failed", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, opErr("health", "read body", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, opErr("health",
			fmt.Sprintf("unexpected status %d: %s", resp.StatusCode, body), nil)
	}

	var info HealthInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, opErr("health", "decode response", err)
	}
	return &info, nil
}

// Ready polls until the instance is ready or the context is cancelled.
func (c *Client) Ready(ctx context.Context, opts ConnectOptions) error {
	opts = opts.withDefaults()

	deadline := time.Now().Add(opts.ReadyTimeout)
	ticker := time.NewTicker(opts.ReadyInterval)
	defer ticker.Stop()

	for {
		info, err := c.Health(ctx)
		if err == nil && info.Ready {
			return nil
		}

		if time.Now().After(deadline) {
			if err != nil {
				return opErr("ready", "timeout waiting for instance", err)
			}
			return &NotReadyError{}
		}

		select {
		case <-ctx.Done():
			return opErr("ready", "context cancelled", ctx.Err())
		case <-ticker.C:
		}
	}
}
