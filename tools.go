package gorthanc

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/proencaj/gorthanc/types"
)

// Find searches for DICOM resources in the local Orthanc database
// This endpoint implements the /tools/find POST request
// Returns a list of resource IDs (when Expand is false or nil)
func (c *Client) Find(request *types.ToolsFindRequest) ([]string, error) {
	if request.Expand != nil && *request.Expand {
		return nil, fmt.Errorf("use FindExpanded() when Expand is true")
	}

	var result []string
	if err := c.post("tools/find", request, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindExpanded searches for DICOM resources with expanded details
// This endpoint implements the /tools/find POST request with Expand=true
// Returns a list of expanded resource objects
func (c *Client) FindExpanded(request *types.ToolsFindRequest) ([]types.ToolsFindExpandedResource, error) {
	// Ensure Expand is set to true
	expand := true
	request.Expand = &expand

	var rawResult []json.RawMessage
	if err := c.post("tools/find", request, &rawResult); err != nil {
		return nil, err
	}

	// Parse the raw JSON into ToolsFindExpandedResource structs
	result := make([]types.ToolsFindExpandedResource, 0, len(rawResult))
	for _, raw := range rawResult {
		var resource types.ToolsFindExpandedResource
		if err := json.Unmarshal(raw, &resource); err != nil {
			return nil, fmt.Errorf("failed to unmarshal expanded resource: %w", err)
		}
		result = append(result, resource)
	}

	return result, nil
}

// Reset performs a hot restart of Orthanc
// This endpoint implements the /tools/reset POST request
// The configuration file will be read again
func (c *Client) Reset() error {
	return c.post("tools/reset", nil, nil)
}

// Shutdown shuts down Orthanc
// This endpoint implements the /tools/shutdown POST request
func (c *Client) Shutdown() error {
	return c.post("tools/shutdown", nil, nil)
}

// GetLogLevel retrieves the current log level
// This endpoint implements the GET /tools/log-level request
// Returns one of: "default", "verbose", or "trace"
func (c *Client) GetLogLevel() (types.LogLevel, error) {
	resp, err := c.getWithRawResponse("tools/log-level")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	level := strings.TrimSpace(string(body))
	return types.LogLevel(level), nil
}

// SetLogLevel dynamically changes the log level
// This endpoint implements the PUT /tools/log-level request
// Valid levels: LogLevelDefault, LogLevelVerbose, LogLevelTrace
// Note: This resets all category-specific log levels
func (c *Client) SetLogLevel(level types.LogLevel) error {
	return c.putWithPlainText("tools/log-level", string(level))
}
