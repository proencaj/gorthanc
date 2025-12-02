package gorthanc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/proencaj/gorthanc/types"
)

func (c *Client) GetStudies(params *types.StudiesQueryParams) ([]string, error) {
	path := "studies"
	
	if params != nil {
		path = c.buildStudiesPath(path, params)
	}

	// If expand is requested, we need to handle the response differently
	if params != nil && params.Expand {
		return nil, fmt.Errorf("for expanded results, use GetStudiesExpanded method")
	}

	var studyIDs []string
	if err := c.get(path, &studyIDs); err != nil {
		return nil, err
	}

	return studyIDs, nil
}


func (c *Client) GetStudiesExpanded(params *types.StudiesQueryParams) ([]types.Study, error) {
	// Ensure expand is set
	if params == nil {
		params = &types.StudiesQueryParams{Expand: true}
	} else {
		params.Expand = true
	}

	path := c.buildStudiesPath("studies", params)

	var studies []types.Study
	if err := c.get(path, &studies); err != nil {
		return nil, err
	}

	return studies, nil
}


func (c *Client) GetStudy(studyID string) (*types.Study, error) {
	var study types.Study
	path := fmt.Sprintf("studies/%s", studyID)

	if err := c.get(path, &study); err != nil {
		return nil, err
	}

	return &study, nil
}

// buildStudiesPath constructs the path with query parameters
func (c *Client) buildStudiesPath(basePath string, params *types.StudiesQueryParams) string {
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

	if params.Limit > 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
	}

	if params.Short {
		queryParams += separator + "short"
		separator = "&"
	}

	if params.Full {
		queryParams += separator + "full"
		separator = "&"
	}

	return basePath + queryParams
}

func (c *Client) DeleteStudy(studyID string) error { 
	path := fmt.Sprintf("studies/%s", studyID)
	return c.delete(path, nil) 
}

// Only implemented the synchronous version, asynchronous version will be available later
func (c *Client) AnonymizeStudy(studyID string, anonymizeRequest *types.StudyAnonymizeRequest) (*types.StudyAnonymizeResponse, error) {
	var result types.StudyAnonymizeResponse
	path := fmt.Sprintf("studies/%s/anonymize", studyID)
	anonymizeRequest.Asynchronous = BoolPtr(false)

	if err := c.post(path, anonymizeRequest, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// getWithRawResponse performs a GET request and returns the raw response
func (c *Client) getWithRawResponse(path string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, path, nil)
}

// putWithPlainText performs a PUT request with a plain text body
func (c *Client) putWithPlainText(path string, body string) error {
	bodyReader := strings.NewReader(body)

	resp, err := c.doRequest(http.MethodPut, path, bodyReader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// postWithRawResponse performs a GET request and returns the raw response
func (c *Client) postWithBodyAndRawResponse(path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader

	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		bodyReader = strings.NewReader(string(bodyBytes))
	}

	return c.doRequest(http.MethodPost, path, bodyReader)
}

func (c *Client) DownloadStudyArchive(studyID string) (*http.Response, error) {
	path := fmt.Sprintf("studies/%s/archive", studyID)
	return c.getWithRawResponse(path)
}

func (c *Client) GetStudyStatistics(studyID string) (*types.Statistics, error) {
	var stats types.Statistics
	path := fmt.Sprintf("studies/%s/statistics", studyID)

	if err := c.get(path, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func (c *Client) GetStudySeries(studyID string) ([]string, error) {
	var seriesIDs []string
	path := fmt.Sprintf("studies/%s/series?expand=false", studyID)

	if err := c.get(path, &seriesIDs); err != nil {
		return nil, err
	}

	return seriesIDs, nil
}

func (c *Client) GetStudySeriesExpanded(studyID string) ([]types.Series, error) {
	var series []types.Series
	path := fmt.Sprintf("studies/%s/series?expand=true", studyID)

	if err := c.get(path, &series); err != nil {
		return nil, err
	}

	return series, nil
}

func (c *Client) GetStudyInstances(studyID string) ([]string, error) {
	var instanceIDs []string
	path := fmt.Sprintf("studies/%s/instances?expand=false", studyID)

	if err := c.get(path, &instanceIDs); err != nil {
		return nil, err
	}

	return instanceIDs, nil
}

func (c *Client) GetStudyInstancesExpanded(studyID string) ([]types.Instance, error) {
	var instances []types.Instance
	path := fmt.Sprintf("studies/%s/instances?expand=true", studyID)

	if err := c.get(path, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}
