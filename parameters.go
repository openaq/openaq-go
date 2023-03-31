package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type ParametersArgs struct {
	BaseArgs BaseArgs

	Coordinates *Coordinates
}

// QueryParams translates ParametersArgs struct into url.Values
func (args ParametersArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetParameters fetches all parameters filtered by any params passed.
func (c *Client) GetParameters(ctx context.Context, args ParametersArgs) (*ParametersResponse, error) {
	resp := &ParametersResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/parameters", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetParameter fetches a single parameter by ID.
func (c *Client) GetParameter(ctx context.Context, parametersID int64) (*ParametersResponse, error) {
	path := fmt.Sprintf("/parameters/%d", parametersID)
	resp := &ParametersResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
