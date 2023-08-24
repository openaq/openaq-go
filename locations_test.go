package openaq

import (
	"net/url"
	"testing"
)

const locations = `{"meta":{"name":"openaq-api","website":"/","page":1,"limit":100,"found":1},"results":[{"id":2178,"name":"Del Norte","locality":"Albuquerque","timezone":"America/Denver","country":{"id":13,"code":"US","name":"United States of America"},"owner":{"id":4,"name":"Unknown Governmental Organization"},"provider":{"id":119,"name":"AirNow"},"isMobile":false,"isMonitor":true,"instruments":[{"id":2,"name":"Government Monitor"}],"sensors":[{"id":25227,"name":"co ppm","parameter":{"id":8,"name":"co","units":"ppm","displayName":"CO"}},{"id":3916,"name":"no2 ppm","parameter":{"id":7,"name":"no2","units":"ppm","displayName":"NO₂"}},{"id":4272226,"name":"no ppm","parameter":{"id":35,"name":"no","units":"ppm","displayName":"NO"}},{"id":3917,"name":"o3 ppm","parameter":{"id":10,"name":"o3","units":"ppm","displayName":"O₃"}},{"id":3918,"name":"so2 ppm","parameter":{"id":9,"name":"so2","units":"ppm","displayName":"SO₂"}},{"id":3919,"name":"pm10 µg/m³","parameter":{"id":1,"name":"pm10","units":"µg/m³","displayName":"PM10"}},{"id":4272103,"name":"nox ppm","parameter":{"id":19840,"name":"nox","units":"ppm","displayName":"NOx"}},{"id":3920,"name":"pm25 µg/m³","parameter":{"id":2,"name":"pm25","units":"µg/m³","displayName":"PM2.5"}}],"coordinates":{"latitude":35.1353,"longitude":-106.584702},"bounds":[-106.584702,35.1353,-106.584702,35.1353],"distance":null,"datetimeFirst":{"utc":"2016-03-06T19:00:00+00:00","local":"2016-03-06T12:00:00-07:00"},"datetimeLast":{"utc":"2023-08-16T18:00:00+00:00","local":"2023-08-16T12:00:00-06:00"}}]}`

func TestLocationArgsQueryParams(t *testing.T) {
	baseArgs := BaseArgs{
		Limit: 1000,
		Page:  1,
	}
	locationBaseArgs := LocationBaseArgs{
		BaseArgs: &baseArgs,
		Monitor:  true,
		Mobile:   false,
	}
	providers := Providers{IDs: []int64{1, 2}}
	countries := Countries{IDs: []int64{6, 7}}

	args := LocationArgs{
		LocationBaseArgs: locationBaseArgs,
		Providers:        &providers,
		Countries:        &countries,
	}
	queryParams, _ := args.QueryParams()
	expected := url.Values{"countries_id": []string{"6,7"}, "providers_id": []string{"1,2"}}
	equals(t, expected, queryParams)
}

// func TestGetLocation(t *testing.T) {

// 	client := NewTestClient(func(req *http.Request) *http.Response {
// 		// Test request parameters
// 		equals(t, req.URL.String(), "https://api.openaq.org/v3/locations/2178")
// 		return &http.Response{
// 			StatusCode: 200,
// 			Body:       io.NopCloser(strings.NewReader(locations)),
// 			Header:     make(http.Header),
// 		}
// 	})
// 	config := &Config{
// 		Client: client,
// 	}
// 	openAQClient, err := NewClient(*config)
// 	ok(t, err)
// 	ctx := context.Background()
// 	body, err := openAQClient.GetLocation(ctx, 2178)
// 	ok(t, err)
// 	var locationsResponse LocationsResponse
// 	err = json.Unmarshal([]byte(locations), &locationsResponse)
// 	ok(t, err)
// 	equals(t, body, locationsResponse)
// 	equals(t, body.Results[len(body.Results)-1].ID, int64(2178))
// }

// func TestGetLocations(t *testing.T) {

// 	client := NewTestClient(func(req *http.Request) *http.Response {
// 		// Test request parameters
// 		equals(t, req.URL.String(), "https://api.openaq.org/v3/locations")
// 		return &http.Response{
// 			StatusCode: 200,
// 			Body:       io.NopCloser(strings.NewReader(locations)),
// 			Header:     make(http.Header),
// 		}
// 	})
// 	config := &Config{
// 		Client: client,
// 	}
// 	openAQClient, err := NewClient(*config)
// 	ok(t, err)
// 	ctx := context.Background()
// 	body, err := openAQClient.GetLocations(ctx, LocationArgs{})
// 	ok(t, err)
// 	var locationsResponse LocationsResponse
// 	err = json.Unmarshal([]byte(locations), &locationsResponse)
// 	ok(t, err)
// 	equals(t, body, locationsResponse)
// 	equals(t, body.Results[len(body.Results)-1].ID, int64(2178))

// }
