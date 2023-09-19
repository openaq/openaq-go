package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type InstrumentArgs struct {
	BaseArgs BaseArgs
}

// QueryParams translates InstrumentArgs struct into url.Values
func (instrumentArgs InstrumentArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := instrumentArgs.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// GetInstruments fetches all instruments filtered by any params passed.
func (c *Client) GetInstruments(ctx context.Context, args InstrumentArgs) (*InstrumentsResponse, error) {
	resp := &InstrumentsResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/instruments", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetManufacturerInstruments fetches instruments by manufacturer ID.
func (c *Client) GetManufacturerInstruments(ctx context.Context, manufacturerID int64) (*InstrumentsResponse, error) {
	path := fmt.Sprintf("/manufacturers/%d/instruments", manufacturerID)
	resp := &InstrumentsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetInstrument fetches a single instrument by ID.
func (c *Client) GetInstrument(ctx context.Context, instrumentsID int64) (*InstrumentsResponse, error) {
	path := fmt.Sprintf("/instruments/%d", instrumentsID)
	resp := &InstrumentsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
