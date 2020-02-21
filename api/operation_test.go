package api

import (
	"fmt"
	"testing"
)

func TestOperationValidate(t *testing.T) {

	getBrowseNodes := GetBrowseNodes
	getItems := GetItems

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"ValidResources", nil, getBrowseNodes.Validate([]Resource{BrowseNodesAncestor, BrowseNodesChildren})},
		{"InvalidResources", fmt.Sprintf(`Invalid resource "%s" for operation "%s"`, BrowseNodesAncestor, getItems), getItems.Validate([]Resource{BrowseNodesAncestor, BrowseNodesChildren}).Error()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
