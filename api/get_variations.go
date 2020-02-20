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

func (p GetVariationsParams) ResourceList() []Resource {
	return p.Resources
}

func (p GetVariationsParams) Payload() (map[string]interface{}, error) {
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
