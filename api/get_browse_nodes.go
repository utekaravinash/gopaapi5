package api

import (
	"errors"
)

// GetBrowseNodesResponse holds response from GetBrowseNodes operation
type GetBrowseNodesResponse struct {
	Errors            []Error           `json:"Errors,omitempty"`
	BrowseNodesResult BrowseNodesResult `json:"BrowseNodesResult,omitempty"`
}

// GetBrowseNodesParams holds parameters to be passed to GetBrowseNodes operation
type GetBrowseNodesParams struct {
	BrowseNodeIds         []string
	LanguagesOfPreference []Language
	Resources             []Resource
}

// ResourceList returns the list of resources in GetBrowseNodesParams
func (p GetBrowseNodesParams) ResourceList() []Resource {
	return p.Resources
}

// Payload constructs payload to be sent along with the API request
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
