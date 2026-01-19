package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Identity handles identity-related API operations.
type Identity struct {
	client interfaces.Client
}

// NewIdentity creates a new Identity resource.
func NewIdentity(client interfaces.Client) *Identity {
	return &Identity{client: client}
}

// Me returns information about the currently authenticated user or service account.
func (r *Identity) Me(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/me", nil)
}

// Teams returns all teams the authenticated user belongs to.
func (r *Identity) Teams(ctx context.Context) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/teams", nil)
}
