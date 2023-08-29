package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type ParametersArgs struct {
	BaseArgs      BaseArgs
	ParameterType string
}

func (parametersArgs *ParametersArgs) Values(q url.Values) (url.Values, error) {
	if parametersArgs.ParameterType != "" {
		q.Add("parameter_type", parametersArgs.ParameterType)
	}
	return q, nil
}

// QueryParams translates ParametersArgs struct into url.Values
func (parametersArgs ParametersArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := parametersArgs.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	q, err = parametersArgs.Values(q)
	if err != nil {
		return nil, err
	}
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
