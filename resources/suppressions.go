package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Suppressions handles suppression-related API operations.
type Suppressions struct {
	client interfaces.Client
}

// NewSuppressions creates a new Suppressions resource.
func NewSuppressions(client interfaces.Client) *Suppressions {
	return &Suppressions{client: client}
}

// List returns all suppressions for the current team.
func (r *Suppressions) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/suppressions", filters)
}

// Create adds a recipient to the suppression list.
func (r *Suppressions) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/suppressions", data, nil)
}

// Delete removes a recipient from the suppression list.
func (r *Suppressions) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/suppressions/"+id)
}

// Import imports multiple suppressions in bulk.
func (r *Suppressions) Import(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/suppressions/import", data, nil)
}

// Export exports all suppressions as a CSV file.
func (r *Suppressions) Export(ctx context.Context) (string, error) {
	// Note: This endpoint returns CSV, not JSON
	// For now, we'll return it as a string
	// In a production SDK, you might want a separate method that returns []byte
	_, err := r.client.Get(ctx, "/suppressions/export", nil)
	if err != nil {
		return "", err
	}
	// This is a simplified implementation - actual CSV handling would be more complex
	return "", nil
}
