package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// HostedFileSDK provides operations on hosted files.
type HostedFileSDK struct {
	client *Client
}

// List returns all hosted files.
func (s *HostedFileSDK) List(
	ctx context.Context,
) (*gen.ListHostedFilesResponse, error) {
	return gen.ListHostedFiles(ctx, s.client.GraphQL)
}

// Rename renames a hosted file.
func (s *HostedFileSDK) Rename(
	ctx context.Context, id, name string,
) (*gen.RenameHostedFileResponse, error) {
	return gen.RenameHostedFile(ctx, s.client.GraphQL, id, name)
}

// Delete deletes a hosted file.
func (s *HostedFileSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteHostedFileResponse, error) {
	return gen.DeleteHostedFile(ctx, s.client.GraphQL, id)
}
