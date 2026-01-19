package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// Projects handles project-related API operations.
type Projects struct {
	client interfaces.Client
}

// NewProjects creates a new Projects resource.
func NewProjects(client interfaces.Client) *Projects {
	return &Projects{client: client}
}

// List returns all projects for the current team.
func (r *Projects) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/projects", filters)
}

// Get returns a specific project by ID.
func (r *Projects) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/projects/"+id, nil)
}

// Create creates a new project.
func (r *Projects) Create(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Post(ctx, "/projects", data, nil)
}

// Update updates an existing project.
func (r *Projects) Update(ctx context.Context, id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Patch(ctx, "/projects/"+id, data)
}

// Delete deletes a project.
func (r *Projects) Delete(ctx context.Context, id string) error {
	return r.client.Delete(ctx, "/projects/"+id)
}
