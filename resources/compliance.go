package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Compliance handles compliance-related API operations.
type Compliance struct {
	client interfaces.Client
}

// NewCompliance creates a new Compliance resource.
func NewCompliance(client interfaces.Client) *Compliance {
	return &Compliance{client: client}
}

// GetRetention returns data retention settings for the current team.
func (r *Compliance) GetRetention(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/compliance/retention", nil)
}

// UpdateRetention updates data retention settings.
func (r *Compliance) UpdateRetention(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Patch(ctx, "/compliance/retention", data)
}

// GetExportConfig returns available export formats and configuration.
func (r *Compliance) GetExportConfig(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/compliance/exports/config", nil)
}
