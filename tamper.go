package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// TamperSDK provides operations on Match & Replace (tamper) rules.
type TamperSDK struct {
	client *Client
}

// ListCollections returns all tamper rule collections with their rules.
func (s *TamperSDK) ListCollections(
	ctx context.Context,
) (*gen.ListTamperRuleCollectionsResponse, error) {
	return gen.ListTamperRuleCollections(ctx, s.client.GraphQL)
}

// GetRule returns a single tamper rule by ID.
func (s *TamperSDK) GetRule(
	ctx context.Context, id string,
) (*gen.GetTamperRuleResponse, error) {
	return gen.GetTamperRule(ctx, s.client.GraphQL, id)
}

// CreateCollection creates a new tamper rule collection.
func (s *TamperSDK) CreateCollection(
	ctx context.Context, input *gen.CreateTamperRuleCollectionInput,
) (*gen.CreateTamperRuleCollectionResponse, error) {
	return gen.CreateTamperRuleCollection(
		ctx, s.client.GraphQL, *input,
	)
}

// CreateRule creates a new tamper rule.
func (s *TamperSDK) CreateRule(
	ctx context.Context, input *gen.CreateTamperRuleInput,
) (*gen.CreateTamperRuleResponse, error) {
	return gen.CreateTamperRule(ctx, s.client.GraphQL, *input)
}

// UpdateRule updates an existing tamper rule.
func (s *TamperSDK) UpdateRule(
	ctx context.Context, id string, input *gen.UpdateTamperRuleInput,
) (*gen.UpdateTamperRuleResponse, error) {
	return gen.UpdateTamperRule(ctx, s.client.GraphQL, id, *input)
}

// DeleteRule deletes a tamper rule.
func (s *TamperSDK) DeleteRule(
	ctx context.Context, id string,
) (*gen.DeleteTamperRuleResponse, error) {
	return gen.DeleteTamperRule(ctx, s.client.GraphQL, id)
}

// ToggleRule enables or disables a tamper rule.
func (s *TamperSDK) ToggleRule(
	ctx context.Context, id string, enabled bool,
) (*gen.ToggleTamperRuleResponse, error) {
	return gen.ToggleTamperRule(ctx, s.client.GraphQL, id, enabled)
}

// DeleteCollection deletes a tamper rule collection.
func (s *TamperSDK) DeleteCollection(
	ctx context.Context, id string,
) (*gen.DeleteTamperRuleCollectionResponse, error) {
	return gen.DeleteTamperRuleCollection(ctx, s.client.GraphQL, id)
}
