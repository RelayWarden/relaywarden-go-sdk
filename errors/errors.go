package errors

import "fmt"

// APIError represents an error from the RelayWarden API.
type APIError struct {
	Message   string
	Code      int
	ErrorCode string
	RequestID string
	Details   []ValidationError
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (%d): %s [Request ID: %s]", e.Code, e.Message, e.RequestID)
}

// ValidationError represents a validation error detail.
type ValidationError struct {
	Field   string
	Message string
}

// AuthenticationError represents an authentication error.
type AuthenticationError struct {
	*APIError
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication failed: %s [Request ID: %s]", e.Message, e.RequestID)
}

// RateLimitError represents a rate limit error.
type RateLimitError struct {
	*APIError
	RetryAfter int
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("Rate limit exceeded: %s [Retry after: %d seconds, Request ID: %s]",
		e.Message, e.RetryAfter, e.RequestID)
}

// ValidationError represents validation errors.
type ValidationErrorResponse struct {
	*APIError
}

func (e *ValidationErrorResponse) Error() string {
	return fmt.Sprintf("Validation failed: %s [Request ID: %s]", e.Message, e.RequestID)
}
