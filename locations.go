package openaq

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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
}

func (args *LocationBaseArgs) Values(q url.Values) (url.Values, error) {

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

	return q, nil
}

type Countries struct {
	//
	IDs []int64
}

func (args *Countries) Values(q url.Values) url.Values {
	if args != nil {
		q.Add("countries_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args.IDs)), ","), "[]"))
	}
	return q
}

type Providers struct {
	//
	IDs []int64
}

func (args *Providers) Values(q url.Values) url.Values {
	if args != nil {
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
	if args.Countries != nil {
		q = args.Countries.Values(q)
	}
	if args.Providers != nil {
		q = args.Providers.Values(q)
	}
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
