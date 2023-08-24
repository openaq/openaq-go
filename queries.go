package openaq

import (
	"fmt"
	"net/url"
	"strings"
)

type Parameters struct {
	//
	IDs []int64
}

func (parameters *Parameters) Values(q url.Values) url.Values {
	if parameters != nil {
		q.Add("parameters_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(parameters.IDs)), ","), "[]"))
	}
	return q
}

type Countries struct {
	// a slice of countries IDs
	IDs []int64
}

func (args *Countries) Values(q url.Values) url.Values {
	if args != nil {
		q.Add("countries_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(args.IDs)), ","), "[]"))
	}
	return q
}

type Providers struct {
	//
	IDs []int64
}

func (providers *Providers) Values(q url.Values) url.Values {
	if providers != nil {
		q.Add("providers_id", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(providers.IDs)), ","), "[]"))
	}
	return q
}
