package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type CountryArgs struct {
	//
	BaseArgs BaseArgs
}

func (countryArgs CountryArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := countryArgs.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// GetCountries fetches all countries filtered by any params passed.
func (c *Client) GetCountries(ctx context.Context, countryArgs CountryArgs) (*CountriesResponse, error) {
	resp := &CountriesResponse{}
	queryParams, err := countryArgs.QueryParams()
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
