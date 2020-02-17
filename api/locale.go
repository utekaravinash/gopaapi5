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

type hostRegion struct {
	host   string
	region string
}

var localeHostRegionMap = map[Locale]hostRegion{
	Australia:          {"webservices.amazon.com.au", "us-west-2"},
	Brazil:             {"webservices.amazon.com.br", "us-east-1"},
	Canada:             {"webservices.amazon.ca", "us-east-1"},
	France:             {"webservices.amazon.fr", "eu-west-1"},
	Germany:            {"webservices.amazon.de", "eu-west-1"},
	India:              {"webservices.amazon.in", "eu-west-1"},
	Italy:              {"webservices.amazon.it", "eu-west-1"},
	Japan:              {"webservices.amazon.co.jp", "us-west-2"},
	Mexico:             {"webservices.amazon.com.mx", "us-east-1"},
	Singapore:          {"webservices.amazon.sg", "us-west-2"},
	Spain:              {"webservices.amazon.es", "eu-west-1"},
	Turkey:             {"webservices.amazon.com.tr", "eu-west-1"},
	UnitedArabEmirates: {"webservices.amazon.ae", "eu-west-1"},
	UnitedKingdom:      {"webservices.amazon.co.uk", "eu-west-1"},
	UnitedStates:       {"webservices.amazon.com", "us-east-1"},
}

func (l Locale) Host() string {
	if hr, ok := localeHostRegionMap[l]; ok {
		return hr.host
	}

	return ""
}

func (l Locale) Region() string {
	if hr, ok := localeHostRegionMap[l]; ok {
		return hr.region
	}

	return ""
}

func (l Locale) IsValid() bool {
	_, valid := localeHostRegionMap[l]
	return valid
}
