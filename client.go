package relaywarden

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/relaywarden/go-sdk/errors"
)

// client is the internal client for making HTTP requests.
type client struct {
	baseURL    string
	token      string
	httpClient *http.Client
	projectID  *string
	teamID     *string
	maxRetries int
	timeout    time.Duration
}

// ClientOptions contains optional configuration for the client.
type ClientOptions struct {
	MaxRetries int
	Timeout    time.Duration
}

// newClient creates a new internal client for making HTTP requests.
func newClient(baseURL, token string, opts ...ClientOptions) *client {
	options := ClientOptions{
		MaxRetries: 3,
		Timeout:    30 * time.Second,
	}
	if len(opts) > 0 {
		options = opts[0]
	}

	c := &client{
		baseURL:    baseURL,
		token:      token,
		maxRetries: options.MaxRetries,
		timeout:    options.Timeout,
		httpClient: &http.Client{
			Timeout: options.Timeout,
		},
	}

	return c
}

// SetProjectID sets the project ID for project-scoped operations.
func (c *client) SetProjectID(projectID string) {
	c.projectID = &projectID
}

// GetProjectID returns the current project ID.
func (c *client) GetProjectID() *string {
	return c.projectID
}

// SetTeamID sets the team ID.
func (c *client) SetTeamID(teamID string) {
	c.teamID = &teamID
}

// GetTeamID returns the current team ID.
func (c *client) GetTeamID() *string {
	return c.teamID
}

// request makes an HTTP request with retry logic.
func (c *client) request(ctx context.Context, method, path string, body interface{}, headers map[string]string) (map[string]interface{}, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set project/team headers
	if c.projectID != nil {
		req.Header.Set("X-Project-Id", *c.projectID)
	}
	if c.teamID != nil {
		req.Header.Set("X-Team-Id", *c.teamID)
	}

	// Set custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
				continue
			}
			return nil, fmt.Errorf("request failed: %w", err)
		}

		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}

		if resp.StatusCode == 204 {
			return nil, nil
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var result map[string]interface{}
			if len(bodyBytes) > 0 {
				if err := json.Unmarshal(bodyBytes, &result); err != nil {
					return nil, fmt.Errorf("failed to unmarshal response: %w", err)
				}
			}
			return result, nil
		}

		// Handle errors
		var errorResp map[string]interface{}
		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &errorResp); err == nil {
				if apiErr := c.handleErrorResponse(resp, errorResp); apiErr != nil {
					// Check if rate limit and retry
					if rateLimitErr, ok := apiErr.(*errors.RateLimitError); ok && attempt < c.maxRetries {
						time.Sleep(time.Duration(rateLimitErr.RetryAfter) * time.Second)
						lastErr = apiErr
						continue
					}
					return nil, apiErr
				}
			}
		}

		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if lastErr != nil {
		return nil, lastErr
	}

	return nil, fmt.Errorf("request failed after %d retries", c.maxRetries)
}

// handleErrorResponse parses error responses and returns appropriate error types.
func (c *client) handleErrorResponse(resp *http.Response, body map[string]interface{}) error {
	statusCode := resp.StatusCode
	requestID := ""
	errorCode := ""
	message := "An error occurred"
	var details []errors.ValidationError

	if meta, ok := body["meta"].(map[string]interface{}); ok {
		if id, ok := meta["request_id"].(string); ok {
			requestID = id
		}
	}

	if err, ok := body["error"].(map[string]interface{}); ok {
		if code, ok := err["code"].(string); ok {
			errorCode = code
		}
		if msg, ok := err["message"].(string); ok {
			message = msg
		}
		if det, ok := err["details"].([]interface{}); ok {
			for _, d := range det {
				if detailMap, ok := d.(map[string]interface{}); ok {
					detail := errors.ValidationError{}
					if field, ok := detailMap["field"].(string); ok {
						detail.Field = field
					}
					if msg, ok := detailMap["message"].(string); ok {
						detail.Message = msg
					}
					details = append(details, detail)
				}
			}
		}
	}

	apiErr := &errors.APIError{
		Message:   message,
		Code:      statusCode,
		ErrorCode: errorCode,
		RequestID: requestID,
		Details:   details,
	}

	switch statusCode {
	case 401:
		return &errors.AuthenticationError{APIError: apiErr}
	case 422:
		return &errors.ValidationErrorResponse{APIError: apiErr}
	case 429:
		retryAfter := 60
		if retryHeader := resp.Header.Get("Retry-After"); retryHeader != "" {
			if ra, err := time.ParseDuration(retryHeader + "s"); err == nil {
				retryAfter = int(ra.Seconds())
			}
		}
		return &errors.RateLimitError{
			APIError:   apiErr,
			RetryAfter: retryAfter,
		}
	default:
		return apiErr
	}
}

// Get makes a GET request.
func (c *client) Get(ctx context.Context, path string, query map[string]string) (map[string]interface{}, error) {
	if len(query) > 0 {
		path += "?"
		first := true
		for k, v := range query {
			if !first {
				path += "&"
			}
			path += k + "=" + url.QueryEscape(v)
			first = false
		}
	}
	return c.request(ctx, "GET", path, nil, nil)
}

// Post makes a POST request.
func (c *client) Post(ctx context.Context, path string, body interface{}, headers map[string]string) (map[string]interface{}, error) {
	return c.request(ctx, "POST", path, body, headers)
}

// Patch makes a PATCH request.
func (c *client) Patch(ctx context.Context, path string, body interface{}) (map[string]interface{}, error) {
	return c.request(ctx, "PATCH", path, body, nil)
}

// Delete makes a DELETE request.
func (c *client) Delete(ctx context.Context, path string) error {
	_, err := c.request(ctx, "DELETE", path, nil, nil)
	return err
}
