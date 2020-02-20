package gopaapi5

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/utekaravinash/gopaapi5/api"
)

type payloadResourceListGetter interface {
	Payload() (map[string]interface{}, error)
	ResourceList() []api.Resource
}

func (c *Client) GetBrowseNodes(params *api.GetBrowseNodesParams) (*api.GetBrowseNodesResponse, error) {
	response := api.GetBrowseNodesResponse{}
	err := c.executeRequest(api.GetBrowseNodes, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetItems(params *api.GetItemsParams) (*api.GetItemsResponse, error) {
	response := api.GetItemsResponse{}

	err := c.executeRequest(api.GetItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetVariations(params *api.GetVariationsParams) (*api.GetVariationsResponse, error) {
	response := api.GetVariationsResponse{}

	err := c.executeRequest(api.GetVariations, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SearchItems(params *api.SearchItemsParams) (*api.SearchItemsResponse, error) {
	response := api.SearchItemsResponse{}

	err := c.executeRequest(api.SearchItems, params, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) executeRequest(operation api.Operation, params payloadResourceListGetter, v interface{}) error {

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

	req := &request{
		operation: operation,
		payload:   payload,
		client:    c,
		path:      "paapi5/getitems",
	}

	err = c.execute(req, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) execute(req *request, v interface{}) error {

	req.build()
	req.sign()

	resp, err := c.httpClient.Do(req.httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
