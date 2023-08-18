package openaq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestMeasurementsArgQueryParams(t *testing.T) {
	baseArgs := BaseArgs{}
	measurementArgs := MeasurementsArgs{
		BaseArgs:     baseArgs,
		DatetimeFrom: time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC),
		DatetimeTo:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC),
		Parameters:   []int64{1, 2, 3},
		PeriodName:   "hour",
	}
	queryString, err := measurementArgs.QueryParams()
	if err != nil {
		fmt.Println(err)
	}
	equals(t, url.Values{"date_from": []string{"2021-08-15T14:30:45Z"}, "date_to": []string{"2021-08-15T14:30:45Z"}, "limit": []string{"100"}, "page": []string{"1"}, "period_name": []string{"hour"}}, queryString)
}

func TestGetLocationMeasurments(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "https://api.openaq.org/v3/locations/2178/measurements?limit=1&page=1")
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
		fmt.Println(err)
	}
	ctx := context.Background()
	openAQClient.GetLocationMeasurements(ctx, 2178, MeasurementsArgs{BaseArgs: BaseArgs{Limit: 1, Page: 1}})
	ok(t, err)

	//equals(t, body.Results[len(body.Results)-1].ID, int64(11))
}
