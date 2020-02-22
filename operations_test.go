package gopaapi5

import (
	"testing"

	"github.com/utekaravinash/gopaapi5/api"
)

func TestGetBrowseNodes(t *testing.T) {
	params := api.GetBrowseNodesParams{
		BrowseNodeIds: []string{"1234"},
	}

	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Fatalf("Error getting client: %s", err)
	}

	response, err := client.GetBrowseNodes(&params)
	if err != nil {
		t.Fatalf("Error in client.GetBrowseNodes: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"ResponseWithNoError", 3, len(response.ItemsResult.Items)},
		{"ResponseWithError", 2, len(response.Errors)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
