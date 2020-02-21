package api

import "testing"

func TestExistsInStrings(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"StringExists", true, existsInStrings("ABC", []string{"ABC", "DEF"})},
		{"StringDoesNotExist", false, existsInStrings("XYZ", []string{"ABC", "DEF"})},
		{"CheckInEmptySlice", false, existsInStrings("XYZ", []string{})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
