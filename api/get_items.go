package api

type GetItemsResponse struct{}

type GetItemsParams struct {
	Condition             Condition
	CurrencyOfPreference  Currency
	ItemIds               []string
	LanguagesOfPreference Language
	Merchant              string
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

func (p *GetItemsParams) SetLanguagesOfPreference(language Language) {
	p.LanguagesOfPreference = language
}

func (p *GetItemsParams) SetMerchant(merchant string) {
	p.Merchant = merchant
}

func (p *GetItemsParams) SetOfferCount(offerCount int) {
	p.OfferCount = offerCount
}

func (p *GetItemsParams) AddResources(resource Resource) {
	p.Resources = append(p.Resources, resource)
}

func (p GetItemsParams) Map() map[string]interface{} {
	return nil
}
