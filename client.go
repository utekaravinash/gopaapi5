package gopaapi5

import (
	"errors"

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
	}

	return client, nil
}

func (c *Client) sendRequest(payload map[string]interface{}, response interface{}) error {

	payload["PartnerType"] = c.partnerType
	payload["PartnerTag"] = c.AssociateTag
	payload["Marketplace"] = c.marketplace

	return nil
}
