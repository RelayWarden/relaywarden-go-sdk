package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Templates handles template-related API operations.
type Templates struct {
	client interfaces.Client
}

// NewTemplates creates a new Templates resource.
func NewTemplates(client interfaces.Client) *Templates {
	return &Templates{client: client}
}

// List returns all templates for the current project.
func (r *Templates) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/templates", filters)
}

// Get returns a specific template by ID.
func (r *Templates) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/templates/"+id, nil)
}

// Create creates a new template.
func (r *Templates) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/templates", data, nil)
}

// Update updates an existing template.
func (r *Templates) Update(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Patch(ctx, "/templates/"+id, data)
}

// Delete deletes a template.
func (r *Templates) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/templates/"+id)
}

// ListVersions returns all versions of a template.
func (r *Templates) ListVersions(ctx context.Context, id string, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/templates/"+id+"/versions", filters)
}

// CreateVersion creates a new version of a template.
func (r *Templates) CreateVersion(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/templates/"+id+"/versions", data, nil)
}

// Render renders a template with provided data.
func (r *Templates) Render(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/templates/"+id+"/render", data, nil)
}

// TestSend sends a test email using the template.
func (r *Templates) TestSend(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/templates/"+id+"/test-send", data, nil)
}
