// Code generated by goa v3.0.10, DO NOT EDIT.
//
// resume client
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package resume

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "resume" service client.
type Client struct {
	ListEndpoint goa.Endpoint
	AddEndpoint  goa.Endpoint
}

// NewClient initializes a "resume" service client given the endpoints.
func NewClient(list, add goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
		AddEndpoint:  add,
	}
}

// List calls the "list" endpoint of the "resume" service.
func (c *Client) List(ctx context.Context) (res StoredResumeCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredResumeCollection), nil
}

// Add calls the "add" endpoint of the "resume" service.
func (c *Client) Add(ctx context.Context, p []*Resume) (res []int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]int), nil
}
