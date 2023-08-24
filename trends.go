package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type TrendsArgs struct {
	BaseArgs BaseArgs
}

// QueryParams translates TrendsArgs struct into url.Values
func (args TrendsArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetTrends fetches trends of a given location and parameter ID
func (c *Client) GetTrends(ctx context.Context, locationsID int64, parameterID int64) (*TrendsResponse, error) {
	path := fmt.Sprintf("locations/%d/sensors", locationsID)
	resp := &TrendsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
