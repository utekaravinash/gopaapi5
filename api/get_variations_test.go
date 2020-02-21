package api

import "testing"

func newGetVariationsParams() *GetVariationsParams {
	params := GetVariationsParams{
		ASIN:                  "0545162076",
		Condition:             New,
		CurrencyOfPreference:  UnitedStatesDollar,
		LanguagesOfPreference: []Language{EnglishUnitedStates, SpanishUnitedStates},
		Merchant:              AllMerchants,
		OfferCount:            7,
		Resources:             []Resource{BrowseNodeInfoBrowseNodes, CustomerReviewsCount},
		VariationCount:        3,
		VariationPage:         5,
	}

	return &params
}

func TestGetVariationsParamsPayloadAndResourceList(t *testing.T) {
	params := newGetVariationsParams()
	payload, err := params.Payload()
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"ASIN", "0545162076", payload["ASIN"].(string)},
		{"Condition", New, payload["Condition"].(Condition)},
		{"CurrencyOfPreference", UnitedStatesDollar, payload["CurrencyOfPreference"].(Currency)},
		{"Count LanguagesOfPreference", 2, len(payload["LanguagesOfPreference"].([]Language))},
		{"Merchant", AllMerchants, payload["Merchant"].(Merchant)},
		{"OfferCount", 7, payload["OfferCount"].(int)},
		{"Count Resources", 2, len(payload["Resources"].([]Resource))},
		{"VariationCount", 3, payload["VariationCount"].(int)},
		{"VariationPage", 5, payload["VariationPage"].(int)},
		{"Count ResourceList", 2, len(params.ResourceList())},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
