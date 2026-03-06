package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// RequestSDK provides operations on proxied HTTP requests.
type RequestSDK struct {
	client *Client
}

// ListRequestsOptions configures the List query.
type ListRequestsOptions struct {
	First   *int
	Last    *int
	After   *string
	Before  *string
	Filter  *string // HTTPQL filter expression
	Order   *gen.RequestResponseOrderInput
	ScopeID *string
}

// List returns paginated proxied requests.
func (s *RequestSDK) List(
	ctx context.Context, opts *ListRequestsOptions,
) (*gen.ListRequestsResponse, error) {
	var o ListRequestsOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListRequests(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before,
		o.Filter, o.Order, o.ScopeID,
	)
}

// ListByOffset returns requests using offset-based pagination.
func (s *RequestSDK) ListByOffset(
	ctx context.Context, opts *ListRequestsByOffsetOptions,
) (*gen.ListRequestsByOffsetResponse, error) {
	var o ListRequestsByOffsetOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListRequestsByOffset(
		ctx, s.client.GraphQL,
		o.Limit, o.Offset, o.Filter, o.Order, o.ScopeID,
	)
}

// ListRequestsByOffsetOptions configures the offset-based List query.
type ListRequestsByOffsetOptions struct {
	Limit   *int
	Offset  *int
	Filter  *string
	Order   *gen.RequestResponseOrderInput
	ScopeID *string
}

// Get returns a single request by ID, including raw bodies.
func (s *RequestSDK) Get(
	ctx context.Context, id string,
) (*gen.GetRequestResponse, error) {
	return gen.GetRequest(ctx, s.client.GraphQL, id)
}

// GetMetadata returns a single request without raw bodies.
func (s *RequestSDK) GetMetadata(
	ctx context.Context, id string,
) (*gen.GetRequestMetadataResponse, error) {
	return gen.GetRequestMetadata(ctx, s.client.GraphQL, id)
}
