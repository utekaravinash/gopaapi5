package gopaapi5

import (
	"errors"

	"github.com/utekaravinash/gopaapi5/api"
)

type Operation string

const (
	GetBrowseNodes Operation = "GetBrowseNodes"
	GetItems       Operation = "GetItems"
	GetVariations  Operation = "GetVariations"
	SearchItems    Operation = "SearchItems"
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

	req := &request{
		Operation: GetItems,
		Payload:   payload,
		client:    c,
		path:      "paapi5/getitems?say=hello&do=otherwise",
	}

	err = c.execute(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
