package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Webhooks handles webhook-related API operations.
type Webhooks struct {
	client interfaces.Client
}

// NewWebhooks creates a new Webhooks resource.
func NewWebhooks(client interfaces.Client) *Webhooks {
	return &Webhooks{client: client}
}

// ListEndpoints returns all webhook endpoints for the current project.
func (r *Webhooks) ListEndpoints(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/webhooks/endpoints", filters)
}

// CreateEndpoint creates a new webhook endpoint.
func (r *Webhooks) CreateEndpoint(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/webhooks/endpoints", data, nil)
}

// UpdateEndpoint updates a webhook endpoint.
func (r *Webhooks) UpdateEndpoint(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Patch(ctx, "/webhooks/endpoints/"+id, data)
}

// DeleteEndpoint deletes a webhook endpoint.
func (r *Webhooks) DeleteEndpoint(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/webhooks/endpoints/"+id)
}

// ListDeliveries returns all delivery attempts for a webhook endpoint.
func (r *Webhooks) ListDeliveries(ctx context.Context, endpointID string, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/webhooks/endpoints/"+endpointID+"/deliveries", filters)
}

// TestEndpoint sends a test webhook to verify the endpoint is working.
func (r *Webhooks) TestEndpoint(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/webhooks/endpoints/"+id+"/test", nil, nil)
}

// ReplayDelivery replays a failed webhook delivery.
func (r *Webhooks) ReplayDelivery(ctx context.Context, deliveryID string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/webhooks/deliveries/"+deliveryID+"/replay", nil, nil)
}
