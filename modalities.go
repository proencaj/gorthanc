package gorthanc

import (
	"fmt"
	"log/slog"

	"github.com/proencaj/gorthanc/types"
)


func (c *Client) GetModalities() ([]string, error) {
	var modalities []string
	if err := c.get("modalities", &modalities); err != nil {
		return nil, err
	}
	return modalities, nil
}


func (c *Client) GetModalityDetails(modalityName string) (*types.Modality, error) {
	var modality types.Modality
	path := fmt.Sprintf("modalities/%s/configuration", modalityName)

	if err := c.get(path, &modality); err != nil {
		return nil, err
	}

	return &modality, nil
}


func (c *Client) CreateOrUpdateModality(modalityName string, request *types.ModalityCreateRequest) error {
	path := fmt.Sprintf("modalities/%s", modalityName)

	// Convert the request to the format expected by Orthanc
	// Orthanc expects an array: [AET, Host, Port, Manufacturer]
	modalityArray := []interface{}{
		request.AET,
		request.Host,
		request.Port,
	}

	// Add manufacturer if provided
	if request.Manufacturer != "" {
		modalityArray = append(modalityArray, request.Manufacturer)
	}

	if err := c.put(path, modalityArray, nil); err != nil {
		return err
	}

	return nil
}


func (c *Client) DeleteModality(modalityName string) error {
	path := fmt.Sprintf("modalities/%s", modalityName)

	if err := c.delete(path, nil); err != nil {
		return err
	}

	return nil
}


func (c *Client) EchoModality(modalityName string) error {
	path := fmt.Sprintf("modalities/%s/echo", modalityName)

	if err := c.post(path, nil, nil); err != nil {
		return err
	}

	return nil
}


func (c *Client) StoreToModality(modalityName, resourceID string) error {
	path := fmt.Sprintf("modalities/%s/store", modalityName)

	if err := c.post(path, resourceID, nil); err != nil {
		return err
	}

	return nil
}


func (c *Client) StoreToModalityWithOptions(modalityName string, request *types.ModalityStoreRequest) (*types.ModalityStoreResult, error) {
	path := fmt.Sprintf("modalities/%s/store", modalityName)

	var result types.ModalityStoreResult
	if err := c.post(path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}


func (c *Client) FindInModality(modalityName string, request *types.ModalityFindRequest) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("modalities/%s/query", modalityName)

	var queryResponse map[string]interface{}
	if err := c.post(path, request, &queryResponse); err != nil {
		return nil, err
	}

	fmt.Println("----------------")
	fmt.Println(queryResponse)

	// Extract the query ID from the response
	queryID, ok := queryResponse["ID"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get query ID from response")
	}

	// TODO: Turn 'expand' and 'simplify' into new parameters or options, so the user can change based on what is needs
	answersPath := fmt.Sprintf("queries/%s/answers?expand=true&simplify=true", queryID)
	
	var answers []map[string]interface{}
	if err := c.get(answersPath, &answers); err != nil {
		return nil, err
	}

	return answers, nil
}


func (c *Client) MoveFromModality(modalityName string, request *types.ModalityMoveRequest) (*types.ModalityMoveResult, error) {
	path := fmt.Sprintf("modalities/%s/move", modalityName)

	slog.Info("Path", "path", path)

	var result types.ModalityMoveResult
	if err := c.post(path, request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}


func (c *Client) GetFromModality(modalityName string, request *types.ModalityGetRequest) error {
	path := fmt.Sprintf("modalities/%s/get", modalityName)

	if err := c.post(path, request, nil); err != nil {
		return err
	}

	return nil
}
