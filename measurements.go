package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type MeasurementsArgs struct {
	BaseArgs BaseArgs

	Coordinates *Coordinates
}

// QueryParams translates MeasurementsArgs struct into url.Values
func (args MeasurementsArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetLocationMeasurements fetches measurements for a single location by ID and any params passed
func (c *Client) GetLocationMeasurements(ctx context.Context, locationsID int64, args MeasurementsArgs) (*MeasurementsResponse, error) {
	path := fmt.Sprintf("/locations/%d/measurements", locationsID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &MeasurementsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSensorMeasurements fetches measurements for a single sensor by ID and any params passed
func (c *Client) GetSensorMeasurements(ctx context.Context, sensorsID int64, args MeasurementsArgs) (*MeasurementsResponse, error) {
	path := fmt.Sprintf("/sensors/%d/measurements", sensorsID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &MeasurementsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
