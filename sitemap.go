package caido

import (
	"context"

	gen "github.com/caido-community/sdk-go/graphql"
)

// SitemapSDK provides operations on the site map tree.
type SitemapSDK struct {
	client *Client
}

// ListRootEntries returns root-level sitemap entries.
func (s *SitemapSDK) ListRootEntries(
	ctx context.Context, scopeID *string,
) (*gen.ListSitemapRootEntriesResponse, error) {
	return gen.ListSitemapRootEntries(ctx, s.client.GraphQL, scopeID)
}

// ListDescendantEntries returns children of a sitemap entry.
func (s *SitemapSDK) ListDescendantEntries(
	ctx context.Context,
	parentID string,
	depth gen.SitemapDescendantsDepth,
) (*gen.ListSitemapDescendantEntriesResponse, error) {
	return gen.ListSitemapDescendantEntries(
		ctx, s.client.GraphQL, parentID, depth,
	)
}

// GetEntry returns a single sitemap entry.
func (s *SitemapSDK) GetEntry(
	ctx context.Context, id string,
) (*gen.GetSitemapEntryResponse, error) {
	return gen.GetSitemapEntry(ctx, s.client.GraphQL, id)
}

// Clear removes all sitemap entries.
func (s *SitemapSDK) Clear(
	ctx context.Context,
) (*gen.ClearSitemapEntriesResponse, error) {
	return gen.ClearSitemapEntries(ctx, s.client.GraphQL)
}

// Delete removes specific sitemap entries.
func (s *SitemapSDK) Delete(
	ctx context.Context, ids []string,
) (*gen.DeleteSitemapEntriesResponse, error) {
	return gen.DeleteSitemapEntries(ctx, s.client.GraphQL, ids)
}
