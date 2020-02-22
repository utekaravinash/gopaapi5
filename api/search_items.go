package api

type SearchItemsResponse struct {
	Errors       []Error      `json:"Errors,omitempty"`
	SearchResult SearchResult `json:"SearchResult,omitempty"`
}

type SearchItemsParams struct {
	Actor                 string
	Artist                string
	Author                string
	Availability          Availability
	Brand                 string
	BrowseNodeId          string
	Condition             Condition
	CurrencyOfPreference  Currency
	DeliveryFlags         []DeliveryFlag
	ItemCount             int
	ItemPage              int
	Keywords              string
	LanguagesOfPreference []Language
	MaxPrice              int
	Merchant              Merchant
	MinPrice              int
	MinReviewsRating      int
	MinSavingPercent      int
	OfferCount            int
	Properties            Properties
	Resources             []Resource
	SearchIndex           string
	SortBy                SortBy
	Title                 string
}

func (p SearchItemsParams) ResourceList() []Resource {
	return p.Resources
}

func (p SearchItemsParams) Payload() (map[string]interface{}, error) {
	kv := map[string]interface{}{}

	if p.Actor != "" {
		kv["Actor"] = p.Actor
	}

	if p.Artist != "" {
		kv["Artist"] = p.Artist
	}

	if p.Author != "" {
		kv["Author"] = p.Author
	}

	if p.Availability != "" {
		kv["Availability"] = p.Availability
	}

	if p.Brand != "" {
		kv["Brand"] = p.Brand
	}

	if p.BrowseNodeId != "" {
		kv["BrowseNodeId"] = p.BrowseNodeId
	}

	if p.Condition != "" {
		kv["Condition"] = p.Condition
	}

	if p.CurrencyOfPreference != "" {
		kv["CurrencyOfPreference"] = p.CurrencyOfPreference
	}

	if p.Keywords != "" {
		kv["Keywords"] = p.Keywords
	}

	if p.Merchant != "" {
		kv["Merchant"] = p.Merchant
	}

	if p.Properties.Length() > 0 {
		kv["Properties"] = p.Properties
	}

	if p.SearchIndex != "" {
		kv["SearchIndex"] = p.SearchIndex
	}

	if p.SortBy != "" {
		kv["SortBy"] = p.SortBy
	}

	if p.Title != "" {
		kv["Title"] = p.Title
	}

	if p.ItemCount > 0 {
		kv["ItemCount"] = p.ItemCount
	}

	if p.ItemPage > 0 {
		kv["ItemPage"] = p.ItemPage
	}

	if p.MaxPrice > 0 {
		kv["MaxPrice"] = p.MaxPrice
	}

	if p.MinPrice > 0 {
		kv["MinPrice"] = p.MinPrice
	}

	if p.MinReviewsRating > 0 {
		kv["MinReviewsRating"] = p.MinReviewsRating
	}

	if p.MinSavingPercent > 0 {
		kv["MinSavingPercent"] = p.MinSavingPercent
	}

	if p.OfferCount > 0 {
		kv["OfferCount"] = p.OfferCount
	}

	if len(p.DeliveryFlags) > 0 {
		kv["DeliveryFlags"] = p.DeliveryFlags
	}

	if len(p.LanguagesOfPreference) > 0 {
		kv["LanguagesOfPreference"] = p.LanguagesOfPreference
	}

	if len(p.Resources) > 0 {
		kv["Resources"] = p.Resources
	}

	return kv, nil
}
