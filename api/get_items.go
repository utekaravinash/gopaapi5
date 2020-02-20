package api

import (
	"errors"
)

type GetItemsResponse struct {
	Errors      []Error     `json:"Errors,omitempty"`
	ItemsResult ItemsResult `json:"ItemsResult,omitempty"`
}

type GetItemsParams struct {
	Condition             Condition
	CurrencyOfPreference  Currency
	ItemIds               []string
	LanguagesOfPreference []Language
	Merchant              Merchant
	OfferCount            int
	Resources             []Resource
}

func NewGetItemsParams(itemIds []string) *GetItemsParams {
	uniqueItemIds := []string{}

	for _, item := range itemIds {
		if !existsInStrings(item, uniqueItemIds) {
			uniqueItemIds = append(uniqueItemIds, item)
		}
	}

	return &GetItemsParams{
		ItemIds: uniqueItemIds,
	}
}

func (p *GetItemsParams) SetCondition(condition Condition) {
	p.Condition = condition
}

func (p *GetItemsParams) SetCurrencyOfPreference(currency Currency) {
	p.CurrencyOfPreference = currency
}

func (p *GetItemsParams) AddItemId(itemId string) {
	if !existsInStrings(itemId, p.ItemIds) {
		p.ItemIds = append(p.ItemIds, itemId)
	}
}

func (p *GetItemsParams) AddLanguagesOfPreference(language Language) {
	p.LanguagesOfPreference = append(p.LanguagesOfPreference, language)
}

func (p *GetItemsParams) SetMerchant(merchant Merchant) {
	p.Merchant = merchant
}

func (p *GetItemsParams) SetOfferCount(offerCount int) {
	p.OfferCount = offerCount
}

func (p *GetItemsParams) AddResources(resource Resource) {
	p.Resources = append(p.Resources, resource)
}

func (p GetItemsParams) GetResources() []Resource {
	return p.Resources
}

func (p GetItemsParams) Map() (map[string]interface{}, error) {
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
