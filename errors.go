package gorthanc

import "fmt"

// HTTPError represents an HTTP error response from the Orthanc server
type HTTPError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s - %s", e.StatusCode, e.Status, e.Body)
}

// IsHTTPError checks if an error is an HTTPError
func IsHTTPError(err error) bool {
	_, ok := err.(*HTTPError)
	return ok
}

// IsNotFound checks if an error is a 404 Not Found error
func IsNotFound(err error) bool {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.StatusCode == 404
	}
	return false
}

// IsUnauthorized checks if an error is a 401 Unauthorized error
func IsUnauthorized(err error) bool {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.StatusCode == 401
	}
	return false
}

// IsForbidden checks if an error is a 403 Forbidden error
func IsForbidden(err error) bool {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.StatusCode == 403
	}
	return false
}