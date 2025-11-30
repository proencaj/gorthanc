package gorthanc

import (
	"fmt"
	"strconv"

	"github.com/proencaj/gorthanc/types"
)


func (c *Client) GetPatients(params *types.PatientQueryParams) ([]string, error) {
	path := "patients"

	if params != nil {
		path = c.buildPatientsPath(path, params)
	}

	// If expand is requested, we need to handle the response differently
	if params != nil && params.Expand {
		return nil, fmt.Errorf("not implemented, use GetPatient for each patient to get details")
	}

	var patientIDs []string
	if err := c.get(path, &patientIDs); err != nil {
		return nil, err
	}

	return patientIDs, nil
}


func (c *Client) GetPatientDetails(patientID string) (*types.Patient, error) {
	var patient types.Patient
	path := fmt.Sprintf("patients/%s", patientID)

	if err := c.get(path, &patient); err != nil {
		return nil, err
	}

	return &patient, nil
}

func (c *Client) AnonymizePatient(studyID string, anonymizeRequest *types.PatientAnonymizeRequest) (*types.PatientAnonymizeResponse, error) {
	var result types.PatientAnonymizeResponse
	path := fmt.Sprintf("patients/%s/anonymize", studyID)
	anonymizeRequest.Asynchronous = BoolPtr(false)

	if err := c.post(path, anonymizeRequest, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// buildPatientsPath constructs the path with query parameters
func (c *Client) buildPatientsPath(basePath string, params *types.PatientQueryParams) string {
	if params == nil {
		return basePath
	}

	// Build query string manually
	queryParams := ""
	separator := "?"

	if params.Expand {
		queryParams += separator + "expand"
		separator = "&"
	}

	if params.Since >= 0 {
		queryParams += separator + "since=" + strconv.Itoa(params.Since)
		separator = "&"
	}

	if params.Limit >= 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
	}

	return basePath + queryParams
}

func (c *Client) DeletePatient(patientID string) error {
	path := fmt.Sprintf("patients/%s", patientID)
	return c.delete(path, nil)
}

func (c *Client) GetPatientStatistics(patientID string) (*types.PatientStatistics, error) {
	var stats types.PatientStatistics
	path := fmt.Sprintf("patients/%s/statistics", patientID)

	if err := c.get(path, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
