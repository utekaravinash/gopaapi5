package api

type SearchItemsResponse struct {
	Errors      []Error     `json:"Errors,omitempty"`
	ItemsResult ItemsResult `json:"ItemsResult,omitempty"`
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

func NewSearchItemsParams() *SearchItemsParams {
	return &SearchItemsParams{}
}

func (p *SearchItemsParams) SetActor(actor string) {
	p.Actor = actor
}

func (p *SearchItemsParams) SetArtist(artist string) {
	p.Artist = artist
}

func (p *SearchItemsParams) SetAuthor(author string) {
	p.Author = author
}

func (p *SearchItemsParams) SetAvailability(availability Availability) {
	p.Availability = availability
}

func (p *SearchItemsParams) SetBrand(brand string) {
	p.Brand = brand
}

func (p *SearchItemsParams) SetBrowseNodeId(browseNodeId string) {
	p.BrowseNodeId = browseNodeId
}

func (p *SearchItemsParams) SetCondition(condition Condition) {
	p.Condition = condition
}

func (p *SearchItemsParams) SetCurrencyOfPreference(currencyOfPreference Currency) {
	p.CurrencyOfPreference = currencyOfPreference
}

func (p *SearchItemsParams) AddDeliveryFlag(deliveryFlag DeliveryFlag) {
	p.DeliveryFlags = append(p.DeliveryFlags, deliveryFlag)
}

func (p *SearchItemsParams) AddLanguagesOfPreference(language Language) {
	p.LanguagesOfPreference = append(p.LanguagesOfPreference, language)
}

func (p *SearchItemsParams) AddResources(resource Resource) {
	p.Resources = append(p.Resources, resource)
}

func (p *SearchItemsParams) SetItemCount(itemCount int) {
	p.ItemCount = itemCount
}

func (p *SearchItemsParams) SetItemPage(itemPage int) {
	p.ItemPage = itemPage
}

func (p *SearchItemsParams) SetKeywords(keywords string) {
	p.Keywords = keywords
}

func (p *SearchItemsParams) SetMaxPrice(maxPrice int) {
	p.MaxPrice = maxPrice
}

func (p *SearchItemsParams) SetMerchant(merchant Merchant) {
	p.Merchant = merchant
}

func (p *SearchItemsParams) SetMinPrice(minPrice int) {
	p.MinPrice = minPrice
}

func (p *SearchItemsParams) SetMinReviewsRating(minReviewsRating int) {
	p.MinReviewsRating = minReviewsRating
}

func (p *SearchItemsParams) SetMinSavingPercent(minSavingPercent int) {
	p.MinSavingPercent = minSavingPercent
}

func (p *SearchItemsParams) SetOfferCount(offerCount int) {
	p.OfferCount = offerCount
}

func (p *SearchItemsParams) SetProperties(properties Properties) {
	p.Properties = properties
}

func (p *SearchItemsParams) SetSearchIndex(searchIndex string) {
	p.SearchIndex = searchIndex
}

func (p *SearchItemsParams) SetSortBy(sortBy SortBy) {
	p.SortBy = sortBy
}

func (p *SearchItemsParams) SetTitle(title string) {
	p.Title = title
}

func (p SearchItemsParams) Map() (map[string]interface{}, error) {
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
