// Package caido provides a Go client for the Caido Web Security Platform.
//
// It mirrors the API surface of the official JavaScript SDK (@caido/sdk-client)
// and uses genqlient for type-safe GraphQL code generation.
//
// Basic usage:
//
//	client, err := caido.NewClient(caido.Options{
//	    URL:  "http://localhost:8080",
//	    Auth: caido.PATAuth("your-pat-token"),
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if err := client.Connect(ctx); err != nil {
//	    log.Fatal(err)
//	}
//
//	requests, err := client.Requests.List(ctx, nil)
package caido

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	gql "github.com/Khan/genqlient/graphql"
)

// TokenRefreshFunc is called before each request when set.
// It should return a valid access token or an error.
type TokenRefreshFunc func(ctx context.Context) (string, error)

// Client is the main Caido SDK client.
// It exposes domain-specific SDKs as fields, mirroring the JS SDK pattern.
type Client struct {
	// Domain SDKs
	Requests     *RequestSDK
	Replay       *ReplaySDK
	Findings     *FindingSDK
	Scopes       *ScopeSDK
	Projects     *ProjectSDK
	Environments *EnvironmentSDK
	HostedFiles  *HostedFileSDK
	Workflows    *WorkflowSDK
	Tasks        *TaskSDK
	Instance     *InstanceSDK
	Filters      *FilterSDK
	Users        *UserSDK
	Plugins      *PluginSDK
	Automate     *AutomateSDK
	Sitemap      *SitemapSDK
	Intercept    *InterceptSDK
	Tamper       *TamperSDK
	Auth         *AuthSDK

	// Low-level access
	GraphQL gql.Client

	baseURL    string
	httpClient *http.Client
	auth       *authState
}

// authState holds mutable auth state, protected by a mutex.
type authState struct {
	mu           sync.RWMutex
	pat          string
	accessToken  string
	refreshToken string
	refreshFn    TokenRefreshFunc
}

// NewClient creates a new Caido client with the given options.
func NewClient(opts Options) (*Client, error) {
	if opts.URL == "" {
		return nil, fmt.Errorf("caido: URL is required")
	}

	state := &authState{}
	if opts.Auth != nil {
		var cfg authConfig
		opts.Auth.apply(&cfg)
		state.pat = cfg.pat
		state.accessToken = cfg.accessToken
		state.refreshToken = cfg.refreshToken
	}

	httpClient := &http.Client{
		Transport: &authTransport{
			base: http.DefaultTransport,
			auth: state,
		},
	}

	gqlClient := gql.NewClient(
		opts.URL+"/graphql",
		httpClient,
	)

	c := &Client{
		baseURL:    opts.URL,
		httpClient: httpClient,
		auth:       state,
		GraphQL:    gqlClient,
	}

	// Initialize domain SDKs
	c.Requests = &RequestSDK{client: c}
	c.Replay = &ReplaySDK{client: c}
	c.Findings = &FindingSDK{client: c}
	c.Scopes = &ScopeSDK{client: c}
	c.Projects = &ProjectSDK{client: c}
	c.Environments = &EnvironmentSDK{client: c}
	c.HostedFiles = &HostedFileSDK{client: c}
	c.Workflows = &WorkflowSDK{client: c}
	c.Tasks = &TaskSDK{client: c}
	c.Instance = &InstanceSDK{client: c}
	c.Filters = &FilterSDK{client: c}
	c.Users = &UserSDK{client: c}
	c.Plugins = &PluginSDK{client: c}
	c.Automate = &AutomateSDK{client: c}
	c.Sitemap = &SitemapSDK{client: c}
	c.Intercept = &InterceptSDK{client: c}
	c.Tamper = &TamperSDK{client: c}
	c.Auth = &AuthSDK{client: c}

	return c, nil
}

// SetAccessToken updates the access token used for authentication.
// This is safe for concurrent use.
func (c *Client) SetAccessToken(token string) {
	c.auth.mu.Lock()
	c.auth.accessToken = token
	c.auth.mu.Unlock()
}

// SetRefreshToken updates the refresh token.
// This is safe for concurrent use.
func (c *Client) SetRefreshToken(token string) {
	c.auth.mu.Lock()
	c.auth.refreshToken = token
	c.auth.mu.Unlock()
}

// RefreshToken returns the current refresh token.
func (c *Client) RefreshToken() string {
	c.auth.mu.RLock()
	defer c.auth.mu.RUnlock()
	return c.auth.refreshToken
}

// SetTokenRefresher sets a callback that is invoked before each request
// to obtain a fresh access token. Pass nil to clear.
func (c *Client) SetTokenRefresher(fn TokenRefreshFunc) {
	c.auth.mu.Lock()
	c.auth.refreshFn = fn
	c.auth.mu.Unlock()
}

// BaseURL returns the base URL of the Caido instance.
func (c *Client) BaseURL() string { return c.baseURL }

// WebSocketEndpoint returns the WebSocket URL for GraphQL subscriptions.
// It converts http/https schemes to ws/wss as required by WebSocket libraries.
func (c *Client) WebSocketEndpoint() string {
	base := c.baseURL
	switch {
	case strings.HasPrefix(base, "https://"):
		base = "wss://" + base[8:]
	case strings.HasPrefix(base, "http://"):
		base = "ws://" + base[7:]
	}
	return base + "/ws/graphql"
}

// Connect verifies connectivity and authentication with the Caido instance.
// This should be called before making any API requests.
func (c *Client) Connect(ctx context.Context) error {
	return c.ConnectWithOptions(ctx, ConnectOptions{})
}

// ConnectWithOptions connects with custom readiness options.
func (c *Client) ConnectWithOptions(
	ctx context.Context, opts ConnectOptions,
) error {
	if opts.WaitForReady {
		if err := c.Ready(ctx, opts); err != nil {
			return err
		}
	}

	info, err := c.Health(ctx)
	if err != nil {
		return opErr("connect", "health check failed", err)
	}
	if !info.Ready {
		return &NotReadyError{}
	}

	return nil
}

// authTransport injects auth headers into HTTP requests.
type authTransport struct {
	base http.RoundTripper
	auth *authState
}

func (t *authTransport) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	// If a refresh function is set, call it first.
	t.auth.mu.RLock()
	refreshFn := t.auth.refreshFn
	t.auth.mu.RUnlock()

	if refreshFn != nil {
		token, err := refreshFn(req.Context())
		if err != nil {
			return nil, fmt.Errorf("caido: token refresh: %w", err)
		}
		t.auth.mu.Lock()
		t.auth.accessToken = token
		t.auth.mu.Unlock()
	}

	t.auth.mu.RLock()
	pat := t.auth.pat
	accessToken := t.auth.accessToken
	t.auth.mu.RUnlock()

	if pat != "" {
		req.Header.Set("Authorization", "Bearer "+pat)
	} else if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	return t.base.RoundTrip(req)
}
