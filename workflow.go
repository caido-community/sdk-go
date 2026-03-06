package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// WorkflowSDK provides operations on automation workflows.
type WorkflowSDK struct {
	client *Client
}

// List returns all workflows.
func (s *WorkflowSDK) List(
	ctx context.Context,
) (*gen.ListWorkflowsResponse, error) {
	return gen.ListWorkflows(ctx, s.client.GraphQL)
}

// Get returns a single workflow by ID, including its definition.
func (s *WorkflowSDK) Get(
	ctx context.Context, id string,
) (*gen.GetWorkflowResponse, error) {
	return gen.GetWorkflow(ctx, s.client.GraphQL, id)
}

// ListNodeDefinitions returns available workflow node definitions.
func (s *WorkflowSDK) ListNodeDefinitions(
	ctx context.Context,
) (*gen.ListWorkflowNodeDefinitionsResponse, error) {
	return gen.ListWorkflowNodeDefinitions(ctx, s.client.GraphQL)
}

// Create creates a new workflow.
func (s *WorkflowSDK) Create(
	ctx context.Context, input *gen.CreateWorkflowInput,
) (*gen.CreateWorkflowResponse, error) {
	return gen.CreateWorkflow(ctx, s.client.GraphQL, input)
}

// Rename renames a workflow.
func (s *WorkflowSDK) Rename(
	ctx context.Context, id, name string,
) (*gen.RenameWorkflowResponse, error) {
	return gen.RenameWorkflow(ctx, s.client.GraphQL, id, name)
}

// Delete deletes a workflow.
func (s *WorkflowSDK) Delete(
	ctx context.Context, id string,
) (*gen.DeleteWorkflowResponse, error) {
	return gen.DeleteWorkflow(ctx, s.client.GraphQL, id)
}

// Globalize converts a project workflow to a global one.
func (s *WorkflowSDK) Globalize(
	ctx context.Context, id string,
) (*gen.GlobalizeWorkflowResponse, error) {
	return gen.GlobalizeWorkflow(ctx, s.client.GraphQL, id)
}

// Localize converts a global workflow to a project one.
func (s *WorkflowSDK) Localize(
	ctx context.Context, id string,
) (*gen.LocalizeWorkflowResponse, error) {
	return gen.LocalizeWorkflow(ctx, s.client.GraphQL, id)
}
