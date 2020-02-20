package gopaapi5

import (
	"errors"

	"github.com/utekaravinash/gopaapi5/api"
)

func (c *Client) GetBrowseNodes(params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {

	if params == nil {
		return nil, errors.New("Nil parameters")
	}

	operation := api.GetBrowseNodes
	err := operation.Validate(params.Resources)
	if err != nil {
		return nil, err
	}

	response := api.GetBrowseNodesResponse{}

	payload, err := params.Map()
	if err != nil {
		return nil, err
	}

	req := &request{
		Operation: operation,
		Payload:   payload,
		client:    c,
		path:      "paapi5/getitems",
	}

	err = c.execute(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetItems(params *api.GetItemsParams) (*api.GetItemsResponse, error) {

	if params == nil {
		return nil, errors.New("Nil parameters")
	}

	operation := api.GetItems
	err := operation.Validate(params.Resources)
	if err != nil {
		return nil, err
	}

	response := api.GetItemsResponse{}

	payload, err := params.Map()
	if err != nil {
		return nil, err
	}

	req := &request{
		Operation: operation,
		Payload:   payload,
		client:    c,
		path:      "paapi5/getitems",
	}

	err = c.execute(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetVariations(params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {

	if params == nil {
		return nil, errors.New("Nil parameters")
	}

	operation := api.GetVariations
	err := operation.Validate(params.Resources)
	if err != nil {
		return nil, err
	}

	response := api.GetVariationsResponse{}

	payload, err := params.Map()
	if err != nil {
		return nil, err
	}

	req := &request{
		Operation: operation,
		Payload:   payload,
		client:    c,
		path:      "paapi5/getitems",
	}

	err = c.execute(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
