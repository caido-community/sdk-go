package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// AutomateSDK provides operations on fuzzing (Automate) sessions.
type AutomateSDK struct {
	client *Client
}

// ListSessionsOptions configures session listing.
type ListAutomateSessionsOptions struct {
	First  *int
	Last   *int
	After  *string
	Before *string
}

// ListSessions returns paginated Automate sessions.
func (s *AutomateSDK) ListSessions(
	ctx context.Context, opts *ListAutomateSessionsOptions,
) (*gen.ListAutomateSessionsResponse, error) {
	var o ListAutomateSessionsOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListAutomateSessions(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before,
	)
}

// GetSession returns a single Automate session by ID.
func (s *AutomateSDK) GetSession(
	ctx context.Context, id string,
) (*gen.GetAutomateSessionResponse, error) {
	return gen.GetAutomateSession(ctx, s.client.GraphQL, id)
}

// GetEntry returns a single Automate entry by ID.
func (s *AutomateSDK) GetEntry(
	ctx context.Context, id string,
) (*gen.GetAutomateEntryResponse, error) {
	return gen.GetAutomateEntry(ctx, s.client.GraphQL, id)
}

// CreateSession creates a new Automate session.
func (s *AutomateSDK) CreateSession(
	ctx context.Context, input *gen.CreateAutomateSessionInput,
) (*gen.CreateAutomateSessionResponse, error) {
	return gen.CreateAutomateSession(ctx, s.client.GraphQL, input)
}

// RenameSession renames an Automate session.
func (s *AutomateSDK) RenameSession(
	ctx context.Context, id, name string,
) (*gen.RenameAutomateSessionResponse, error) {
	return gen.RenameAutomateSession(ctx, s.client.GraphQL, id, name)
}

// DeleteSession deletes an Automate session.
func (s *AutomateSDK) DeleteSession(
	ctx context.Context, id string,
) (*gen.DeleteAutomateSessionResponse, error) {
	return gen.DeleteAutomateSession(ctx, s.client.GraphQL, id)
}

// StartTask starts an Automate fuzzing task.
func (s *AutomateSDK) StartTask(
	ctx context.Context, sessionID string,
) (*gen.StartAutomateTaskResponse, error) {
	return gen.StartAutomateTask(ctx, s.client.GraphQL, sessionID)
}

// CancelTask cancels an Automate task.
func (s *AutomateSDK) CancelTask(
	ctx context.Context, id string,
) (*gen.CancelAutomateTaskResponse, error) {
	return gen.CancelAutomateTask(ctx, s.client.GraphQL, id)
}

// PauseTask pauses an Automate task.
func (s *AutomateSDK) PauseTask(
	ctx context.Context, id string,
) (*gen.PauseAutomateTaskResponse, error) {
	return gen.PauseAutomateTask(ctx, s.client.GraphQL, id)
}

// ResumeTask resumes a paused Automate task.
func (s *AutomateSDK) ResumeTask(
	ctx context.Context, id string,
) (*gen.ResumeAutomateTaskResponse, error) {
	return gen.ResumeAutomateTask(ctx, s.client.GraphQL, id)
}
