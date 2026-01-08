package gorthanc

import (
	"fmt"

	"github.com/proencaj/gorthanc/types"
)


func (c *Client) GetDicomWebServers() ([]string, error) {
	var servers []string
	if err := c.get("dicom-web/servers", &servers); err != nil {
		return nil, err
	}
	return servers, nil
}


func (c *Client) GetDicomWebServersExpanded() (map[string]types.DicomWebServer, error) {
	var servers map[string]types.DicomWebServer
	if err := c.get("dicom-web/servers?expand=true", &servers); err != nil {
		return nil, err
	}
	return servers, nil
}


func (c *Client) CreateOrUpdateDicomWebServer(serverName string, request *types.DicomWebServerCreateRequest) error {
	path := fmt.Sprintf("dicom-web/servers/%s", serverName)

	if err := c.put(path, request, nil); err != nil {
		return err
	}

	return nil
}


func (c *Client) DeleteDicomWebServer(serverName string) error {
	path := fmt.Sprintf("dicom-web/servers/%s", serverName)

	if err := c.delete(path, nil); err != nil {
		return err
	}

	return nil
}

