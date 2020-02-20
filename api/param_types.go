package api

type Condition string

const (
	Any         Condition = "Any"
	New         Condition = "New"
	Used        Condition = "Used"
	Collectible Condition = "Collectible"
	Refurbished Condition = "Refurbished"
)

type Availability string

const (
	Available         Availability = "Available"
	IncludeOutOfStock Availability = "IncludeOutOfStock"
)

type DeliveryFlag string

const (
	AmazonGlobal      DeliveryFlag = "AmazonGlobal"
	FreeShipping      DeliveryFlag = "FreeShipping"
	FulfilledByAmazon DeliveryFlag = "FulfilledByAmazon"
	Prime             DeliveryFlag = "Prime"
)

type Merchant string

const (
	AllMerchants Merchant = "All"
	Amazon       Merchant = "Amazon"
)

type SortBy string

const (
	AvgCustomerReviews SortBy = "AvgCustomerReviews"
	Featured           SortBy = "Featured"
	NewestArrivals     SortBy = "NewestArrivals"
	PriceHighToLow     SortBy = "Price:HighToLow"
	PriceLowToHigh     SortBy = "Price:LowToHigh"
	Relevance          SortBy = "Relevance"
)

type Properties map[string]string

func (p Properties) Add(key, value string) {
	p[key] = value
}

func (p Properties) Remove(key string) {
	delete(p, key)
}

func (p Properties) Exists(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Properties) Length() int {
	return len(p)
}
