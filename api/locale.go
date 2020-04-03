package api

// Locale custom type for locales
type Locale string

const (
	// Australia Locale for Australia
	Australia Locale = "AU"
	// Brazil Locale for Brazil
	Brazil Locale = "BR"
	// Canada Locale for Canada
	Canada Locale = "CA"
	// France Locale for France
	France Locale = "FR"
	// Germany Locale for Germany
	Germany Locale = "DE"
	// India Locale for India
	India Locale = "IN"
	// Italy Locale for Italy
	Italy Locale = "IT"
	// Japan Locale for Japan
	Japan Locale = "JP"
	// Mexico Locale for Mexico
	Mexico Locale = "MX"
	// Netherlands Locale for Netherlands
	Netherlands Locale = "NL"
	// Singapore Locale for Singapore
	Singapore Locale = "SG"
	// Spain Locale for Spain
	Spain Locale = "ES"
	// Turkey Locale for Turkey
	Turkey Locale = "TR"
	// UnitedArabEmirates Locale for United Arab Emirates
	UnitedArabEmirates Locale = "AE"
	// UnitedKingdom Locale for United Kingdom
	UnitedKingdom Locale = "UK"
	// UnitedStates Locale for United States
	UnitedStates Locale = "US"
)

// hostRegionMarketplace holds host, region and marketplace
type hostRegionMarketplace struct {
	host        string
	region      string
	marketplace string
}

// localeHostRegionMarketplaceMap maps a locale to its host, region and marketplace
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
	Netherlands:        {"webservices.amazon.nl", "eu-west-1", "www.amazon.nl"},
	Singapore:          {"webservices.amazon.sg", "us-west-2", "www.amazon.sg"},
	Spain:              {"webservices.amazon.es", "eu-west-1", "www.amazon.es"},
	Turkey:             {"webservices.amazon.com.tr", "eu-west-1", "www.amazon.com.tr"},
	UnitedArabEmirates: {"webservices.amazon.ae", "eu-west-1", "www.amazon.ae"},
	UnitedKingdom:      {"webservices.amazon.co.uk", "eu-west-1", "www.amazon.co.uk"},
	UnitedStates:       {"webservices.amazon.com", "us-east-1", "www.amazon.com"},
}

// Host returns host for a locale
func (l Locale) Host() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.host
	}

	return ""
}

// Region returns a region for a locale
func (l Locale) Region() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.region
	}

	return ""
}

// Marketplace returns a marketplace for a locale
func (l Locale) Marketplace() string {
	if hrm, ok := localeHostRegionMarketplaceMap[l]; ok {
		return hrm.marketplace
	}

	return ""
}

// IsValid checks if a locale is valie or not
func (l Locale) IsValid() bool {
	_, valid := localeHostRegionMarketplaceMap[l]
	return valid
}
