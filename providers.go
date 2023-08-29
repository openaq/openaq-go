package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type ProvidersArgs struct {
	BaseArgs BaseArgs
}

// QueryParams translates ProvidersArgs struct into url.Values
func (args *ProvidersArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetProviders fetches all providers filtered by any params passed.
func (c *Client) GetProviders(ctx context.Context, args ProvidersArgs) (*ProvidersResponse, error) {
	resp := &ProvidersResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/providers", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetProvider fetches a single provider by ID.
func (c *Client) GetProvider(ctx context.Context, providersID int64) (*ProvidersResponse, error) {
	path := fmt.Sprintf("/providers/%d", providersID)
	resp := &ProvidersResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
