package openaq

import (
	"net/url"
	"strconv"
)

type BaseArgs struct {
	//
	Limit int64
	//
	Page int64
}

func (b *BaseArgs) Values(q url.Values) (url.Values, error) {
	if b.Limit != 0 {
		q.Add("limit", strconv.Itoa(int(b.Limit)))
	} else {
		q.Add("limit", "100")
	}
	if b.Page != 0 {
		q.Add("page", strconv.Itoa(int(b.Page)))
	} else {
		q.Add("page", "1")
	}
	return q, nil
}

type CoordinatesArgs struct {
	Lat float64
	Lon float64
}
