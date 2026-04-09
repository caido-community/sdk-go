package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// FindingSDK provides operations on security findings.
type FindingSDK struct {
	client *Client
}

// ListFindingsOptions configures the List query.
type ListFindingsOptions struct {
	First  *int
	Last   *int
	After  *string
	Before *string
	Filter *gen.FilterClauseFindingInput
	Order  *gen.FindingOrderInput
}

// List returns paginated findings.
func (s *FindingSDK) List(
	ctx context.Context, opts *ListFindingsOptions,
) (*gen.ListFindingsResponse, error) {
	var o ListFindingsOptions
	if opts != nil {
		o = *opts
	}
	return gen.ListFindings(
		ctx, s.client.GraphQL,
		o.First, o.Last, o.After, o.Before, o.Filter, o.Order,
	)
}

// Get returns a single finding by ID.
func (s *FindingSDK) Get(
	ctx context.Context, id string,
) (*gen.GetFindingResponse, error) {
	return gen.GetFinding(ctx, s.client.GraphQL, id)
}

// ListReporters returns all finding reporter names.
func (s *FindingSDK) ListReporters(
	ctx context.Context,
) (*gen.ListFindingReportersResponse, error) {
	return gen.ListFindingReporters(ctx, s.client.GraphQL)
}

// Create creates a new finding attached to a request.
func (s *FindingSDK) Create(
	ctx context.Context, requestID string, input *gen.CreateFindingInput,
) (*gen.CreateFindingResponse, error) {
	return gen.CreateFinding(ctx, s.client.GraphQL, requestID, *input)
}

// Delete deletes findings.
func (s *FindingSDK) Delete(
	ctx context.Context, input *gen.DeleteFindingsInput,
) (*gen.DeleteFindingsResponse, error) {
	return gen.DeleteFindings(ctx, s.client.GraphQL, input)
}

// Export exports findings.
func (s *FindingSDK) Export(
	ctx context.Context, input *gen.ExportFindingsInput,
) (*gen.ExportFindingsResponse, error) {
	return gen.ExportFindings(ctx, s.client.GraphQL, *input)
}
