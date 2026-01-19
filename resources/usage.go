package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Usage handles usage-related API operations.
type Usage struct {
	client interfaces.Client
}

// NewUsage creates a new Usage resource.
func NewUsage(client interfaces.Client) *Usage {
	return &Usage{client: client}
}

// GetDaily returns daily usage statistics for the current team.
func (r *Usage) GetDaily(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/usage/daily", filters)
}

// GetLimits returns current usage limits and remaining quota.
func (r *Usage) GetLimits(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/limits", nil)
}

// GetDiagnostics returns system health and diagnostic information.
func (r *Usage) GetDiagnostics(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/diagnostics", nil)
}
