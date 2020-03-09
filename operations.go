package gopaapi5

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/utekaravinash/gopaapi5/api"
)

// GetBrowseNodes gets information for Browse Nodes
func (c *Client) GetBrowseNodes(params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {
	return c.GetBrowseNodesCtx(context.Background(), params)
}

// GetItems gets information for items
func (c *Client) GetItems(params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	return c.GetItemsCtx(context.Background(), params)
}

// GetVariations gets information for variations
func (c *Client) GetVariations(params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	return c.GetVariationsCtx(context.Background(), params)
}

// SearchItems searches for items on Amazon
func (c *Client) SearchItems(params *api.SearchItemsParams) (*api.SearchItemsResponse, error) {
	return c.SearchItemsCtx(context.Background(), params)
}

// GetBrowseNodesCtx gets information for Browse Nodes, also accepts context.Context
func (c *Client) GetBrowseNodesCtx(ctx context.Context, params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {
	response := api.GetBrowseNodesResponse{}
	err := c.executeOperation(ctx, api.GetBrowseNodes, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetItemsCtx gets information for items, also accepts context.Context
func (c *Client) GetItemsCtx(ctx context.Context, params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	response := api.GetItemsResponse{}

	err := c.executeOperation(ctx, api.GetItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetVariationsCtx gets information for variations, also accepts context.Context
func (c *Client) GetVariationsCtx(ctx context.Context, params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	response := api.GetVariationsResponse{}

	err := c.executeOperation(ctx, api.GetVariations, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SearchItemsCtx searches for items on Amazon, also accepts context.Context
func (c *Client) SearchItemsCtx(ctx context.Context, params *api.SearchItemsParams) (*api.SearchItemsResponse, error) {
	response := api.SearchItemsResponse{}

	err := c.executeOperation(ctx, api.SearchItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// payloadResourceListGetter is an interface for getting Payload and Resources
type payloadResourceListGetter interface {
	Payload() (map[string]interface{}, error)
	ResourceList() []api.Resource
}

// executeOperation validates request parameters and builds request payload
func (c *Client) executeOperation(ctx context.Context, operation api.Operation, params payloadResourceListGetter, v interface{}) error {

	if params == nil {
		return errors.New("Nil parameters")
	}

	err := operation.Validate(params.ResourceList())
	if err != nil {
		return err
	}

	payload, err := params.Payload()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("paapi5/%s", strings.ToLower(string(operation)))

	req := &request{
		operation: operation,
		payload:   payload,
		client:    c,
		path:      path,
		dateTime:  time.Now().UTC(),
		ctx:       ctx,
	}

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}
