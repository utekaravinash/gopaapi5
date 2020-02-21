package api

import "testing"

func TestPropertiesAddRemoveExistsLength(t *testing.T) {
	properties := Properties{}

	t.Run("Properties Length", func(t *testing.T) {
		if properties.Length() > 0 {
			t.Errorf("Length of nil Properties variable cannot be greater than zero")
		}
	})

	t.Run("Properties Add", func(t *testing.T) {
		properties.Add("k1", "v1")
		properties.Add("k2", "v2")
		expected := 2
		actual := properties.Length()
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	})

	t.Run("Properties Remove", func(t *testing.T) {
		properties.Remove("k1")
		expected := 1
		actual := properties.Length()
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	})

	t.Run("Properties Exists", func(t *testing.T) {
		expected := true
		actual := properties.Exists("k2")
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	})
}
