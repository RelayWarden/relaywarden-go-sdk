package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Domains handles domain-related API operations.
type Domains struct {
	client interfaces.Client
}

// NewDomains creates a new Domains resource.
func NewDomains(client interfaces.Client) *Domains {
	return &Domains{client: client}
}

// List returns all sending domains for the current project.
func (r *Domains) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/domains", filters)
}

// Get returns a specific domain by ID.
func (r *Domains) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/domains/"+id, nil)
}

// Create creates a new sending domain.
func (r *Domains) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/domains", data, nil)
}

// Update updates a domain.
func (r *Domains) Update(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Patch(ctx, "/domains/"+id, data)
}

// Delete deletes a domain.
func (r *Domains) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/domains/"+id)
}

// GetDNSRecords returns DNS records required for domain verification.
func (r *Domains) GetDNSRecords(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/domains/"+id+"/dns-records", nil)
}

// GetChecks returns the current status of domain verification checks.
func (r *Domains) GetChecks(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/domains/"+id+"/checks", nil)
}

// Verify initiates domain verification.
func (r *Domains) Verify(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/domains/"+id+"/verify", nil, nil)
}

// RotateDKIM rotates DKIM signing keys for a domain.
func (r *Domains) RotateDKIM(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/domains/"+id+"/dkim/rotate", nil, nil)
}

// EnableProduction enables a domain for production use.
func (r *Domains) EnableProduction(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/domains/"+id+"/enable-production", nil, nil)
}
