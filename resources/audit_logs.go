package resources

import (
	"context"

	"github.com/relaywarden/go-sdk/interfaces"
)

// AuditLogs handles audit log-related API operations.
type AuditLogs struct {
	client interfaces.Client
}

// NewAuditLogs creates a new AuditLogs resource.
func NewAuditLogs(client interfaces.Client) *AuditLogs {
	return &AuditLogs{client: client}
}

// List returns audit logs for the current team.
func (r *AuditLogs) List(ctx context.Context, filters map[string]string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/audit-logs", filters)
}

// Get returns a specific audit log entry by ID.
func (r *AuditLogs) Get(ctx context.Context, id string) (map[string]interface{}, error) {
	return r.client.Get(ctx, "/audit-logs/"+id, nil)
}
