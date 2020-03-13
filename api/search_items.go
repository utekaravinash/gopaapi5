package api

import (
	"encoding/json"
)

// SearchItemsResponse holds response from SearchItems operation
type SearchItemsResponse struct {
	Errors       []Error      `json:"Errors,omitempty"`
	SearchResult SearchResult `json:"SearchResult,omitempty"`
}

// SearchItemsParams holds parameters to be passed to SearchItems operation
type SearchItemsParams struct {
	Actor                 string         `json:"Actor,omitempty"`
	Artist                string         `json:"Artist,omitempty"`
	Author                string         `json:"Author,omitempty"`
	Availability          Availability   `json:"Availability,omitempty"`
	Brand                 string         `json:"Brand,omitempty"`
	BrowseNodeId          string         `json:"BrowseNodeId,omitempty"`
	Condition             Condition      `json:"Condition,omitempty"`
	CurrencyOfPreference  Currency       `json:"CurrencyOfPreference,omitempty"`
	DeliveryFlags         []DeliveryFlag `json:"DeliveryFlags,omitempty"`
	ItemCount             int            `json:"ItemCount,omitempty"`
	ItemPage              int            `json:"ItemPage,omitempty"`
	Keywords              string         `json:"Keywords,omitempty"`
	LanguagesOfPreference []Language     `json:"LanguagesOfPreference,omitempty"`
	MaxPrice              int            `json:"MaxPrice,omitempty"`
	Merchant              Merchant       `json:"Merchant,omitempty"`
	MinPrice              int            `json:"MinPrice,omitempty"`
	MinReviewsRating      int            `json:"MinReviewsRating,omitempty"`
	MinSavingPercent      int            `json:"MinSavingPercent,omitempty"`
	OfferCount            int            `json:"OfferCount,omitempty"`
	Properties            Properties     `json:"Properties,omitempty"`
	Resources             []Resource     `json:"Resources,omitempty"`
	SearchIndex           string         `json:"SearchIndex,omitempty"`
	SortBy                SortBy         `json:"SortBy,omitempty"`
	Title                 string         `json:"Title,omitempty"`
}

// ResourceList returns the list of resources in SearchItemsParams
func (p SearchItemsParams) ResourceList() []Resource {
	return p.Resources
}

// Payload constructs payload to be sent along with the API request
func (p SearchItemsParams) Payload() (map[string]interface{}, error) {
	kv := map[string]interface{}{}

	pj, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(pj, &kv)
	if err != nil {
		return nil, err
	}

	return kv, nil
}
