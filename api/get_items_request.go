package api

type GetItemsRequest struct {
	condition             Condition
	currencyOfPreference  Currency
	itemIds               []string
	languagesOfPreference Language
	merchant              string
	offerCount            int
	resources             []Resource
}

func NewGetItemsRequest(itemIds []string) *GetItemsRequest {
	uniqueItemIds := []string{}

	for _, item := range itemIds {
		if !existsInStrings(item, uniqueItemIds) {
			uniqueItemIds = append(uniqueItemIds, item)
		}
	}

	return &GetItemsRequest{
		itemIds: uniqueItemIds,
	}
}

func (r *GetItemsRequest) SetCondition(condition Condition) {
	r.condition = condition
}

func (r *GetItemsRequest) SetCurrencyOfPreference(currency Currency) {
	r.currencyOfPreference = currency
}

func (r *GetItemsRequest) AddItemId(itemId string) {
	if !existsInStrings(itemId, r.itemIds) {
		r.itemIds = append(r.itemIds, itemId)
	}
}

func (r *GetItemsRequest) SetLanguagesOfPreference(language Language) {
	r.languagesOfPreference = language
}

func (r *GetItemsRequest) SetMerchant(merchant string) {
	r.merchant = merchant
}

func (r *GetItemsRequest) SetOfferCount(offerCount int) {
	r.offerCount = offerCount
}

func (r *GetItemsRequest) AddResources(resource Resource) {
	r.resources = append(r.resources, resource)
}
