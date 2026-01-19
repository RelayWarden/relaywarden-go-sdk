package relaywarden

import (
	"github.com/relaywarden/go-sdk/interfaces"
	"github.com/relaywarden/go-sdk/resources"
)

// Client provides access to all API resources.
type Client struct {
	*client
	Identity        *resources.Identity
	Projects        *resources.Projects
	ServiceAccounts *resources.ServiceAccounts
	Domains         *resources.Domains
	Senders         *resources.Senders
	Templates       *resources.Templates
	Messages        *resources.Messages
	Events          *resources.Events
	Webhooks        *resources.Webhooks
	Suppressions    *resources.Suppressions
	Usage           *resources.Usage
	AuditLogs       *resources.AuditLogs
	Compliance      *resources.Compliance
}

// NewClient creates a new RelayWarden API client with all resources initialized.
func NewClient(baseURL, token string, opts ...ClientOptions) *Client {
	baseClient := newClient(baseURL, token, opts...)
	var clientInterface interfaces.Client = baseClient

	return &Client{
		client:          baseClient,
		Identity:        resources.NewIdentity(clientInterface),
		Projects:        resources.NewProjects(clientInterface),
		ServiceAccounts: resources.NewServiceAccounts(clientInterface),
		Domains:         resources.NewDomains(clientInterface),
		Senders:         resources.NewSenders(clientInterface),
		Templates:       resources.NewTemplates(clientInterface),
		Messages:        resources.NewMessages(clientInterface),
		Events:          resources.NewEvents(clientInterface),
		Webhooks:        resources.NewWebhooks(clientInterface),
		Suppressions:    resources.NewSuppressions(clientInterface),
		Usage:           resources.NewUsage(clientInterface),
		AuditLogs:       resources.NewAuditLogs(clientInterface),
		Compliance:      resources.NewCompliance(clientInterface),
	}
}

// SetProjectID sets the project ID for project-scoped operations.
func (c *Client) SetProjectID(projectID string) {
	c.client.SetProjectID(projectID)
}

// GetProjectID returns the current project ID.
func (c *Client) GetProjectID() *string {
	return c.client.GetProjectID()
}

// SetTeamID sets the team ID.
func (c *Client) SetTeamID(teamID string) {
	c.client.SetTeamID(teamID)
}

// GetTeamID returns the current team ID.
func (c *Client) GetTeamID() *string {
	return c.client.GetTeamID()
}
