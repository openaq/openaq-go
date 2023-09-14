package openaq

import (
	"context"
	"fmt"
	"net/url"
)

type OwnersArgs struct {
	BaseArgs BaseArgs
}

func (ownersArgs *OwnersArgs) Values(q url.Values) (url.Values, error) {
	return q, nil
}

// QueryParams translates OwnerArgs struct into url.Values
func (args OwnersArgs) QueryParams() (url.Values, error) {
	q := make(url.Values)
	return q, nil
}

// GetOwners fetches all owners filtered by any params passed.
func (c *Client) GetOwners(ctx context.Context, args OwnersArgs) (*OwnersResponse, error) {
	resp := &OwnersResponse{}
	queryParams, err := args.QueryParams()
	if err != nil {
		return nil, err
	}
	err = c.request(ctx, "GET", "/owners", queryParams, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOwner fetches a single owner by ID.
func (c *Client) GetOwner(ctx context.Context, OwnersID int64) (*OwnersResponse, error) {
	path := fmt.Sprintf("/owners/%d", OwnersID)
	resp := &OwnersResponse{}
	err := c.request(ctx, "GET", path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
