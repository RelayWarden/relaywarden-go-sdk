package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Messages handles message-related API operations.
type Messages struct {
	client interfaces.Client
}

// NewMessages creates a new Messages resource.
func NewMessages(client interfaces.Client) *Messages {
	return &Messages{client: client}
}

// Send sends an email message.
func (r *Messages) Send(ctx context.Context, data map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	headers := make(map[string]string)
	if idempotencyKey != "" {
		headers["Idempotency-Key"] = idempotencyKey
	}
	return r.client.Post(ctx, "/messages", data, headers)
}

// List returns all messages for the current project.
func (r *Messages) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/messages", filters)
}

// Get returns a specific message by ID.
func (r *Messages) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/messages/"+id, nil)
}

// GetTimeline returns the complete timeline of events for a message.
func (r *Messages) GetTimeline(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/messages/"+id+"/timeline", nil)
}

// Cancel cancels a message that hasn't been sent yet.
func (r *Messages) Cancel(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/messages/"+id+"/cancel", nil, nil)
}

// Resend resends a previously sent message.
func (r *Messages) Resend(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/messages/"+id+"/resend", nil, nil)
}
