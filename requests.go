package gopaapi5

import (
	"github.com/utekaravinash/gopaapi5/api"
)

func (c *Client) GetItems(params api.GetItemsParams) (*api.GetItemsResponse, error) {

	response := &api.GetItemsResponse{}
	payload := params.Map()

	err := c.SendRequest(payload, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
