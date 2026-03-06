package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// TaskSDK provides operations on background tasks.
type TaskSDK struct {
	client *Client
}

// List returns all running tasks.
func (s *TaskSDK) List(
	ctx context.Context,
) (*gen.ListTasksResponse, error) {
	return gen.ListTasks(ctx, s.client.GraphQL)
}

// Cancel cancels a running task.
func (s *TaskSDK) Cancel(
	ctx context.Context, id string,
) (*gen.CancelTaskResponse, error) {
	return gen.CancelTask(ctx, s.client.GraphQL, id)
}
