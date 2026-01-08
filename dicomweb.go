package gorthanc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/proencaj/gorthanc/types"
)

func (c *Client) QidoSearchStudies(params *types.QidoStudyQueryParams) ([]map[string]interface{}, error) {
	path := "dicom-web/studies"

	if params != nil {
		path = c.buildQidoStudiesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) QidoSearchSeries(studyUID string, params *types.QidoSeriesQueryParams) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series", studyUID)

	if params != nil {
		path = c.buildQidoSeriesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) QidoSearchAllSeries(params *types.QidoSeriesQueryParams) ([]map[string]interface{}, error) {
	path := "dicom-web/series"

	if params != nil {
		path = c.buildQidoSeriesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) QidoSearchInstances(studyUID string, seriesUID string, params *types.QidoInstanceQueryParams) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances", studyUID, seriesUID)

	if params != nil {
		path = c.buildQidoInstancesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) QidoSearchStudyInstances(studyUID string, params *types.QidoInstanceQueryParams) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/instances", studyUID)

	if params != nil {
		path = c.buildQidoInstancesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) QidoSearchAllInstances(params *types.QidoInstanceQueryParams) ([]map[string]interface{}, error) {
	path := "dicom-web/instances"

	if params != nil {
		path = c.buildQidoInstancesPath(path, params)
	}

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) WadoRsRetrieveStudy(studyUID string) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s", studyUID)
	return c.getWithAcceptRawResponse(path, "multipart/related; type=application/dicom")
}

func (c *Client) WadoRsRetrieveSeries(studyUID string, seriesUID string) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s", studyUID, seriesUID)
	return c.getWithAcceptRawResponse(path, "multipart/related; type=application/dicom")
}

func (c *Client) WadoRsRetrieveInstance(studyUID string, seriesUID string, instanceUID string) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances/%s", studyUID, seriesUID, instanceUID)
	return c.getWithAcceptRawResponse(path, "multipart/related; type=application/dicom")
}

func (c *Client) WadoRsRetrieveStudyMetadata(studyUID string) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/metadata", studyUID)

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) WadoRsRetrieveSeriesMetadata(studyUID string, seriesUID string) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/metadata", studyUID, seriesUID)

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) WadoRsRetrieveInstanceMetadata(studyUID string, seriesUID string, instanceUID string) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances/%s/metadata", studyUID, seriesUID, instanceUID)

	var results []map[string]interface{}
	if err := c.get(path, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) WadoRsRetrieveFrames(studyUID string, seriesUID string, instanceUID string, frameList string) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances/%s/frames/%s", studyUID, seriesUID, instanceUID, frameList)
	return c.getWithAcceptRawResponse(path, "multipart/related; type=application/octet-stream")
}

func (c *Client) WadoRsRetrieveRenderedInstance(studyUID string, seriesUID string, instanceUID string, params *types.WadoRsRenderedParams) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances/%s/rendered", studyUID, seriesUID, instanceUID)

	accept := "image/jpeg"
	if params != nil {
		path = c.buildWadoRsRenderedPath(path, params)
		if params.Accept != "" {
			accept = params.Accept
		}
	}

	return c.getWithAcceptRawResponse(path, accept)
}

func (c *Client) WadoRsRetrieveRenderedFrames(studyUID string, seriesUID string, instanceUID string, frameList string, params *types.WadoRsRenderedParams) (*http.Response, error) {
	path := fmt.Sprintf("dicom-web/studies/%s/series/%s/instances/%s/frames/%s/rendered", studyUID, seriesUID, instanceUID, frameList)

	accept := "image/jpeg"
	if params != nil {
		path = c.buildWadoRsRenderedPath(path, params)
		if params.Accept != "" {
			accept = params.Accept
		}
	}

	return c.getWithAcceptRawResponse(path, accept)
}

func (c *Client) WadoUriRetrieve(params *types.WadoUriParams) (*http.Response, error) {
	if params == nil {
		return nil, fmt.Errorf("WadoUriParams is required")
	}

	path := c.buildWadoUriPath("wado", params)
	return c.getWithRawResponse(path)
}

func (c *Client) buildQidoStudiesPath(basePath string, params *types.QidoStudyQueryParams) string {
	if params == nil {
		return basePath
	}

	queryParams := ""
	separator := "?"

	if params.Limit > 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
		separator = "&"
	}

	if params.Offset > 0 {
		queryParams += separator + "offset=" + strconv.Itoa(params.Offset)
		separator = "&"
	}

	if params.Includefield != "" {
		queryParams += separator + "includefield=" + params.Includefield
		separator = "&"
	}

	if params.FuzzyMatching {
		queryParams += separator + "fuzzymatching=true"
		separator = "&"
	}

	if params.StudyInstanceUID != "" {
		queryParams += separator + "StudyInstanceUID=" + params.StudyInstanceUID
		separator = "&"
	}

	if params.PatientID != "" {
		queryParams += separator + "PatientID=" + params.PatientID
		separator = "&"
	}

	if params.PatientName != "" {
		queryParams += separator + "PatientName=" + params.PatientName
		separator = "&"
	}

	if params.AccessionNumber != "" {
		queryParams += separator + "AccessionNumber=" + params.AccessionNumber
		separator = "&"
	}

	if params.StudyDate != "" {
		queryParams += separator + "StudyDate=" + params.StudyDate
		separator = "&"
	}

	if params.ModalitiesInStudy != "" {
		queryParams += separator + "ModalitiesInStudy=" + params.ModalitiesInStudy
		separator = "&"
	}

	return basePath + queryParams
}

func (c *Client) buildQidoSeriesPath(basePath string, params *types.QidoSeriesQueryParams) string {
	if params == nil {
		return basePath
	}

	queryParams := ""
	separator := "?"

	if params.Limit > 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
		separator = "&"
	}

	if params.Offset > 0 {
		queryParams += separator + "offset=" + strconv.Itoa(params.Offset)
		separator = "&"
	}

	if params.Includefield != "" {
		queryParams += separator + "includefield=" + params.Includefield
		separator = "&"
	}

	if params.FuzzyMatching {
		queryParams += separator + "fuzzymatching=true"
		separator = "&"
	}

	if params.StudyInstanceUID != "" {
		queryParams += separator + "StudyInstanceUID=" + params.StudyInstanceUID
		separator = "&"
	}

	if params.SeriesInstanceUID != "" {
		queryParams += separator + "SeriesInstanceUID=" + params.SeriesInstanceUID
		separator = "&"
	}

	if params.Modality != "" {
		queryParams += separator + "Modality=" + params.Modality
		separator = "&"
	}

	if params.SeriesNumber != "" {
		queryParams += separator + "SeriesNumber=" + params.SeriesNumber
		separator = "&"
	}

	return basePath + queryParams
}

func (c *Client) buildQidoInstancesPath(basePath string, params *types.QidoInstanceQueryParams) string {
	if params == nil {
		return basePath
	}

	queryParams := ""
	separator := "?"

	if params.Limit > 0 {
		queryParams += separator + "limit=" + strconv.Itoa(params.Limit)
		separator = "&"
	}

	if params.Offset > 0 {
		queryParams += separator + "offset=" + strconv.Itoa(params.Offset)
		separator = "&"
	}

	if params.Includefield != "" {
		queryParams += separator + "includefield=" + params.Includefield
		separator = "&"
	}

	if params.FuzzyMatching {
		queryParams += separator + "fuzzymatching=true"
		separator = "&"
	}

	if params.StudyInstanceUID != "" {
		queryParams += separator + "StudyInstanceUID=" + params.StudyInstanceUID
		separator = "&"
	}

	if params.SeriesInstanceUID != "" {
		queryParams += separator + "SeriesInstanceUID=" + params.SeriesInstanceUID
		separator = "&"
	}

	if params.SOPInstanceUID != "" {
		queryParams += separator + "SOPInstanceUID=" + params.SOPInstanceUID
		separator = "&"
	}

	if params.SOPClassUID != "" {
		queryParams += separator + "SOPClassUID=" + params.SOPClassUID
		separator = "&"
	}

	return basePath + queryParams
}

func (c *Client) buildWadoRsRenderedPath(basePath string, params *types.WadoRsRenderedParams) string {
	if params == nil {
		return basePath
	}

	queryParams := ""
	separator := "?"

	if params.WindowCenter != "" {
		queryParams += separator + "window-center=" + params.WindowCenter
		separator = "&"
	}

	if params.WindowWidth != "" {
		queryParams += separator + "window-width=" + params.WindowWidth
		separator = "&"
	}

	if params.Quality > 0 {
		queryParams += separator + "quality=" + strconv.Itoa(params.Quality)
		separator = "&"
	}

	if params.Viewport != "" {
		queryParams += separator + "viewport=" + params.Viewport
		separator = "&"
	}

	return basePath + queryParams
}

func (c *Client) buildWadoUriPath(basePath string, params *types.WadoUriParams) string {
	if params == nil {
		return basePath
	}

	queryParams := "?"

	requestType := params.RequestType
	if requestType == "" {
		requestType = "WADO"
	}
	queryParams += "requestType=" + requestType

	if params.StudyUID != "" {
		queryParams += "&studyUID=" + params.StudyUID
	}

	if params.SeriesUID != "" {
		queryParams += "&seriesUID=" + params.SeriesUID
	}

	if params.ObjectUID != "" {
		queryParams += "&objectUID=" + params.ObjectUID
	}

	if params.ContentType != "" {
		queryParams += "&contentType=" + params.ContentType
	}

	if params.TransferSyntax != "" {
		queryParams += "&transferSyntax=" + params.TransferSyntax
	}

	if params.Anonymize != "" {
		queryParams += "&anonymize=" + params.Anonymize
	}

	if params.FrameNumber > 0 {
		queryParams += "&frameNumber=" + strconv.Itoa(params.FrameNumber)
	}

	if params.ImageQuality > 0 {
		queryParams += "&imageQuality=" + strconv.Itoa(params.ImageQuality)
	}

	if params.WindowCenter != "" {
		queryParams += "&windowCenter=" + params.WindowCenter
	}

	if params.WindowWidth != "" {
		queryParams += "&windowWidth=" + params.WindowWidth
	}

	if params.Rows > 0 {
		queryParams += "&rows=" + strconv.Itoa(params.Rows)
	}

	if params.Columns > 0 {
		queryParams += "&columns=" + strconv.Itoa(params.Columns)
	}

	if params.Region != "" {
		queryParams += "&region=" + params.Region
	}

	return basePath + queryParams
}
