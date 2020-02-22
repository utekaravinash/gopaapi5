package gopaapi5

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
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
	client.testing = true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("_response/get_browse_nodes_success.json")
		if err != nil {
			t.Errorf("Cannot open _response/get_browse_nodes_success.json")
		}

		jo, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Cannot read _response/get_browse_nodes_success.json")
		}

		w.Write(jo)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client.httpClient = httpClient

	response, err := client.GetBrowseNodes(&params)
	if err != nil {
		t.Fatalf("Error in client.GetBrowseNodes: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count BrowseNodes", 2, len(response.BrowseNodesResult.BrowseNodes)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	params := api.GetItemsParams{
		ItemIds: []string{"0892131349"},
	}

	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Fatalf("Error getting client: %s", err)
	}
	client.testing = true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("_response/get_items_success.json")
		if err != nil {
			t.Errorf("Cannot open _response/get_items_success.json")
		}

		jo, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Cannot read _response/get_items_success.json")
		}

		w.Write(jo)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client.httpClient = httpClient

	response, err := client.GetItems(&params)
	if err != nil {
		t.Fatalf("Error in client.GetItems: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count Items", 2, len(response.ItemsResult.Items)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestGetVariations(t *testing.T) {
	params := api.GetVariationsParams{}

	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Fatalf("Error getting client: %s", err)
	}
	client.testing = true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("_response/get_variations_success.json")
		if err != nil {
			t.Errorf("Cannot open _response/get_variations_success.json")
		}

		jo, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Cannot read _response/get_variations_success.json")
		}

		w.Write(jo)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client.httpClient = httpClient

	response, err := client.GetVariations(&params)
	if err != nil {
		t.Fatalf("Error in client.GetVariations: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count Items", 9, len(response.VariationsResult.Items)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestSearchItems(t *testing.T) {
	params := api.SearchItemsParams{}

	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Fatalf("Error getting client: %s", err)
	}
	client.testing = true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("_response/search_items_success.json")
		if err != nil {
			t.Errorf("Cannot open _response/search_items_success.json")
		}

		jo, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("Cannot read _response/search_items_success.json")
		}

		w.Write(jo)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	client.httpClient = httpClient

	response, err := client.SearchItems(&params)
	if err != nil {
		t.Fatalf("Error in client.SearchItems: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Count Items", 3, len(response.SearchResult.Items)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return cli, s.Close
}
