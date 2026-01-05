package gorthanc

import (
	"fmt"

	"github.com/proencaj/gorthanc/types"
)

// GetDicomWebServers retrieves the list of configured remote DICOMweb servers
// This endpoint implements the GET /dicom-web/servers request
// Returns a list of server names
func (c *Client) GetDicomWebServers() ([]string, error) {
	var servers []string
	if err := c.get("dicom-web/servers", &servers); err != nil {
		return nil, err
	}
	return servers, nil
}

// GetDicomWebServersExpanded retrieves all DICOMweb servers with their full configurations
// This endpoint implements the GET /dicom-web/servers?expand=true request
// Returns a map of server names to their configurations
func (c *Client) GetDicomWebServersExpanded() (map[string]types.DicomWebServer, error) {
	var servers map[string]types.DicomWebServer
	if err := c.get("dicom-web/servers?expand=true", &servers); err != nil {
		return nil, err
	}
	return servers, nil
}

// CreateOrUpdateDicomWebServer creates or updates a DICOMweb server configuration
// This endpoint implements the PUT /dicom-web/servers/{name} request
func (c *Client) CreateOrUpdateDicomWebServer(serverName string, request *types.DicomWebServerCreateRequest) error {
	path := fmt.Sprintf("dicom-web/servers/%s", serverName)

	if err := c.put(path, request, nil); err != nil {
		return err
	}

	return nil
}

// DeleteDicomWebServer removes a DICOMweb server configuration
// This endpoint implements the DELETE /dicom-web/servers/{name} request
func (c *Client) DeleteDicomWebServer(serverName string) error {
	path := fmt.Sprintf("dicom-web/servers/%s", serverName)

	if err := c.delete(path, nil); err != nil {
		return err
	}

	return nil
}

