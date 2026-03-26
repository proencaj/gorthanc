package gorthanc

import "github.com/proencaj/gorthanc/types"

func (c *Client) GetSystem() (*types.SystemInfo, error) {
	var info types.SystemInfo
	if err := c.get("system", &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *Client) GetSystemStatistics() (*types.SystemStatistics, error) {
	var statistics types.SystemStatistics
	if err := c.get("statistics", &statistics); err != nil {
		return nil, err
	}
	return &statistics, nil
}