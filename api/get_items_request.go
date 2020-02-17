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
	return &GetItemsRequest{
		itemIds: itemIds,
	}
}

func (r *GetItemsRequest) SetCondition(condition Condition) {
	r.condition = condition
}

func (r *GetItemsRequest) SetCurrencyOfPreference(currency Currency) {
	r.currencyOfPreference = currency
}

func (r *GetItemsRequest) AddItemId(itemId string) {
	r.itemIds = append(r.itemIds, itemId)
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
