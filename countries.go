package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type CountryArgs struct {
	//
	BaseArgs BaseArgs
	//
	Coordinates *Coordinates
}

func (args CountryArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetCountries fetches all countries filtered by any params passed.
func (c *Client) GetCountries(ctx context.Context, args CountryArgs) (*CountriesResponse, error) {
	resp := &CountriesResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/countries", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetCountry fetches a single country by ID.
func (c *Client) GetCountry(ctx context.Context, countriesID int64) (*CountriesResponse, error) {
	path := fmt.Sprintf("/countries/%d", countriesID)
	resp := &CountriesResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
