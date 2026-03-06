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

	gql "github.com/Khan/genqlient/graphql"
)

// Client is the main Caido SDK client.
// It exposes domain-specific SDKs as fields, mirroring the JS SDK pattern.
type Client struct {
	// Domain SDKs
	Requests    *RequestSDK
	Replay      *ReplaySDK
	Findings    *FindingSDK
	Scopes      *ScopeSDK
	Projects    *ProjectSDK
	Environments *EnvironmentSDK
	HostedFiles *HostedFileSDK
	Workflows   *WorkflowSDK
	Tasks       *TaskSDK
	Instance    *InstanceSDK
	Filters     *FilterSDK
	Users       *UserSDK
	Plugins     *PluginSDK
	Automate    *AutomateSDK
	Sitemap     *SitemapSDK
	Intercept   *InterceptSDK

	// Low-level access
	GraphQL gql.Client

	baseURL    string
	httpClient *http.Client
	authCfg    authConfig
}

// NewClient creates a new Caido client with the given options.
func NewClient(opts Options) (*Client, error) {
	if opts.URL == "" {
		return nil, fmt.Errorf("caido: URL is required")
	}

	var auth authConfig
	if opts.Auth != nil {
		opts.Auth.apply(&auth)
	}

	httpClient := &http.Client{
		Transport: &authTransport{
			base: http.DefaultTransport,
			auth: &auth,
		},
	}

	gqlClient := gql.NewClient(
		opts.URL+"/graphql",
		httpClient,
	)

	c := &Client{
		baseURL:    opts.URL,
		httpClient: httpClient,
		authCfg:    auth,
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

	return c, nil
}

// Connect verifies connectivity and authentication with the Caido instance.
// This should be called before making any API requests.
func (c *Client) Connect(ctx context.Context) error {
	return c.ConnectWithOptions(ctx, ConnectOptions{})
}

// ConnectWithOptions connects with custom readiness options.
func (c *Client) ConnectWithOptions(ctx context.Context, opts ConnectOptions) error {
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
	auth *authConfig
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.auth.pat != "" {
		req.Header.Set("Authorization", "Bearer "+t.auth.pat)
	} else if t.auth.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+t.auth.accessToken)
	}
	return t.base.RoundTrip(req)
}
