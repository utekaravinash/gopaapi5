package api

import "testing"

func TestLocale(t *testing.T) {
	auLocale := Australia
	brLocale := Brazil
	caLocale := Canada
	frLocale := France
	deLocale := Germany
	inLocale := India
	itLocale := Italy
	jpLocale := Japan
	mxLocale := Mexico
	nlLocale := Netherlands
	sgLocale := Singapore
	esLocale := Spain
	trLocale := Turkey
	aeLocale := UnitedArabEmirates
	ukLocale := UnitedKingdom
	usLocale := UnitedStates
	fakeLocale := Locale("Fake Country")

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"AU Host", "webservices.amazon.com.au", auLocale.Host()},
		{"AU Region", "us-west-2", auLocale.Region()},
		{"AU Marketplace", "www.amazon.com.au", auLocale.Marketplace()},
		{"BR Host", "webservices.amazon.com.br", brLocale.Host()},
		{"BR Region", "us-east-1", brLocale.Region()},
		{"BR Marketplace", "www.amazon.com.br", brLocale.Marketplace()},
		{"CA Host", "webservices.amazon.ca", caLocale.Host()},
		{"CA Region", "us-east-1", caLocale.Region()},
		{"CA Marketplace", "www.amazon.ca", caLocale.Marketplace()},
		{"FR Host", "webservices.amazon.fr", frLocale.Host()},
		{"FR Region", "eu-west-1", frLocale.Region()},
		{"FR Marketplace", "www.amazon.fr", frLocale.Marketplace()},
		{"DE Host", "webservices.amazon.de", deLocale.Host()},
		{"DE Region", "eu-west-1", deLocale.Region()},
		{"DE Marketplace", "www.amazon.de", deLocale.Marketplace()},
		{"IN Host", "webservices.amazon.in", inLocale.Host()},
		{"IN Region", "eu-west-1", inLocale.Region()},
		{"IN Marketplace", "www.amazon.in", inLocale.Marketplace()},
		{"IT Host", "webservices.amazon.it", itLocale.Host()},
		{"IT Region", "eu-west-1", itLocale.Region()},
		{"IT Marketplace", "www.amazon.it", itLocale.Marketplace()},
		{"JP Host", "webservices.amazon.co.jp", jpLocale.Host()},
		{"JP Region", "us-west-2", jpLocale.Region()},
		{"JP Marketplace", "www.amazon.co.jp", jpLocale.Marketplace()},
		{"MX Host", "webservices.amazon.com.mx", mxLocale.Host()},
		{"MX Region", "us-east-1", mxLocale.Region()},
		{"MX Marketplace", "www.amazon.com.mx", mxLocale.Marketplace()},
		{"NL Host", "webservices.amazon.nl", nlLocale.Host()},
		{"NL Region", "eu-west-1", nlLocale.Region()},
		{"NL Marketplace", "www.amazon.nl", nlLocale.Marketplace()},
		{"SG Host", "webservices.amazon.sg", sgLocale.Host()},
		{"SG Region", "us-west-2", sgLocale.Region()},
		{"SG Marketplace", "www.amazon.sg", sgLocale.Marketplace()},
		{"ES Host", "webservices.amazon.es", esLocale.Host()},
		{"ES Region", "eu-west-1", esLocale.Region()},
		{"ES Marketplace", "www.amazon.es", esLocale.Marketplace()},
		{"TR Host", "webservices.amazon.com.tr", trLocale.Host()},
		{"TR Region", "eu-west-1", trLocale.Region()},
		{"TR Marketplace", "www.amazon.com.tr", trLocale.Marketplace()},
		{"AE Host", "webservices.amazon.ae", aeLocale.Host()},
		{"AE Region", "eu-west-1", aeLocale.Region()},
		{"AE Marketplace", "www.amazon.ae", aeLocale.Marketplace()},
		{"UK Host", "webservices.amazon.co.uk", ukLocale.Host()},
		{"UK Region", "eu-west-1", ukLocale.Region()},
		{"UK Marketplace", "www.amazon.co.uk", ukLocale.Marketplace()},
		{"US Host", "webservices.amazon.com", usLocale.Host()},
		{"US Region", "us-east-1", usLocale.Region()},
		{"US Marketplace", "www.amazon.com", usLocale.Marketplace()},
		{"Valid Locale", true, usLocale.IsValid()},
		{"Invalid Locale", false, fakeLocale.IsValid()},
		{"Invalid Locale Host", "", fakeLocale.Host()},
		{"Invalid Locale Region", "", fakeLocale.Region()},
		{"Invalid Locale Marketplace", "", fakeLocale.Marketplace()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expected != test.actual {
				t.Errorf("Expected: %v, Actual: %v", test.expected, test.actual)
			}
		})
	}
}
