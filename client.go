package gopaapi5

import (
	"errors"
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
	testing      bool
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
		testing:      false,
	}

	return client, nil
}
