package gopaapi5

import (
	"errors"

	"github.com/utekaravinash/gopaapi5/api"
)

func (c *Client) GetItems(params *api.GetItemsParams) (*api.GetItemsResponse, error) {

	if params == nil {
		return nil, errors.New("Nil parameters")
	}

	response := &api.GetItemsResponse{}

	payload, err := params.Map()
	if err != nil {
		return nil, err
	}

	err = c.sendRequest(payload, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
