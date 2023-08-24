package openaq

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type MeasurementsArgs struct {
	BaseArgs     BaseArgs
	DatetimeFrom time.Time
	DatetimeTo   time.Time
	Parameters   *Parameters
	PeriodName   string
}

func (m *MeasurementsArgs) Values(q url.Values) (url.Values, error) {
	if !m.DatetimeFrom.IsZero() {
		q.Add("date_from", m.DatetimeFrom.UTC().Format("2006-01-02T15:04:05Z07:00"))
	}
	if !m.DatetimeTo.IsZero() {
		q.Add("date_to", m.DatetimeTo.UTC().Format("2006-01-02T15:04:05Z07:00"))
	}
	if m.PeriodName != "" {
		q.Add("period_name", m.PeriodName)
	}
	if m.Parameters != nil {
		q = m.Parameters.Values(q)
	}
	return q, nil
}

// QueryParams translates MeasurementsArgs struct into url.Values
func (args MeasurementsArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	q, err := args.BaseArgs.Values(q)
	if err != nil {
		return nil, err
	}
	q, err = args.Values(q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// GetLocationMeasurements fetches measurements for a single location by ID and any params passed
func (c *Client) GetLocationMeasurements(ctx context.Context, locationsID int64, args MeasurementsArgs) (*MeasurementsResponse, error) {
	path := fmt.Sprintf("/locations/%d/measurements", locationsID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &MeasurementsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSensorMeasurements fetches measurements for a single sensor by ID and any params passed
func (c *Client) GetSensorMeasurements(ctx context.Context, sensorsID int64, args MeasurementsArgs) (*MeasurementsResponse, error) {
	path := fmt.Sprintf("/sensors/%d/measurements", sensorsID)
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	resp := &MeasurementsResponse{}
	err = c.request(ctx, "GET", path, queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
