package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type ManufacturerArgs struct {
	BaseArgs BaseArgs
}

// QueryParams translates ManufacturerArgs struct into url.Values
func (manufacturerArgs ManufacturerArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := manufacturerArgs.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// GetManufacturers fetches all manufacturers filtered by any params passed.
func (c *Client) GetManufacturers(ctx context.Context, args ManufacturerArgs) (*ManufacturersResponse, error) {
	resp := &ManufacturersResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/manufacturers", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetManufacturer fetches a single manufacturer by ID.
func (c *Client) GetManufacturer(ctx context.Context, manufacturersID int64) (*ManufacturersResponse, error) {
	path := fmt.Sprintf("/manufacturers/%d", manufacturersID)
	resp := &ManufacturersResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
