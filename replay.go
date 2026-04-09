package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// ReplaySDK provides operations on replay sessions and entries.
type ReplaySDK struct {
	client *Client
}

// ListSessionsOptions configures session listing.
type ListSessionsOptions struct {
	First  *int
	Last   *int
	After  *string
	Before *string
}

// ListSessions returns paginated replay sessions.
func (s *ReplaySDK) ListSessions(
	ctx context.Context, opts *ListSessionsOptions,
) (*gen.ListReplaySessionsResponse, error) {
	var o ListSessionsOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListReplaySessions(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before,
	)
}

// GetSession returns a single replay session by ID.
func (s *ReplaySDK) GetSession(
	ctx context.Context, id string,
) (*gen.GetReplaySessionResponse, error) {
	return gen.GetReplaySession(ctx, s.client.GraphQL, id)
}

// GetEntry returns a single replay entry by ID.
func (s *ReplaySDK) GetEntry(
	ctx context.Context, id string,
) (*gen.GetReplayEntryResponse, error) {
	return gen.GetReplayEntry(ctx, s.client.GraphQL, id)
}

// ListCollections returns paginated replay session collections.
func (s *ReplaySDK) ListCollections(
	ctx context.Context, opts *ListSessionsOptions,
) (*gen.ListReplaySessionCollectionsResponse, error) {
	var o ListSessionsOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListReplaySessionCollections(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before,
	)
}

// CreateSession creates a new replay session.
func (s *ReplaySDK) CreateSession(
	ctx context.Context, input *gen.CreateReplaySessionInput,
) (*gen.CreateReplaySessionResponse, error) {
	return gen.CreateReplaySession(ctx, s.client.GraphQL, *input)
}

// CreateCollection creates a new replay session collection.
func (s *ReplaySDK) CreateCollection(
	ctx context.Context, input *gen.CreateReplaySessionCollectionInput,
) (*gen.CreateReplaySessionCollectionResponse, error) {
	return gen.CreateReplaySessionCollection(ctx, s.client.GraphQL, *input)
}

// RenameSession renames a replay session.
func (s *ReplaySDK) RenameSession(
	ctx context.Context, id, name string,
) (*gen.RenameReplaySessionResponse, error) {
	return gen.RenameReplaySession(ctx, s.client.GraphQL, id, name)
}

// RenameCollection renames a replay session collection.
func (s *ReplaySDK) RenameCollection(
	ctx context.Context, id, name string,
) (*gen.RenameReplaySessionCollectionResponse, error) {
	return gen.RenameReplaySessionCollection(
		ctx, s.client.GraphQL, id, name,
	)
}

// DeleteSessions deletes replay sessions by IDs.
func (s *ReplaySDK) DeleteSessions(
	ctx context.Context, ids []string,
) (*gen.DeleteReplaySessionsResponse, error) {
	return gen.DeleteReplaySessions(ctx, s.client.GraphQL, ids)
}

// DeleteCollection deletes a replay session collection.
func (s *ReplaySDK) DeleteCollection(
	ctx context.Context, id string,
) (*gen.DeleteReplaySessionCollectionResponse, error) {
	return gen.DeleteReplaySessionCollection(ctx, s.client.GraphQL, id)
}

// MoveSession moves a session to a different collection.
func (s *ReplaySDK) MoveSession(
	ctx context.Context, id, collectionID string,
) (*gen.MoveReplaySessionResponse, error) {
	return gen.MoveReplaySession(ctx, s.client.GraphQL, id, collectionID)
}

// SendRequest starts a replay task (sends a request).
func (s *ReplaySDK) SendRequest(
	ctx context.Context, sessionID string, input *gen.StartReplayTaskInput,
) (*gen.StartReplayTaskResponse, error) {
	return gen.StartReplayTask(ctx, s.client.GraphQL, sessionID, *input)
}
