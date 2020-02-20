package api

type GetVariationsResponse struct {
	Errors      []Error     `json:"Errors,omitempty"`
	ItemsResult ItemsResult `json:"ItemsResult,omitempty"`
}

type GetVariationsParams struct {
	ASIN                  string
	Condition             Condition
	CurrencyOfPreference  Currency
	LanguagesOfPreference []Language
	Merchant              Merchant
	OfferCount            int
	Resources             []Resource
	VariationCount        int
	VariationPage         int
}

func NewGetVariationsParams(asin string) *GetVariationsParams {

	return &GetVariationsParams{
		ASIN: asin,
	}
}

func (p *GetVariationsParams) SetASIN(asin string) {
	p.ASIN = asin
}

func (p *GetVariationsParams) SetCondition(condition Condition) {
	p.Condition = condition
}

func (p *GetVariationsParams) SetCurrencyOfPreference(currency Currency) {
	p.CurrencyOfPreference = currency
}

func (p *GetVariationsParams) AddLanguagesOfPreference(language Language) {
	p.LanguagesOfPreference = append(p.LanguagesOfPreference, language)
}

func (p *GetVariationsParams) SetMerchant(merchant Merchant) {
	p.Merchant = merchant
}

func (p *GetVariationsParams) SetOfferCount(offerCount int) {
	p.OfferCount = offerCount
}

func (p *GetVariationsParams) AddResources(resource Resource) {
	p.Resources = append(p.Resources, resource)
}

func (p *GetVariationsParams) SetVariationCount(variationCount int) {
	p.VariationCount = variationCount
}

func (p *GetVariationsParams) SetVariationPage(variationPage int) {
	p.VariationPage = variationPage
}

func (p GetVariationsParams) GetResources() []Resource {
	return p.Resources
}

func (p GetVariationsParams) Map() (map[string]interface{}, error) {
	kv := map[string]interface{}{}
	kv["ASIN"] = p.ASIN

	if p.Condition != "" {
		kv["Condition"] = p.Condition
	}

	if p.CurrencyOfPreference != "" {
		kv["CurrencyOfPreference"] = p.CurrencyOfPreference
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

	if p.VariationCount > 1 {
		kv["VariationCount"] = p.VariationCount
	}

	if p.VariationPage > 1 {
		kv["VariationPage"] = p.VariationPage
	}

	return kv, nil
}
