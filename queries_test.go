package openaq

import (
	"net/url"
	"testing"
)

func TestParametersValues(t *testing.T) {
	parameters := Parameters{IDs: []int64{1, 2, 3}}
	q := make(url.Values)
	q = parameters.Values(q)
	equals(t, url.Values{"parameters_id": []string{"1,2,3"}}, q)
}

func TestCountriesValues(t *testing.T) {
	countries := Countries{IDs: []int64{1, 2, 3}}
	q := make(url.Values)
	q = countries.Values(q)
	equals(t, url.Values{"countries_id": []string{"1,2,3"}}, q)
}

func TestProvidersValues(t *testing.T) {
	providers := Providers{IDs: []int64{1, 2, 3}}
	q := make(url.Values)
	q = providers.Values(q)
	equals(t, url.Values{"providers_id": []string{"1,2,3"}}, q)
}
