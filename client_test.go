package gopaapi5

import (
	"testing"

	"github.com/utekaravinash/gopaapi5/api"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Errorf("Error getting api client instance")
	}

	_, err1 := NewClient("", "", "", api.UnitedStates)
	_, err2 := NewClient("accessKey", "", "", api.UnitedStates)
	_, err3 := NewClient("accessKey", "secretKey", "", api.UnitedStates)
	_, err4 := NewClient("accessKey", "secretKey", "associateTag", api.Locale("Fake Locale"))

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"AccessKey", "accessKey", client.AccessKey},
		{"SecretKey", "secretKey", client.SecretKey},
		{"AssociateTag", "associateTag", client.AssociateTag},
		{"Empty AccessKey", "Empty access key", err1.Error()},
		{"Empty SecretKey", "Empty secret key", err2.Error()},
		{"Empty AssociateTag", "Empty associate tag", err3.Error()},
		{"Invalid Locale", "Invalid locale", err4.Error()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
