package openaq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const manufacturers = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":1,"name":"OpenAQ admin","locationsCount":7061}]}`

func TestManufacturersQueryParams(t *testing.T) {
	manufacturerArgs := ManufacturerArgs{
		BaseArgs: BaseArgs{
			Limit: 100,
			Page:  42,
		},
	}
	queryString, err := manufacturerArgs.QueryParams()
	ok(t, err)
	expected := url.Values{"limit": []string{"100"}, "page": []string{"42"}}
	equals(t, expected, queryString)
}

func TestGetManufacturers(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://api.openaq.org/v3/manufacturers?limit=100&page=1")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(manufacturers)),
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
	args := &ManufacturerArgs{}
	ctx := context.Background()
	body, err := openAQClient.GetManufacturers(ctx, *args)
	expected := ManufacturersResponse{
		Meta: Meta{
			Name:    "openaq-api",
			Website: "/",
			Limit:   100,
			Page:    1,
			Found:   float64(1),
		},
		Results: []Manufacturer{
			{
				ID:   1,
				Name: "OpenAQ admin",
			},
		},
	}
	ok(t, err)
	equals(t, &expected, body)
}

func TestGetManufacturer(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/manufacturers/1")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(owners)),
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
	body, err := openAQClient.GetManufacturer(ctx, 1)
	ok(t, err)
	equals(t, body.Results[len(body.Results)-1].ID, int64(1))
}
