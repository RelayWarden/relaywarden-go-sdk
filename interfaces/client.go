package interfaces

import (
	"context"
)

// Client defines the interface for making HTTP requests.
// This allows resources to use the client without creating import cycles.
type Client interface {
	Get(ctx context.Context, path string, query map[string]string) (map[string]interface{}, error)
	Post(ctx context.Context, path string, body interface{}, headers map[string]string) (map[string]interface{}, error)
	Patch(ctx context.Context, path string, body interface{}) (map[string]interface{}, error)
	Delete(ctx context.Context, path string) error
	SetProjectID(projectID string)
	GetProjectID() *string
	SetTeamID(teamID string)
	GetTeamID() *string
}
