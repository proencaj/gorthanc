package gorthanc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"io"
	
	"github.com/proencaj/gorthanc/types"
)


func (c *Client) GetAllInstances(params *types.InstancesQueryParams) ([]string, error) {
	path := "instances"

	if params != nil {
		path = c.buildInstancePath(path, params)
	}	

	var instanceIDs []string
	if err := c.get(path, &instanceIDs); err != nil {
		return nil, err
	}

	return instanceIDs, nil
}


func (c *Client) GetInstanceDetails(instanceID string) (*types.Instance, error) {
	var instance types.Instance
	path := fmt.Sprintf("instances/%s", instanceID)

	if err := c.get(path, &instance); err != nil {
		return nil, err
	}

	return &instance, nil
}


func (c *Client) DeleteInstance(instanceID string) error {
	path := fmt.Sprintf("instances/%s", instanceID)
	return c.delete(path, nil) 
}


func (c *Client) UploadDicomFile(reader io.Reader) (*types.UploadDicomFileResponse, error) {
	resp, err := c.doRequest(http.MethodPost, "instances", reader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UploadDicomFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}


func (c *Client) AnonymizeInstance(instanceID string, anonymizeRequest *types.InstancesAnonymizeRequest) (*http.Response, error) {
	path := fmt.Sprintf("instances/%s/anonymize", instanceID)
	return c.postWithBodyAndRawResponse(path, anonymizeRequest)
}


func (c *Client) DownloadDicomFile(instanceID string) (*http.Response, error) {
	path := fmt.Sprintf("instances/%s/file", instanceID)
	return c.getWithRawResponse(path)
}


func (c *Client) GetInstanceTags(instanceID string, params *types.GetInstanceTagsQueryParams) (map[string]interface{}, error) {
	var tags map[string]interface{}
	path := fmt.Sprintf("instances/%s/tags", instanceID)

	if params != nil {
		path = c.buildInstanceTagsPath(path, params)
	}

	if err := c.get(path, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

// buildInstancePath constructs the path with query parameters
func (c *Client) buildInstancePath(basePath string, params *types.InstancesQueryParams) string {
	if params == nil {
		return basePath
	}

	// Build query string manually
	queryParams := ""
	separator := "?"

	if params.Since >= 0 {
		queryParams += separator + "since=" + strconv.Itoa(params.Since)
		separator = "&"
	}

	if params.Limit >= 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
	}

	return basePath + queryParams
}

func (c *Client) buildInstanceTagsPath(basePath string, params *types.GetInstanceTagsQueryParams) string {
	if params == nil {
		return basePath
	}

	// Build query string manually
	queryParams := ""
	separator := "?"

	if params.Short {
		queryParams += separator + "short=true"
		separator = "&"
	}

	if params.Simplify {
		queryParams += separator + "simplify=true" 
	}

	if params.Whole {
		queryParams += separator + "whole=true" 
	}

	return basePath + queryParams
}
