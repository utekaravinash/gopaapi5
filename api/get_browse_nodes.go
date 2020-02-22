package api

import (
	"errors"
)

type GetBrowseNodesResponse struct {
	Errors            []Error           `json:"Errors,omitempty"`
	BrowseNodesResult BrowseNodesResult `json:"BrowseNodesResult,omitempty"`
}

type GetBrowseNodesParams struct {
	BrowseNodeIds         []string
	LanguagesOfPreference []Language
	Resources             []Resource
}

func (p GetBrowseNodesParams) ResourceList() []Resource {
	return p.Resources
}

func (p GetBrowseNodesParams) Payload() (map[string]interface{}, error) {
	kv := map[string]interface{}{}

	if len(p.BrowseNodeIds) > 0 {
		kv["BrowseNodeIds"] = p.BrowseNodeIds
	} else {
		return nil, errors.New("One or more browse node ids required")
	}

	if len(p.LanguagesOfPreference) > 0 {
		kv["LanguagesOfPreference"] = p.LanguagesOfPreference
	}

	if len(p.Resources) > 0 {
		kv["Resources"] = p.Resources
	}

	return kv, nil
}
