package caido

import "time"

// Options configures the Caido client.
type Options struct {
	// URL is the base URL of the Caido instance (e.g. "http://localhost:8080").
	URL string

	// Auth configures authentication. Use PATAuth() or TokenAuth().
	Auth AuthOption
}

// AuthOption configures client authentication.
type AuthOption interface {
	apply(*authConfig)
}

type authConfig struct {
	pat          string
	accessToken  string
	refreshToken string
}

type patAuth struct{ token string }

func (p patAuth) apply(c *authConfig) { c.pat = p.token }

// PATAuth returns an AuthOption that authenticates with a Personal Access Token.
// This is the recommended auth method for SDK usage.
func PATAuth(token string) AuthOption { return patAuth{token: token} }

type tokenAuth struct {
	access  string
	refresh string
}

func (t tokenAuth) apply(c *authConfig) {
	c.accessToken = t.access
	c.refreshToken = t.refresh
}

// TokenAuth returns an AuthOption that authenticates with access/refresh tokens.
func TokenAuth(accessToken, refreshToken string) AuthOption {
	return tokenAuth{access: accessToken, refresh: refreshToken}
}

// ConnectOptions configures the Connect behavior.
type ConnectOptions struct {
	// WaitForReady polls the instance until it reports ready.
	// Default: false.
	WaitForReady bool

	// ReadyTimeout is the max time to wait for readiness.
	// Default: 60s.
	ReadyTimeout time.Duration

	// ReadyInterval is the polling interval for readiness checks.
	// Default: 1s.
	ReadyInterval time.Duration
}

func (o *ConnectOptions) withDefaults() ConnectOptions {
	out := *o
	if out.ReadyTimeout == 0 {
		out.ReadyTimeout = 60 * time.Second
	}
	if out.ReadyInterval == 0 {
		out.ReadyInterval = 1 * time.Second
	}
	return out
}
