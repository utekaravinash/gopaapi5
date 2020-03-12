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
func (c *Client) GetBrowseNodes(ctx context.Context, params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {
	response := api.GetBrowseNodesResponse{}
	err := c.executeOperation(ctx, api.GetBrowseNodes, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetItems gets information for items
func (c *Client) GetItems(ctx context.Context, params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	response := api.GetItemsResponse{}

	err := c.executeOperation(ctx, api.GetItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetVariations gets information for variations
func (c *Client) GetVariations(ctx context.Context, params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	response := api.GetVariationsResponse{}

	err := c.executeOperation(ctx, api.GetVariations, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SearchItems searches for items on Amazon
func (c *Client) SearchItems(ctx context.Context, params *api.SearchItemsParams) (*api.SearchItemsResponse, error) {
	response := api.SearchItemsResponse{}

	err := c.executeOperation(ctx, api.SearchItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// payloadResourceLister is an interface for getting Payload and Resources
type payloadResourceLister interface {
	Payload() (map[string]interface{}, error)
	ResourceList() []api.Resource
}

// executeOperation validates request parameters and builds request payload
func (c *Client) executeOperation(ctx context.Context, operation api.Operation, params payloadResourceLister, v interface{}) error {

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
