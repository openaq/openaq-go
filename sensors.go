package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type SensorArgs struct {
	BaseArgs BaseArgs
}

// QueryParams translates SensorArgs struct into url.Values
func (args SensorArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetSensors fetches all sensors of a given location
func (c *Client) GetLocationSensors(ctx context.Context, locationsID int64) (*SensorsResponse, error) {
	path := fmt.Sprintf("locations/%d/sensors", locationsID)
	resp := &SensorsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSensor fetches a single Sensor by ID.
func (c *Client) GetSensor(ctx context.Context, SensorsID int64) (*SensorsResponse, error) {
	path := fmt.Sprintf("/sensors/%d", SensorsID)
	resp := &SensorsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
