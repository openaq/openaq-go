package openaq

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type LocationArgs struct {
	//
	BaseArgs *BaseArgs
	// Coordinates represents latitude,longitude (y,x) values to center the search
	// must be paired with Radius
	Coordinates *CoordinatesArgs
	// Radius represents the distance in meters from the Coordinates position
	// to search, must be paired with Coordinates
	Radius int64
	// Bbox is a bounding box to search within and is represented as a slice of four coordinates
	// in the form [ymin xmin ymax xmax]
	Bbox []float64
	// IsoCode is the ISO 3166-1 alpha-2 country code, unlike countries ID ISO code searching can only accept
	// a single value and not a slice of values
	IsoCode string
	// Monitor allows filtering for reference grade/regulatory monitors vs air sensors (AKA low-cost sensors)
	Monitor *bool
	// Mobile allowing filtering for mobile locations vs stationary locations
	Mobile bool
	// A slice of owners IDs
	OwnersIDs []int64
	//
	Countries *Countries
	//
	Providers *Providers
}

func (locationArgs *LocationArgs) Values(q url.Values) (url.Values, error) {
	if locationArgs.Coordinates != nil {
		lat := strconv.FormatFloat(locationArgs.Coordinates.Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(locationArgs.Coordinates.Lon, 'f', -1, 64)
		coords := fmt.Sprintf("%s,%s", lat, lon)
		q.Add("coordinates", coords)
	}
	if locationArgs.Countries != nil {
		q = locationArgs.Countries.Values(q)
	}
	if locationArgs.Providers != nil {
		q = locationArgs.Providers.Values(q)
	}
	if locationArgs.Radius != 0 {
		q.Add("radius", strconv.Itoa(int(locationArgs.Radius)))
	}
	if locationArgs.IsoCode != "" {
		q.Add("iso", locationArgs.IsoCode)
	}
	if locationArgs.Monitor != nil {
		q.Add("monitor", strconv.FormatBool(*locationArgs.Monitor))
	}
	if locationArgs.Bbox != nil {
		minX := strconv.FormatFloat(locationArgs.Bbox[0], 'f', -1, 64)
		minY := strconv.FormatFloat(locationArgs.Bbox[1], 'f', -1, 64)
		maxX := strconv.FormatFloat(locationArgs.Bbox[2], 'f', -1, 64)
		maxY := strconv.FormatFloat(locationArgs.Bbox[3], 'f', -1, 64)
		bboxValue := fmt.Sprintf("%s,%s,%s,%s", minX, minY, maxX, maxY)
		q.Add("bbox", bboxValue)
	}
	return q, nil
}

// QueryParams translates LocationArgs struct into url.Values
func (locationArgs LocationArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := locationArgs.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	q, err = locationArgs.Values(q)
	if err != nil {
		return nil, err
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
