package api

import "testing"

func newSearchItemsParams() *SearchItemsParams {
	properties := Properties{}
	properties.Add("k1", "v1")

	params := SearchItemsParams{
		Actor:                 "Tom Cruise",
		Artist:                "Tom",
		Author:                "JK",
		Availability:          IncludeOutOfStock,
		Brand:                 "string",
		BrowseNodeId:          "string",
		Condition:             Any,
		CurrencyOfPreference:  UnitedStatesDollar,
		DeliveryFlags:         []DeliveryFlag{AmazonGlobal, FreeShipping},
		ItemCount:             7,
		ItemPage:              4,
		Keywords:              "some keywords",
		LanguagesOfPreference: []Language{EnglishUnitedStates, SpanishUnitedStates},
		MaxPrice:              5,
		Merchant:              AllMerchants,
		MinPrice:              300,
		MinReviewsRating:      4,
		MinSavingPercent:      2,
		OfferCount:            8,
		Properties:            properties,
		Resources:             []Resource{BrowseNodeInfoBrowseNodes, BrowseNodeInfoBrowseNodesAncestor, BrowseNodesAncestor},
		SearchIndex:           "stringElectronics",
		SortBy:                Featured,
		Title:                 "Some Title",
	}

	return &params
}

func TestSearchItemsParamsPayloadAndResourceList(t *testing.T) {
	params := newSearchItemsParams()
	payload, err := params.Payload()
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Actor", "Tom Cruise", payload["Actor"].(string)},
		{"Artist", "Tom", payload["Artist"].(string)},
		{"Author", "JK", payload["Author"].(string)},
		{"Availability", IncludeOutOfStock, payload["Availability"].(Availability)},
		{"Brand", "string", payload["Brand"].(string)},
		{"BrowseNodeId", "string", payload["BrowseNodeId"].(string)},
		{"Condition", Any, payload["Condition"].(Condition)},
		{"CurrencyOfPreference", UnitedStatesDollar, payload["CurrencyOfPreference"].(Currency)},
		{"CountDeliveryFlags", 2, len(payload["DeliveryFlags"].([]DeliveryFlag))},
		{"ItemCount", 7, payload["ItemCount"].(int)},
		{"ItemPage", 4, payload["ItemPage"].(int)},
		{"Keywords", "some keywords", payload["Keywords"].(string)},
		{"CountLanguagesOfPreference", 2, len(payload["LanguagesOfPreference"].([]Language))},
		{"MaxPrice", 5, payload["MaxPrice"].(int)},
		{"Merchant", AllMerchants, payload["Merchant"].(Merchant)},
		{"MinPrice", 300, payload["MinPrice"].(int)},
		{"MinReviewsRating", 4, payload["MinReviewsRating"].(int)},
		{"MinSavingPercent", 2, payload["MinSavingPercent"].(int)},
		{"OfferCount", 8, payload["OfferCount"].(int)},
		{"Properties", 1, payload["Properties"].(Properties).Length()},
		{"CountResources", 3, len(payload["Resources"].([]Resource))},
		{"SearchIndex", "stringElectronics", payload["SearchIndex"].(string)},
		{"SortBy", Featured, payload["SortBy"].(SortBy)},
		{"Title", "Some Title", payload["Title"].(string)},
		{"Count ResourceList", 3, len(params.ResourceList())},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
