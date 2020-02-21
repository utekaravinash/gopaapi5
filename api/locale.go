package api

type Locale string

const (
	Australia          Locale = "AU"
	Brazil             Locale = "BR"
	Canada             Locale = "CA"
	France             Locale = "FR"
	Germany            Locale = "DE"
	India              Locale = "IN"
	Italy              Locale = "IT"
	Japan              Locale = "JP"
	Mexico             Locale = "MX"
	Singapore          Locale = "SG"
	Spain              Locale = "ES"
	Turkey             Locale = "TR"
	UnitedArabEmirates Locale = "AE"
	UnitedKingdom      Locale = "UK"
	UnitedStates       Locale = "US"
)

type hostRegionMarketplace struct {
	host        string
	region      string
	marketplace string
}

var localeHostRegionMarketplaceMap = map[Locale]hostRegionMarketplace{
	Australia:          {"webservices.amazon.com.au", "us-west-2", "www.amazon.com.au"},
	Brazil:             {"webservices.amazon.com.br", "us-east-1", "www.amazon.com.br"},
	Canada:             {"webservices.amazon.ca", "us-east-1", "www.amazon.ca"},
	France:             {"webservices.amazon.fr", "eu-west-1", "www.amazon.fr"},
	Germany:            {"webservices.amazon.de", "eu-west-1", "www.amazon.de"},
	India:              {"webservices.amazon.in", "eu-west-1", "www.amazon.in"},
	Italy:              {"webservices.amazon.it", "eu-west-1", "www.amazon.it"},
	Japan:              {"webservices.amazon.co.jp", "us-west-2", "www.amazon.co.jp"},
	Mexico:             {"webservices.amazon.com.mx", "us-east-1", "www.amazon.com.mx"},
	Singapore:          {"webservices.amazon.sg", "us-west-2", "www.amazon.sg"},
	Spain:              {"webservices.amazon.es", "eu-west-1", "www.amazon.es"},
	Turkey:             {"webservices.amazon.com.tr", "eu-west-1", "www.amazon.com.tr"},
	UnitedArabEmirates: {"webservices.amazon.ae", "eu-west-1", "www.amazon.ae"},
	UnitedKingdom:      {"webservices.amazon.co.uk", "eu-west-1", "www.amazon.co.uk"},
	UnitedStates:       {"webservices.amazon.com", "us-east-1", "www.amazon.com"},
}

func (l Locale) Host() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.host
	}

	return ""
}

func (l Locale) Region() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.region
	}

	return ""
}

func (l Locale) Marketplace() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.marketplace
	}

	return ""
}

func (l Locale) IsValid() bool {
	_, valid := localeHostRegionMarketplaceMap[l]
	return valid
}
