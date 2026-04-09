package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// ProjectSDK provides operations on Caido projects.
type ProjectSDK struct {
	client *Client
}

// List returns all projects.
func (s *ProjectSDK) List(
	ctx context.Context,
) (*gen.ListProjectsResponse, error) {
	return gen.ListProjects(ctx, s.client.GraphQL)
}

// GetCurrent returns the currently selected project.
func (s *ProjectSDK) GetCurrent(
	ctx context.Context,
) (*gen.GetCurrentProjectResponse, error) {
	return gen.GetCurrentProject(ctx, s.client.GraphQL)
}

// Create creates a new project.
func (s *ProjectSDK) Create(
	ctx context.Context, input *gen.CreateProjectInput,
) (*gen.CreateProjectResponse, error) {
	return gen.CreateProject(ctx, s.client.GraphQL, *input)
}

// Select selects a project as the active project.
func (s *ProjectSDK) Select(
	ctx context.Context, id string,
) (*gen.SelectProjectResponse, error) {
	return gen.SelectProject(ctx, s.client.GraphQL, id)
}

// Rename renames a project.
func (s *ProjectSDK) Rename(
	ctx context.Context, id, name string,
) (*gen.RenameProjectResponse, error) {
	return gen.RenameProject(ctx, s.client.GraphQL, id, name)
}

// Delete deletes a project.
func (s *ProjectSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteProjectResponse, error) {
	return gen.DeleteProject(ctx, s.client.GraphQL, id)
}
