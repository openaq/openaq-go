package openaq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

const parameters = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":11,"name":"bc","units":"µg/m³","displayName":"BC","description":"Black Carbon mass concentration","locationsCount":93,"measurementsCount":734}]}`

func TestGetParameters(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://api.openaq.org/v3/parameters")
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
		fmt.Println("Failed to create new client")
	}
	args := &ParametersArgs{}
	ctx := context.Background()
	body, err := openAQClient.GetParameters(ctx, *args)
	ok(t, err)
	var results []Parameter
	results = append(results, Parameter{
		ID:                11,
		Name:              "bc",
		Units:             "µg/m³",
		DisplayName:       "BC",
		Description:       "Black Carbon mass concentration",
		LocationsCount:    93,
		MeasurementsCount: 734,
	})
	equals(t, &ParametersResponse{
		Meta: Meta{
			Name:    "openaq-api",
			Website: "/",
			Page:    1,
			Limit:   100,
			Found:   1,
		},
		Results: results,
	}, body)
}

func TestGetParameter(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/parameters/11")
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
	body, err := openAQClient.GetParameter(ctx, 11)
	ok(t, err)

	equals(t, body.Results[len(body.Results)-1].ID, int64(11))
}
