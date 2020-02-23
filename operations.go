package gopaapi5

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/utekaravinash/gopaapi5/api"
)

// GetBrowseNodes gets information for Browse Nodes
func (c *Client) GetBrowseNodes(params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {
	response := api.GetBrowseNodesResponse{}
	err := c.executeOperation(api.GetBrowseNodes, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBrowseNodes gets item information for items
func (c *Client) GetItems(params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	response := api.GetItemsResponse{}

	err := c.executeOperation(api.GetItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetVariations gets information for variations
func (c *Client) GetVariations(params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	response := api.GetVariationsResponse{}

	err := c.executeOperation(api.GetVariations, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SearchItems searches for items on Amazon
func (c *Client) SearchItems(params *api.SearchItemsParams) (*api.SearchItemsResponse, error) {
	response := api.SearchItemsResponse{}

	err := c.executeOperation(api.SearchItems, params, &response)
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
func (c *Client) executeOperation(operation api.Operation, params payloadResourceListGetter, v interface{}) error {

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
	}

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}
