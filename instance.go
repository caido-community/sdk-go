package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// InstanceSDK provides operations on instance runtime and settings.
type InstanceSDK struct {
	client *Client
}

// GetRuntime returns instance runtime info (version, platform).
func (s *InstanceSDK) GetRuntime(
	ctx context.Context,
) (*gen.GetRuntimeResponse, error) {
	return gen.GetRuntime(ctx, s.client.GraphQL)
}

// GetSettings returns instance settings.
func (s *InstanceSDK) GetSettings(
	ctx context.Context,
) (*gen.GetInstanceSettingsResponse, error) {
	return gen.GetInstanceSettings(ctx, s.client.GraphQL)
}

// SetSettings updates instance settings.
func (s *InstanceSDK) SetSettings(
	ctx context.Context, input *gen.SetInstanceSettingsInput,
) (*gen.SetInstanceSettingsResponse, error) {
	return gen.SetInstanceSettings(ctx, s.client.GraphQL, *input)
}
