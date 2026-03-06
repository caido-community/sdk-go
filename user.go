package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// UserSDK provides operations on the current user.
type UserSDK struct {
	client *Client
}

// GetViewer returns the currently authenticated user.
func (s *UserSDK) GetViewer(
	ctx context.Context,
) (*gen.GetViewerResponse, error) {
	return gen.GetViewer(ctx, s.client.GraphQL)
}
