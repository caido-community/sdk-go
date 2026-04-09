package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// ScopeSDK provides operations on target scopes.
type ScopeSDK struct {
	client *Client
}

// List returns all scopes.
func (s *ScopeSDK) List(
	ctx context.Context,
) (*gen.ListScopesResponse, error) {
	return gen.ListScopes(ctx, s.client.GraphQL)
}

// Get returns a single scope by ID.
func (s *ScopeSDK) Get(
	ctx context.Context, id string,
) (*gen.GetScopeResponse, error) {
	return gen.GetScope(ctx, s.client.GraphQL, id)
}

// Create creates a new scope.
func (s *ScopeSDK) Create(
	ctx context.Context, input *gen.CreateScopeInput,
) (*gen.CreateScopeResponse, error) {
	return gen.CreateScope(ctx, s.client.GraphQL, *input)
}

// Rename renames a scope.
func (s *ScopeSDK) Rename(
	ctx context.Context, id, name string,
) (*gen.RenameScopeResponse, error) {
	return gen.RenameScope(ctx, s.client.GraphQL, id, name)
}

// Delete deletes a scope.
func (s *ScopeSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteScopeResponse, error) {
	return gen.DeleteScope(ctx, s.client.GraphQL, id)
}
