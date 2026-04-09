package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// FilterSDK provides operations on saved filter presets.
type FilterSDK struct {
	client *Client
}

// List returns all filter presets.
func (s *FilterSDK) List(
	ctx context.Context,
) (*gen.ListFilterPresetsResponse, error) {
	return gen.ListFilterPresets(ctx, s.client.GraphQL)
}

// Get returns a single filter preset by ID.
func (s *FilterSDK) Get(
	ctx context.Context, id string,
) (*gen.GetFilterPresetResponse, error) {
	return gen.GetFilterPreset(ctx, s.client.GraphQL, id)
}

// Create creates a new filter preset.
func (s *FilterSDK) Create(
	ctx context.Context, input *gen.CreateFilterPresetInput,
) (*gen.CreateFilterPresetResponse, error) {
	return gen.CreateFilterPreset(ctx, s.client.GraphQL, *input)
}

// Delete deletes a filter preset.
func (s *FilterSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteFilterPresetResponse, error) {
	return gen.DeleteFilterPreset(ctx, s.client.GraphQL, id)
}
