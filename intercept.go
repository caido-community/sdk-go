package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// InterceptSDK provides operations on the MITM intercept proxy.
type InterceptSDK struct {
	client *Client
}

// ListInterceptEntriesOptions configures entry listing.
type ListInterceptEntriesOptions struct {
	First   *int
	Last    *int
	After   *string
	Before  *string
	Filter  *string
	Order   *gen.InterceptEntryOrderInput
	ScopeID *string
}

// ListEntries returns paginated intercept entries.
func (s *InterceptSDK) ListEntries(
	ctx context.Context, opts *ListInterceptEntriesOptions,
) (*gen.ListInterceptEntriesResponse, error) {
	var o ListInterceptEntriesOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListInterceptEntries(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before,
		o.Filter, o.Order, o.ScopeID,
	)
}

// GetEntry returns a single intercept entry by ID.
func (s *InterceptSDK) GetEntry(
	ctx context.Context, id string,
) (*gen.GetInterceptEntryResponse, error) {
	return gen.GetInterceptEntry(ctx, s.client.GraphQL, id)
}

// GetStatus returns whether intercept is running or paused.
func (s *InterceptSDK) GetStatus(
	ctx context.Context,
) (*gen.GetInterceptStatusResponse, error) {
	return gen.GetInterceptStatus(ctx, s.client.GraphQL)
}

// GetOptions returns the intercept configuration.
func (s *InterceptSDK) GetOptions(
	ctx context.Context,
) (*gen.GetInterceptOptionsResponse, error) {
	return gen.GetInterceptOptions(ctx, s.client.GraphQL)
}

// Forward forwards an intercepted message.
func (s *InterceptSDK) Forward(
	ctx context.Context,
	id string,
	input *gen.ForwardInterceptMessageInput,
) (*gen.ForwardInterceptMessageResponse, error) {
	return gen.ForwardInterceptMessage(
		ctx, s.client.GraphQL, id, input,
	)
}

// Drop drops an intercepted message.
func (s *InterceptSDK) Drop(
	ctx context.Context, id string,
) (*gen.DropInterceptMessageResponse, error) {
	return gen.DropInterceptMessage(ctx, s.client.GraphQL, id)
}

// Pause pauses interception.
func (s *InterceptSDK) Pause(
	ctx context.Context,
) (*gen.PauseInterceptResponse, error) {
	return gen.PauseIntercept(ctx, s.client.GraphQL)
}

// Resume resumes interception.
func (s *InterceptSDK) Resume(
	ctx context.Context,
) (*gen.ResumeInterceptResponse, error) {
	return gen.ResumeIntercept(ctx, s.client.GraphQL)
}

// SetOptions updates the intercept configuration.
func (s *InterceptSDK) SetOptions(
	ctx context.Context, input *gen.InterceptOptionsInput,
) (*gen.SetInterceptOptionsResponse, error) {
	return gen.SetInterceptOptions(ctx, s.client.GraphQL, input)
}

// DeleteEntries deletes intercept entries matching a filter.
func (s *InterceptSDK) DeleteEntries(
	ctx context.Context, filter *string, scopeID *string,
) (*gen.DeleteInterceptEntriesResponse, error) {
	return gen.DeleteInterceptEntries(
		ctx, s.client.GraphQL, filter, scopeID,
	)
}

// DeleteEntry deletes a single intercept entry.
func (s *InterceptSDK) DeleteEntry(
	ctx context.Context, id string,
) (*gen.DeleteInterceptEntryResponse, error) {
	return gen.DeleteInterceptEntry(ctx, s.client.GraphQL, id)
}
