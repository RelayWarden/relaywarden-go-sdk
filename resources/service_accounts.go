package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// ServiceAccounts handles service account-related API operations.
type ServiceAccounts struct {
	client interfaces.Client
}

// NewServiceAccounts creates a new ServiceAccounts resource.
func NewServiceAccounts(client interfaces.Client) *ServiceAccounts {
	return &ServiceAccounts{client: client}
}

// List returns all service accounts for the current team.
func (r *ServiceAccounts) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/service-accounts", filters)
}

// Create creates a new service account.
func (r *ServiceAccounts) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/service-accounts", data, nil)
}

// Delete deletes a service account.
func (r *ServiceAccounts) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/service-accounts/"+id)
}

// CreateToken creates a new API token for a service account.
func (r *ServiceAccounts) CreateToken(ctx context.Context, serviceAccountID string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/service-accounts/"+serviceAccountID+"/tokens", data, nil)
}

// DeleteToken deletes an API token.
func (r *ServiceAccounts) DeleteToken(ctx context.Context, tokenID string) error {
	return r.client.Delete(ctx, "/tokens/"+tokenID)
}
