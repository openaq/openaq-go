package openaq

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

const countries = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":42,"code":"TG","name":"Togo","datetimeFirst":"2020-01-24T01:13:50+00:00","datetimeLast":"2023-08-19T01:00:59+00:00","parameters":[{"id":1,"name":"pm10","units":"µg/m³","displayName":null},{"id":2,"name":"pm25","units":"µg/m³","displayName":null},{"id":19,"name":"pm1","units":"µg/m³","displayName":null},{"id":100,"name":"temperature","units":"c","displayName":null},{"id":125,"name":"um003","units":"particles/cm³","displayName":null},{"id":126,"name":"um010","units":"particles/cm³","displayName":null},{"id":128,"name":"temperature","units":"f","displayName":null},{"id":129,"name":"um050","units":"particles/cm³","displayName":null},{"id":130,"name":"um025","units":"particles/cm³","displayName":null},{"id":132,"name":"pressure","units":"mb","displayName":null},{"id":133,"name":"um005","units":"particles/cm³","displayName":null},{"id":134,"name":"humidity","units":"%","displayName":null},{"id":135,"name":"um100","units":"particles/cm³","displayName":null}],"locationsCount":11,"measurementsCount":5641624,"providersCount":1}]}`

// func TestGetCountries(t *testing.T) {

// 	client := NewTestClient(func(req *http.Request) *http.Response {
// 		equals(t, req.URL.String(), "https://api.openaq.org/v3/countries")
// 		return &http.Response{
// 			StatusCode: 200,
// 			Body:       io.NopCloser(strings.NewReader(countries)),
// 			Header:     make(http.Header),
// 		}
// 	})
// 	config := &Config{
// 		Client: client,
// 	}
// 	openAQClient, err := NewClient(*config)
// 	ok(t, err)

// 	countryArgs := &CountryArgs{}
// 	ctx := context.Background()
// 	body, err := openAQClient.GetCountries(ctx, *countryArgs)
// 	ok(t, err)
// 	var results []Country
// 	results = append(results, Country{})
// 	equals(t, &CountriesResponse{
// 		Meta: Meta{
// 			Name:    "openaq-api",
// 			Website: "/",
// 			Page:    1,
// 			Limit:   100,
// 			Found:   1,
// 		},
// 		Results: results,
// 	}, body)
// }

func TestGetCountry(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/countries/42")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(countries)),
			Header:     make(http.Header),
		}
	})
	config := &Config{
		Client: client,
	}
	openAQClient, err := NewClient(*config)
	ok(t, err)

	ctx := context.Background()
	body, err := openAQClient.GetCountry(ctx, 42)
	ok(t, err)

	equals(t, body.Results[len(body.Results)-1].ID, int64(42))
}
