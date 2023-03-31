package openaq

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type LocationBaseArgs struct {
	//
	BaseArgs *BaseArgs
	// Coordinates represents latitude,longitude (y,x) values to center the search
	// must be paired with Radius
	Coordinates *CoordinatesArgs
	// Radius represents the distance in meters from the Coordinates position
	// to search, must be paired with Coordinates
	Radius int32
	// Bbox is a bounding box to search within and is represented as a slice of four coordinates
	// in the form [ymin xmin ymax xmax]
	Bbox []float32
	// IsoCode is the ISO 3166-1 alpha-2 country code, unlike countries ID ISO code searching can only accept
	// a single value and not a slice of values
	IsoCode string
	// Monitor allows filtering for reference grade/regulatory monitors vs air sensors (AKA low-cost sensors)
	Monitor bool
	// Mobile allowing filtering for mobile locations vs stationary locations
	Mobile bool
	//
	OwnersIDs []int64
	//
	DatetimeStart time.Time
}

func (args *LocationBaseArgs) Values(q url.Values) (url.Values, error) {

	if (args.Coordinates != nil || args.Radius != 0) && len(args.Bbox) > 0 {
		return nil, &ClientError{}
	}

	if args.Coordinates != nil {
		lat := strconv.FormatFloat(args.Coordinates.Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(args.Coordinates.Lon, 'f', -1, 64)
		coords := fmt.Sprintf("%s,%s", lat, lon)
		q.Add("coordinates", coords)
	}

	if args.Radius != 0 {
		q.Add("radius", strconv.Itoa(int(args.Radius)))
	}

	if len(args.IsoCode) > 0 {
		q.Add("iso", args.IsoCode)
	}

	if !args.DatetimeStart.IsZero() {
		q.Add("datetime_start", args.DatetimeStart.Format(time.RFC3339))
	}

	return q, nil
}

type Countries struct {
	//
	IDs []int64
}

func (args *Countries) Values(q url.Values) url.Values {
	if len(args.IDs) == 0 {
		q.Add("countries_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args.IDs)), ","), "[]"))
	}
	return q
}

type Providers struct {
	//
	IDs []int64
}

func (args *Providers) Values(q url.Values) url.Values {
	if len(args.IDs) == 0 {
		q.Add("providers_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args.IDs)), ","), "[]"))
	}
	return q
}

type LocationArgs struct {
	//
	LocationBaseArgs
	//
	Countries *Countries
	//
	Providers *Providers
}

func (args *LocationArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := args.LocationBaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	q = args.Countries.Values(q)
	q = args.Providers.Values(q)
	return q, nil
}

type LocationsByCountryArgs struct {
	//
	LocationBaseArgs *LocationBaseArgs
	//
	Providers *Providers
}

func (args *LocationsByCountryArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := args.LocationBaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	q = args.Providers.Values(q)
	return q, nil
}

type LocationsByProviderArgs struct {
	//
	LocationBaseArgs *LocationBaseArgs
	//
	Countries *Countries
}

func (args *LocationsByProviderArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetLocations fetches all locations filtered by any params passed.
func (c *Client) GetLocations(ctx context.Context, args LocationArgs) (*LocationsResponse, error) {
	resp := &LocationsResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/locations", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetLocation fetches a single location by ID.
func (c *Client) GetLocation(ctx context.Context, locationsID int64) (*LocationsResponse, error) {
	path := fmt.Sprintf("/locations/%d", locationsID)
	resp := &LocationsResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetLocationsByCountry fetches a single location by ID.
func (c *Client) GetLocationByCountry(ctx context.Context, countriesID int64, args LocationsByCountryArgs) (*LocationsResponse, error) {
	path := fmt.Sprintf("/countries/%d/locations", countriesID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &LocationsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetLocations fetches all locations filtered by providerID and any params passed.
func (c *Client) GetLocationByProvider(ctx context.Context, providerID int64, args LocationsByProviderArgs) (*LocationsResponse, error) {
	path := fmt.Sprintf("/providers/%d/locations", providerID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &LocationsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
