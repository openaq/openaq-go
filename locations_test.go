package openaq

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"
)

func TestLocationsArgsQueryParams(t *testing.T) {
	args := LocationArgs{
		LocationBaseArgs{
			BaseArgs: &BaseArgs{
				Limit: 100,
				Page:  1,
			},
			Radius:      10,
			Coordinates: &CoordinatesArgs{32.3, 32.3},
		},
		&Countries{},
		&Providers{},
	}
	queryParams, err := args.QueryParams()
	ok(t, err)
	expectedParams := &url.Values{}
	expectedParams.Add("radius", "10")
	expectedParams.Add("countries_id", "")
	expectedParams.Add("providers_id", "")
	coords := fmt.Sprintf("%s,%s", strconv.FormatFloat(32.3, 'f', -1, 64), strconv.FormatFloat(32.3, 'f', -1, 64))
	expectedParams.Add("coordinates", coords)
	equals(t, expectedParams, &queryParams)
}

func TestGetLocations(t *testing.T) {

}

func TestGetLocation(t *testing.T) {

}
