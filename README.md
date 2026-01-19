# RelayWarden Go SDK

Official Go SDK for the RelayWarden API v1.

## Installation

```bash
go get github.com/relaywarden/go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/relaywarden/go-sdk"
)

func main() {
    // Initialize the client
    client := relaywarden.NewClient(
        "https://api.relaywarden.eu/api/v1",
        "your-api-token",
    )
    
    // Set project ID for project-scoped operations
    client.SetProjectID("your-project-id")
    
    // Send a message
    message, err := client.Messages.Send(context.Background(), map[string]interface{}{
        "from": map[string]interface{}{
            "email": "noreply@example.com",
            "name":  "Acme Corp",
        },
        "to": []map[string]interface{}{
            {"email": "user@example.com"},
        },
        "subject": "Welcome!",
        "html":    "<h1>Welcome!</h1>",
        "text":    "Welcome!",
    }, "unique-idempotency-key")
    
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Message ID: %v\n", message["data"].(map[string]interface{})["message_id"])
}
```

## Authentication

The SDK uses Bearer token authentication. Pass your API token when creating the client:

```go
client := relaywarden.NewClient(
    "https://api.relaywarden.eu/api/v1",
    "your-api-token",
)
```

## Resources

### Identity

```go
// Get current user/service account info
me, err := client.Identity.Me(context.Background())

// List teams
teams, err := client.Identity.Teams(context.Background())
```

### Projects

```go
// List projects
projects, err := client.Projects.List(context.Background(), map[string]string{
    "environment": "production",
})

// Create project
project, err := client.Projects.Create(context.Background(), map[string]interface{}{
    "name":        "Production",
    "environment": "production",
})

// Get project
project, err := client.Projects.Get(context.Background(), "project-id")

// Update project
project, err := client.Projects.Update(context.Background(), "project-id", map[string]interface{}{
    "name": "Updated Name",
})

// Delete project
err := client.Projects.Delete(context.Background(), "project-id")
```

### Messages

```go
// Send message with idempotency key
message, err := client.Messages.Send(context.Background(), map[string]interface{}{
    "from": map[string]interface{}{
        "email": "noreply@example.com",
    },
    "to": []map[string]interface{}{
        {"email": "user@example.com"},
    },
    "subject": "Hello",
    "html": "<h1>Hello</h1>",
}, "unique-idempotency-key")

// List messages
messages, err := client.Messages.List(context.Background(), map[string]string{
    "status":   "delivered",
    "per_page": "25",
})

// Get message
message, err := client.Messages.Get(context.Background(), "message-id")

// Get message timeline
timeline, err := client.Messages.GetTimeline(context.Background(), "message-id")
```

## Error Handling

The SDK returns specific error types for different error scenarios:

```go
import "github.com/relaywarden/go-sdk/errors"

message, err := client.Messages.Send(ctx, data, "")
if err != nil {
    switch e := err.(type) {
    case *errors.AuthenticationError:
        // 401 - Invalid or missing token
        fmt.Printf("Authentication failed: %s\n", e.Message)
    case *errors.ValidationErrorResponse:
        // 422 - Validation errors
        fmt.Printf("Validation failed: %s\n", e.Message)
        for _, detail := range e.Details {
            fmt.Printf("  %s: %s\n", detail.Field, detail.Message)
        }
    case *errors.RateLimitError:
        // 429 - Rate limit exceeded
        fmt.Printf("Rate limit exceeded. Retry after: %d seconds\n", e.RetryAfter)
    case *errors.APIError:
        // Other API errors
        fmt.Printf("API Error: %s [Request ID: %s]\n", e.Message, e.RequestID)
    default:
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Pagination

List endpoints return paginated responses:

```go
response, err := client.Messages.List(ctx, nil)
if err != nil {
    panic(err)
}

meta := response["meta"].(map[string]interface{})
currentPage := meta["current_page"].(float64)
total := meta["total"].(float64)

data := response["data"].([]interface{})
for _, item := range data {
    // Process each message
}
```

## Rate Limiting

The SDK automatically handles rate limits with exponential backoff. Rate limit information is available in the error:

```go
message, err := client.Messages.Send(ctx, data, "")
if rateLimitErr, ok := err.(*errors.RateLimitError); ok {
    retryAfter := rateLimitErr.RetryAfter // Seconds to wait
    // SDK will automatically retry, but you can also handle manually
}
```

## Configuration

```go
client := relaywarden.NewClient(
    "https://api.relaywarden.eu/api/v1",
    "your-token",
    relaywarden.ClientOptions{
        MaxRetries: 3,
        Timeout:    30 * time.Second,
    },
)
```

## Testing

```bash
go test ./...
go test -v -cover ./...
```

## License

MIT
