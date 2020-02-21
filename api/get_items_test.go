package api

import (
	"testing"
)

func newGetItemsParams() *GetItemsParams {
	params := GetItemsParams{
		Condition:            New,
		CurrencyOfPreference: UnitedStatesDollar,
		Resources: []Resource{
			BrowseNodeInfoBrowseNodes,
			ItemInfoTitle,
		},
		ItemIds: []string{
			"B07HDBZN7Q",
			"B07QTVKNNQ",
			"1119293499",
		},
		LanguagesOfPreference: []Language{EnglishUnitedStates, SpanishUnitedStates},
		Merchant:              Amazon,
		OfferCount:            7,
	}

	return &params
}

func TestGetItemsParamsPayloadAndResourceList(t *testing.T) {
	params := newGetItemsParams()
	payload, err := params.Payload()
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count Resources", 2, len(payload["Resources"].([]Resource))},
		{"Count ItemIds", 3, len(payload["ItemIds"].([]string))},
		{"Condition", New, payload["Condition"].(Condition)},
		{"CurrencyOfPreference", UnitedStatesDollar, payload["CurrencyOfPreference"].(Currency)},
		{"Count LanguagesOfPreference", 2, len(payload["LanguagesOfPreference"].([]Language))},
		{"Merchant", Amazon, payload["Merchant"].(Merchant)},
		{"OfferCount", 7, payload["OfferCount"].(int)},
		{"Count ResourceList", 2, len(params.ResourceList())},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}

	t.Run("Zero ItemIds", func(t *testing.T) {
		nilParams := GetItemsParams{}
		_, err := nilParams.Payload()
		if err.Error() != "One or more item ids required" {
			t.Errorf(`Expected: "One or more item ids required", Actual: %v`, err)
		}
	})
}
