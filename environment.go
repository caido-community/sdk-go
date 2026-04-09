package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// EnvironmentSDK provides operations on variable environments.
type EnvironmentSDK struct {
	client *Client
}

// List returns all environments.
func (s *EnvironmentSDK) List(
	ctx context.Context,
) (*gen.ListEnvironmentsResponse, error) {
	return gen.ListEnvironments(ctx, s.client.GraphQL)
}

// Get returns a single environment by ID.
func (s *EnvironmentSDK) Get(
	ctx context.Context, id string,
) (*gen.GetEnvironmentResponse, error) {
	return gen.GetEnvironment(ctx, s.client.GraphQL, id)
}

// GetContext returns the global and selected environments.
func (s *EnvironmentSDK) GetContext(
	ctx context.Context,
) (*gen.GetEnvironmentContextResponse, error) {
	return gen.GetEnvironmentContext(ctx, s.client.GraphQL)
}

// Create creates a new environment.
func (s *EnvironmentSDK) Create(
	ctx context.Context, input *gen.CreateEnvironmentInput,
) (*gen.CreateEnvironmentResponse, error) {
	return gen.CreateEnvironment(ctx, s.client.GraphQL, *input)
}

// Select selects an environment (pass nil to deselect).
func (s *EnvironmentSDK) Select(
	ctx context.Context, id *string,
) (*gen.SelectEnvironmentResponse, error) {
	return gen.SelectEnvironment(ctx, s.client.GraphQL, id)
}

// Delete deletes an environment.
func (s *EnvironmentSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteEnvironmentResponse, error) {
	return gen.DeleteEnvironment(ctx, s.client.GraphQL, id)
}
