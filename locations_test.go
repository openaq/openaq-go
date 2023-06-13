package openaq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/locations/11")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(parameters)),
			Header:     make(http.Header),
		}
	})
	config := &Config{
		Client: client,
	}
	openAQClient, err := NewClient(*config)
	if err != nil {
		fmt.Println("")
	}
	ctx := context.Background()
	body, err := openAQClient.GetLocation(ctx, 11)
	ok(t, err)

	equals(t, body.Results[len(body.Results)-1].ID, int64(11))
}
