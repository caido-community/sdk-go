package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// PluginSDK provides operations on installed plugins.
type PluginSDK struct {
	client *Client
}

// ListPackages returns all installed plugin packages.
func (s *PluginSDK) ListPackages(
	ctx context.Context,
) (*gen.ListPluginPackagesResponse, error) {
	return gen.ListPluginPackages(ctx, s.client.GraphQL)
}

// InstallPackage installs a plugin package.
func (s *PluginSDK) InstallPackage(
	ctx context.Context, input *gen.InstallPluginPackageInput,
) (*gen.InstallPluginPackageResponse, error) {
	return gen.InstallPluginPackage(ctx, s.client.GraphQL, *input)
}

// DeleteUpstreamPlugin deletes an upstream plugin.
func (s *PluginSDK) DeleteUpstreamPlugin(
	ctx context.Context, id string,
) (*gen.DeleteUpstreamPluginResponse, error) {
	return gen.DeleteUpstreamPlugin(ctx, s.client.GraphQL, id)
}

// Toggle enables or disables a plugin.
func (s *PluginSDK) Toggle(
	ctx context.Context, id string, enabled bool,
) (*gen.TogglePluginResponse, error) {
	return gen.TogglePlugin(ctx, s.client.GraphQL, id, enabled)
}
