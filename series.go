package gorthanc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/proencaj/gorthanc/types"
)

func (c *Client) GetSeries(params *types.SeriesQueryParams) ([]string, error) {
	path := "series"

	if params != nil {
		path = c.buildSeriesPath(path, params)
	}

	// If expand is requested, we need to handle the response differently
	if params != nil && params.Expand {
		return nil, fmt.Errorf("for expanded results, use GetSeriesExpanded method")
	}

	var seriesIDs []string
	if err := c.get(path, &seriesIDs); err != nil {
		return nil, err
	}

	return seriesIDs, nil
}


func (c *Client) GetSeriesExpanded(params *types.SeriesQueryParams) ([]types.Series, error) {
	// Ensure expand is set
	if params == nil {
		params = &types.SeriesQueryParams{Expand: true}
	} else {
		params.Expand = true
	}

	path := c.buildSeriesPath("series", params)

	var seriesList []types.Series
	if err := c.get(path, &seriesList); err != nil {
		return nil, err
	}

	return seriesList, nil
}


func (c *Client) GetSeriesDetail(seriesID string) (*types.Series, error) {
	var series types.Series
	path := fmt.Sprintf("series/%s", seriesID)

	if err := c.get(path, &series); err != nil {
		return nil, err
	}

	return &series, nil
}


func (c *Client) DeleteSeries(seriesID string) error {
	path := fmt.Sprintf("series/%s", seriesID)
	return c.delete(path, nil) 
}


// Only implemented the synchronous version, asynchronous version will be available later
func (c *Client) AnonymizeSeries(seriesID string, anonymizeRequest *types.SeriesAnonymizeRequest) (*types.SeriesAnonymizeResponse, error) {
	var result types.SeriesAnonymizeResponse
	path := fmt.Sprintf("series/%s/anonymize", seriesID)
	anonymizeRequest.Asynchronous = BoolPtr(false)

	if err := c.post(path, anonymizeRequest, &result); err != nil {
		return nil, err
	}

	return &result, nil
}


func (c *Client) DownloadSeriesArchive(seriesID string) (*http.Response, error) {
	path := fmt.Sprintf("series/%s/archive", seriesID)
	return c.getWithRawResponse(path)
}


func (c *Client) GetSeriesStatistics(seriesID string) (*types.Statistics, error) {
	var stats types.Statistics
	path := fmt.Sprintf("series/%s/statistics", seriesID)

	if err := c.get(path, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}


func (c *Client) GetSeriesInstances(seriesID string) ([]string, error) {
	var instanceIDs []string
	path := fmt.Sprintf("series/%s/instances?expand=false", seriesID)

	if err := c.get(path, &instanceIDs); err != nil {
		return nil, err
	}

	return instanceIDs, nil
}

func (c *Client) GetSeriesInstancesExpanded(seriesID string) ([]types.Instance, error) {
	var instances []types.Instance
	path := fmt.Sprintf("series/%s/instances?expand=true", seriesID)

	if err := c.get(path, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}



// buildSeriesPath constructs the path with query parameters
func (c *Client) buildSeriesPath(basePath string, params *types.SeriesQueryParams) string {
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

	return basePath + queryParams
}
