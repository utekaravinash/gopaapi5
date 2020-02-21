package api

import "testing"

func newGetBrowseNodesParams() *GetBrowseNodesParams {
	params := GetBrowseNodesParams{
		BrowseNodeIds:         []string{"123", "4567"},
		LanguagesOfPreference: []Language{EnglishUnitedStates, SpanishUnitedStates},
		Resources:             []Resource{BrowseNodeInfoBrowseNodes, BrowseNodeInfoBrowseNodesAncestor, ImagesPrimaryLarge},
	}

	return &params
}

func TestGetBrowseNodesParamsAndResourceList(t *testing.T) {
	params := newGetBrowseNodesParams()
	payload, err := params.Payload()
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count BrowseNodeIds", 2, len(payload["BrowseNodeIds"].([]string))},
		{"Count LanguagesOfPreference", 2, len(payload["LanguagesOfPreference"].([]Language))},
		{"Count Resources", 3, len(payload["Resources"].([]Resource))},
		{"Count ResourceList", 3, len(params.ResourceList())},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}

	t.Run("Zero BrowseNodeIds", func(t *testing.T) {
		nilParams := GetBrowseNodesParams{}
		_, err := nilParams.Payload()
		if err.Error() != "One or more browse node ids required" {
			t.Errorf(`Expected: "One or more browse node ids required", Actual: %v`, err)
		}
	})
}
