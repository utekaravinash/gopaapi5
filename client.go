package gopaapi5

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/utekaravinash/gopaapi5/api"
)

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
}

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
	}

	return client, nil
}

func (c *Client) execute(req *request, response interface{}) error {

	req.build()
	req.sign()

	resp, err := c.httpClient.Do(req.httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
