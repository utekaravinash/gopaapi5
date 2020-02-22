package gopaapi5

import (
	"fmt"
	"github.com/utekaravinash/gopaapi5/api"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
	"time"
)

var (
	testRequestTime         = time.Date(2020, time.February, 22, 11, 15, 9, 0, time.UTC)
	testAuthorizationHeader = "AWS4-HMAC-SHA256 Credential=accessKey/20200222/us-east-1/ProductAdvertisingAPI/aws4_request " +
		"SignedHeaders=content-encoding;content-type;host;x-amz-date;x-amz-target " +
		"Signature=a9c55a51fd4110f33032f20399d6d3df03df43f4a9b45188292132672dbe7aba"
)

func TestRequestBuildSign(t *testing.T) {
	client, err := NewClient("accessKey", "secretKey", "associateTag", api.UnitedStates)
	if err != nil {
		t.Fatalf("Error getting new client: %s", err)
	}

	req := &request{
		operation: api.GetItems,
		payload:   map[string]interface{}{},
		path:      "paapi5/getitems",
		client:    client,
		dateTime:  testRequestTime,
	}

	err = req.build()
	if err != nil {
		t.Fatalf("Error building request: %s", err)
	}
	err = req.sign()
	if err != nil {
		t.Fatalf("Error signing request: %s", err)
	}

	reqBody, err := ioutil.ReadAll(req.httpReq.Body)
	if err != nil {
		t.Errorf("Error reading request body: %s", err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"PayloadPartnerType", "Associates", req.payload["PartnerType"]},
		{"PayloadPartnerTag", "associateTag", req.payload["PartnerTag"]},
		{"PayloadMarketplace", "www.amazon.com", req.payload["Marketplace"]},
		{"HeaderHost", "webservices.amazon.com", req.httpReq.Header.Get("host")},
		{"HeaderContentType", "application/json; charset=utf-8", req.httpReq.Header.Get("content-type")},
		{"HeaderContentEncoding", "amz-1.0", req.httpReq.Header.Get("content-encoding")},
		{"HeaderXAmzDate", "20200222T111509Z", req.httpReq.Header.Get("x-amz-date")},
		{"HeaderXAmzTarget", "com.amazon.paapi5.v1.ProductAdvertisingAPIv1.GetItems", req.httpReq.Header.Get("x-amz-target")},
		{"AuthorizationHeader", testAuthorizationHeader, req.httpReq.Header.Get(authorizationHeader)},
		{"RequestBody", `{"Marketplace":"www.amazon.com","PartnerTag":"associateTag","PartnerType":"Associates"}`, string(reqBody)},
		{"RequestURL", "https://webservices.amazon.com/paapi5/getitems", req.httpReq.URL.String()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestStripExcessSpaces(t *testing.T) {

	valuesWithSpaceAtBothEnds := []string{"abc", "a b c", " abc", "abc ", " abc "}
	stripExcessSpaces(valuesWithSpaceAtBothEnds)

	valuesWithoutSpaceAtBothEnds := []string{"abc", "a b c", "abc", "abc", "abc"}
	stripExcessSpaces(valuesWithoutSpaceAtBothEnds)

	nilSlice := []string{}
	stripExcessSpaces(nilSlice)

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"ValuesWithSpaceAtBothEnds", "abc,a b c,abc,abc,abc", strings.Join(valuesWithSpaceAtBothEnds, ",")},
		{"ValuesWithoutSpaceAtBothEnds", "abc,a b c,abc,abc,abc", strings.Join(valuesWithoutSpaceAtBothEnds, ",")},
		{"NilSlice", "", strings.Join(nilSlice, ",")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestHmacSHA256(t *testing.T) {

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Key1,Data1", "8736733aaa19a38fe88073449795f85272d7a709a12d0ff226a1c65180a7a52d", fmt.Sprintf("%x", hmacSHA256([]byte("Key1"), []byte("Data1")))},
		{"Key2,Data2", "ef1be62d151e55c16944e80587c17bd29143c00812392fb6113c1b20cadc67bf", fmt.Sprintf("%x", hmacSHA256([]byte("Key2"), []byte("Data2")))},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestHashSHA256(t *testing.T) {

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"Data1", "e9fb5550dbc06ade237d302497f9b8b246984676f10562b9813d45da26ffa6d7", fmt.Sprintf("%x", hashSHA256([]byte("Data1")))},
		{"Data2", "24975629874c2650cab0e97f3a4e97a6e50f91278d5be470277a3bc2f317c459", fmt.Sprintf("%x", hashSHA256([]byte("Data2")))},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestFormatShortLongDate(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"ShortDate", "20200222", formatShortDate(testRequestTime)},
		{"LongDate", "20200222T111509Z", formatDate(testRequestTime)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}

func TestGetURIPath(t *testing.T) {
	url1, err := url.Parse("https://en.wikipedia.org/wiki/URL")
	if err != nil {
		t.Fatalf("Error parsing URL: %s", "https://en.wikipedia.org/wiki/URL")
	}

	url2, err := url.Parse("http://example.com/abc?xyz=def")
	if err != nil {
		t.Fatalf("Error parsing URL: %s", "https://en.wikipedia.org/wiki/URL")
	}

	url3, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("Error parsing URL: %s", "http://example.com")
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"URL1", "/wiki/URL", getURIPath(url1)},
		{"URL2", "/abc", getURIPath(url2)},
		{"URL3", "/", getURIPath(url3)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
