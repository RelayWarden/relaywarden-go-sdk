package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Events handles event-related API operations.
type Events struct {
	client interfaces.Client
}

// NewEvents creates a new Events resource.
func NewEvents(client interfaces.Client) *Events {
	return &Events{client: client}
}

// List returns all events for the current team.
func (r *Events) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/events", filters)
}

// Get returns a specific event by ID.
func (r *Events) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/events/"+id, nil)
}
