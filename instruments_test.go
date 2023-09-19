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

const instruments = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":1,"name":"OpenAQ admin","locationsCount":7061}]}`

func TestInstrumentsQueryParams(t *testing.T) {
	instrumentArgs := InstrumentArgs{
		BaseArgs: BaseArgs{
			Limit: 100,
			Page:  42,
		},
	}
	queryString, err := instrumentArgs.QueryParams()
	ok(t, err)
	expected := url.Values{"limit": []string{"100"}, "page": []string{"42"}}
	equals(t, expected, queryString)
}

func TestGetInstruments(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://api.openaq.org/v3/instruments?limit=100&page=1")
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
	args := &InstrumentArgs{}
	ctx := context.Background()
	body, err := openAQClient.GetInstruments(ctx, *args)
	expected := InstrumentsResponse{
		Meta: Meta{
			Name:    "openaq-api",
			Website: "/",
			Limit:   100,
			Page:    1,
			Found:   float64(1),
		},
		Results: []Instrument{
			{
				ID:             1,
				Name:           "OpenAQ admin",
				LocationsCount: 7061,
			},
		},
	}
	ok(t, err)
	equals(t, &expected, body)
}

func TestGetInstrument(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/instruments/1")
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
	body, err := openAQClient.GetInstrument(ctx, 1)
	ok(t, err)
	equals(t, body.Results[len(body.Results)-1].ID, int64(1))
}

func TestGetManufacturerInstruments(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://api.openaq.org/v3/manufacturers/42/instruments")
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
	ctx := context.Background()
	body, err := openAQClient.GetManufacturerInstruments(ctx, 42)
	expected := InstrumentsResponse{
		Meta: Meta{
			Name:    "openaq-api",
			Website: "/",
			Limit:   100,
			Page:    1,
			Found:   float64(1),
		},
		Results: []Instrument{
			{
				ID:             1,
				Name:           "OpenAQ admin",
				LocationsCount: 7061,
			},
		},
	}
	ok(t, err)
	equals(t, &expected, body)
}
