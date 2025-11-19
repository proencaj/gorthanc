package gorthanc

import "github.com/proencaj/gorthanc/types"

func (c *Client) GetSystem() (*types.SystemInfo, error) {
	var info types.SystemInfo
	if err := c.get("system", &info); err != nil {
		return nil, err
	}
	return &info, nil
}