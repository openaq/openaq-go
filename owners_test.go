package openaq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const owners = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":1,"name":"OpenAQ admin","locationsCount":7061}]}`

func TestGetOwners(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://api.openaq.org/v3/owners")
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
		fmt.Println("Failed to create new client")
	}
	args := &OwnersArgs{}
	ctx := context.Background()
	body, err := openAQClient.GetOwners(ctx, *args)
	expected := OwnersResponse{
		Meta: Meta{
			Name:    "openaq-api",
			Website: "/",
			Limit:   100,
			Page:    1,
			Found:   float64(1),
		},
		Results: []Owner{
			{
				ID:             1,
				Name:           "OpenAQ admin",
				LocationsCount: 7061,
			},
		},
	}
	ok(t, err)
	fmt.Println(cmp.Diff(&expected, body))
	equals(t, &expected, body)
}

func TestGetOwner(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/owners/1")
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
	body, err := openAQClient.GetOwner(ctx, 1)
	ok(t, err)

	equals(t, body.Results[len(body.Results)-1].ID, int64(1))
}
