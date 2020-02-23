package api

// Condition custom type for product conditions
type Condition string

const (
	// Any constant for any product condition
	Any Condition = "Any"
	// New constant for new product condition
	New Condition = "New"
	// Used constant for used product condition
	Used Condition = "Used"
	// Collectible constant for collectible product condition
	Collectible Condition = "Collectible"
	// Refurbished constant for refurbished product condition
	Refurbished Condition = "Refurbished"
)

// Availability custom type for product availabilities
type Availability string

const (
	// Available constant for available products
	Available Availability = "Available"
	// IncludeOutOfStock constant to include out of stock products
	IncludeOutOfStock Availability = "IncludeOutOfStock"
)

// DeliveryFlag custom type for delivery flags
type DeliveryFlag string

const (
	// AmazonGlobal constant for AmazonGlobal delivery flag
	AmazonGlobal DeliveryFlag = "AmazonGlobal"
	// FreeShipping constant for FreeShipping delivery flag
	FreeShipping DeliveryFlag = "FreeShipping"
	// FulfilledByAmazon constant for FulfilledByAmazon delivery flag
	FulfilledByAmazon DeliveryFlag = "FulfilledByAmazon"
	// Prime constant for Prime delivery flag
	Prime DeliveryFlag = "Prime"
)

// Merchant custom type for merchant types
type Merchant string

const (
	// AllMerchants constant for any merchant
	AllMerchants Merchant = "All"
	// Amazon constant for Amazon merchant
	Amazon Merchant = "Amazon"
)

// SortBy custom type for sort by operation
type SortBy string

const (
	// AvgCustomerReviews constant to sort by AvgCustomerReviews
	AvgCustomerReviews SortBy = "AvgCustomerReviews"
	// Featured constant to sort by Featured
	Featured SortBy = "Featured"
	// NewestArrivals constant to sort by NewestArrivals
	NewestArrivals SortBy = "NewestArrivals"
	// PriceHighToLow constant to sort by PriceHighToLow
	PriceHighToLow SortBy = "Price:HighToLow"
	// PriceLowToHigh constant to sort by PriceLowToHigh
	PriceLowToHigh SortBy = "Price:LowToHigh"
	// Relevance constant to sort by Relevance
	Relevance SortBy = "Relevance"
)

// Properties custom type to hold property pairs
type Properties map[string]string

// Add adds property key value pair
func (p Properties) Add(key, value string) {
	p[key] = value
}

// Remove deletes a property value by its key
func (p Properties) Remove(key string) {
	delete(p, key)
}

// Exists checks if key exists
func (p Properties) Exists(key string) bool {
	_, ok := p[key]
	return ok
}

// Length returns the count of key value pairs in Properties typed variable
func (p Properties) Length() int {
	return len(p)
}
