package gopaapi5

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/utekaravinash/gopaapi5/api"
)

// Client stores AccessKey, SecretKey, and, AssociateTag; and exposes GetBrowseNodes, GetItems, GetVariations, and SearchItems operations.
type Client struct {
	AccessKey    string
	SecretKey    string
	AssociateTag string
	Locale       api.Locale
	partnerType  string
	service      string
	host         string
	region       string
	marketplace  string
	httpClient   *http.Client
	testing      bool
}

// NewClient accepts Access Key, Secrete Key, Associate Tag, Locale and returns a new client
func NewClient(accessKey, secretKey, associateTag string, locale api.Locale) (*Client, error) {

	if accessKey == "" {
		return nil, errors.New("Empty access key")
	}

	if secretKey == "" {
		return nil, errors.New("Empty secret key")
	}

	if associateTag == "" {
		return nil, errors.New("Empty associate tag")
	}

	if !locale.IsValid() {
		return nil, errors.New("Invalid locale")
	}

	client := &Client{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		AssociateTag: associateTag,
		partnerType:  "Associates",
		service:      "ProductAdvertisingAPIv1",
		host:         locale.Host(),
		region:       locale.Region(),
		marketplace:  locale.Marketplace(),
		httpClient:   &http.Client{},
		testing:      false,
	}

	return client, nil
}

// send sends a http request to Amazon Product Advertising service and returns response or error
func (c *Client) send(req *request, v interface{}) error {

	req.build()
	req.sign()

	resp, err := c.httpClient.Do(req.httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
