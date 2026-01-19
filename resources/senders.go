package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Senders handles sender-related API operations.
type Senders struct {
	client interfaces.Client
}

// NewSenders creates a new Senders resource.
func NewSenders(client interfaces.Client) *Senders {
	return &Senders{client: client}
}

// List returns all sender addresses for the current project.
func (r *Senders) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/senders", filters)
}

// Get returns a specific sender by ID.
func (r *Senders) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/senders/"+id, nil)
}

// Create creates a new sender address.
func (r *Senders) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/senders", data, nil)
}

// Delete deletes a sender address.
func (r *Senders) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/senders/"+id)
}

// Verify initiates sender verification.
func (r *Senders) Verify(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/senders/"+id+"/verify", nil, nil)
}
