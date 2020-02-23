package api

import (
	"errors"
)

// GetItemsResponse holds response from GetItems operation
type GetItemsResponse struct {
	Errors      []Error     `json:"Errors,omitempty"`
	ItemsResult ItemsResult `json:"ItemsResult,omitempty"`
}

// GetItemsParams holds parameters to be passed to GetItems operation
type GetItemsParams struct {
	Condition             Condition
	CurrencyOfPreference  Currency
	ItemIds               []string
	LanguagesOfPreference []Language
	Merchant              Merchant
	OfferCount            int
	Resources             []Resource
}

// ResourceList returns the list of resources in GetItemsParams
func (p GetItemsParams) ResourceList() []Resource {
	return p.Resources
}

// Payload constructs payload to be sent along with the API request
func (p GetItemsParams) Payload() (map[string]interface{}, error) {
	kv := map[string]interface{}{}
	kv["ItemIdType"] = "ASIN"

	if p.Condition != "" {
		kv["Condition"] = p.Condition
	}

	if p.CurrencyOfPreference != "" {
		kv["CurrencyOfPreference"] = p.CurrencyOfPreference
	}

	if len(p.ItemIds) > 0 {
		kv["ItemIds"] = p.ItemIds
	} else {
		return nil, errors.New("One or more item ids required")
	}

	if len(p.LanguagesOfPreference) > 0 {
		kv["LanguagesOfPreference"] = p.LanguagesOfPreference
	}

	if p.Merchant != "" {
		kv["Merchant"] = p.Merchant
	}

	if p.OfferCount > 1 {
		kv["OfferCount"] = p.OfferCount
	}

	if len(p.Resources) > 0 {
		kv["Resources"] = p.Resources
	}

	return kv, nil
}
