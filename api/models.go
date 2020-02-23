package api

// Error represents Error object in json response
type Error struct {
	Type    string `json:"__type,omitempty"`
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// ItemsResult represents ItemsResult object in json response
type ItemsResult struct {
	Items []Item `json:"Items,omitempty"`
}

// BrowseNodesResult represents BrowseNodesResult object in json response
type BrowseNodesResult struct {
	BrowseNodes []BrowseNode `json:"BrowseNodes,omitempty"`
}

// VariationsResult represents VariationsResult object in json response
type VariationsResult struct {
	Items            []Item           `json:"Items,omitempty"`
	VariationSummary VariationSummary `json:"VariationSummary,omitempty"`
}

// SearchResult represents SearchResult object in json response
type SearchResult struct {
	TotalResultCount  int                    `json:"TotalResultCount,omitempty"`
	SearchURL         string                 `json:"SearchURL,omitempty"`
	Items             []Item                 `json:"Items,omitempty"`
	SearchRefinements SearchRefinementsModel `json:"SearchRefinements,omitempty"`
}

// SearchRefinementsModel represents SearchRefinementsModel object in json response
type SearchRefinementsModel struct {
	BrowseNode       Refinement   `json:"BrowseNode,omitempty"`
	OtherRefinements []Refinement `json:"OtherRefinements,omitempty"`
	SearchIndex      Refinement   `json:"SearchIndex,omitempty"`
}

// Refinement represents Refinement object in json response
type Refinement struct {
	Bins        []RefinementBin `json:"Bins,omitempty"`
	DisplayName string          `json:"DisplayName,omitempty"`
	Id          string          `json:"Id,omitempty"`
}

// RefinementBin represents RefinementBin object in json response
type RefinementBin struct {
	DisplayName string `json:"DisplayName,omitempty"`
	Id          string `json:"Id,omitempty"`
}

// Item represents Item object in json response
type Item struct {
	ASIN                string               `json:"ASIN,omitempty"`
	BrowseNodeInfo      BrowseNodeInfo       `json:"BrowseNodeInfo,omitempty"`
	DetailPageURL       string               `json:"DetailPageURL,omitempty"`
	Images              Images               `json:"Images,omitempty"`
	ItemInfo            ItemInfo             `json:"ItemInfo,omitempty"`
	Offers              Offers               `json:"Offers,omitempty"`
	ParentASIN          string               `json:"ParentASIN,omitempty"`
	RentalOffers        RentalOffers         `json:"RentalOffers,omitempty"`
	Score               float32              `json:"Score,omitempty"`
	VariationAttributes []VariationAttribute `json:"VariationAttributes,omitempty"`
}

// VariationSummary represents VariationSummary object in json response
type VariationSummary struct {
	PageCount           int                  `json:"PageCount,omitempty"`
	Price               Price                `json:"Price,omitempty"`
	VariationCount      int                  `json:"VariationCount,omitempty"`
	VariationDimensions []VariationDimension `json:"VariationDimensions,omitempty"`
}

// Price represents Price object in json response
type Price struct {
	HighestPrice OfferPrice `json:"HighestPrice,omitempty"`
	LowestPrice  OfferPrice `json:"LowestPrice,omitempty"`
}

// VariationDimension represents VariationDimension object in json response
type VariationDimension struct {
	DisplayName string   `json:"DisplayName,omitempty"`
	Locale      string   `json:"Locale,omitempty"`
	Name        string   `json:"Name,omitempty"`
	Values      []string `json:"Values,omitempty"`
}

// BrowseNodeInfo represents BrowseNodeInfo object in json response
type BrowseNodeInfo struct {
	BrowseNodes      []BrowseNode     `json:"BrowseNodes,omitempty"`
	WebsiteSalesRank WebsiteSalesRank `json:"WebsiteSalesRank,omitempty"`
}

// RentalOffers represents RentalOffers object in json response
type RentalOffers struct {
	Listings []RentalOfferListing `json:"Listings,omitempty"`
}

// RentalOfferListing represents RentalOfferListing object in json response
type RentalOfferListing struct {
	Availability OfferAvailability `json:"Availability,omitempty"`
	BasePrice    DurationPrice     `json:"BasePrice,omitempty"`
	Condition    OfferCondition    `json:"Condition,omitempty"`
	DeliveryInfo OfferDeliveryInfo `json:"DeliveryInfo,omitempty"`
	Id           string            `json:"Id,omitempty"`
	MerchantInfo OfferMerchantInfo `json:"MerchantInfo,omitempty"`
}

// DurationPrice represents DurationPrice object in json response
type DurationPrice struct {
	Price    OfferPrice         `json:"Price,omitempty"`
	Duration UnitBasedAttribute `json:"Duration,omitempty"`
}

// VariationAttribute represents VariationAttribute object in json response
type VariationAttribute struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}

// BrowseNode represents BrowseNode object in json response
type BrowseNode struct {
	Ancestor        BrowseNodeAncestor `json:"Ancestor,omitempty"`
	Children        []BrowseNodeChild  `json:"Children,omitempty"`
	ContextFreeName string             `json:"ContextFreeName,omitempty"`
	DisplayName     string             `json:"DisplayName,omitempty"`
	Id              string             `json:"Id,omitempty"`
	IsRoot          bool               `json:"IsRoot,omitempty"`
	SalesRank       int                `json:"SalesRank,omitempty"`
}

// BrowseNodeAncestor represents BrowseNodeAncestor object in json response
type BrowseNodeAncestor struct {
	Ancestor        *BrowseNodeAncestor `json:"Ancestor,omitempty"`
	ContextFreeName string              `json:"ContextFreeName,omitempty"`
	DisplayName     string              `json:"DisplayName,omitempty"`
	Id              string              `json:"Id,omitempty"`
}

// BrowseNodeChild represents BrowseNodeChild object in json response
type BrowseNodeChild struct {
	ContextFreeName string `json:"ContextFreeName,omitempty"`
	DisplayName     string `json:"DisplayName,omitempty"`
	Id              string `json:"Id,omitempty"`
}

// WebsiteSalesRank represents WebsiteSalesRank object in json response
type WebsiteSalesRank struct {
	ContextFreeName string `json:"ContextFreeName,omitempty"`
	DisplayName     string `json:"DisplayName,omitempty"`
	Id              string `json:"Id,omitempty"`
	SalesRank       int    `json:"SalesRank,omitempty"`
}

// Images represents Images object in json response
type Images struct {
	Primary  ImageType   `json:"Primary,omitempty"`
	Variants []ImageType `json:"Variants,omitempty"`
}

// ImageType represents ImageType object in json response
type ImageType struct {
	Small  ImageSize `json:"Small,omitempty"`
	Medium ImageSize `json:"Medium,omitempty"`
	Large  ImageSize `json:"Large,omitempty"`
}

// ImageSize represents ImageSize object in json response
type ImageSize struct {
	URL    string `json:"URL,omitempty"`
	Height int    `json:"Height,omitempty"`
	Width  int    `json:"Width,omitempty"`
}

// ItemInfo represents ItemInfo object in json response
type ItemInfo struct {
	ByLineInfo      ByLineInfo                  `json:"ByLineInfo,omitempty"`
	Classifications Classifications             `json:"Classifications,omitempty"`
	ContentInfo     ContentInfo                 `json:"ContentInfo,omitempty"`
	ContentRating   ContentRating               `json:"ContentRating,omitempty"`
	ExternalIds     ExternalIds                 `json:"ExternalIds,omitempty"`
	Features        MultiValuedAttribute        `json:"Features,omitempty"`
	ManufactureInfo ManufactureInfo             `json:"ManufactureInfo,omitempty"`
	ProductInfo     ProductInfo                 `json:"ProductInfo,omitempty"`
	TechnicalInfo   TechnicalInfo               `json:"TechnicalInfo,omitempty"`
	Title           SingleStringValuedAttribute `json:"Title,omitempty"`
	TradeInInfo     TradeInInfo                 `json:"TradeInInfo,omitempty"`
}

// ByLineInfo represents ByLineInfo object in json response
type ByLineInfo struct {
	Brand        SingleStringValuedAttribute `json:"Brand,omitempty"`
	Contributors []Contributor               `json:"Contributors,omitempty"`
	Manufacturer SingleStringValuedAttribute `json:"Manufacturer,omitempty"`
}

// SingleStringValuedAttribute represents SingleStringValuedAttribute object in json response
type SingleStringValuedAttribute struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

// Contributor represents Contributor object in json response
type Contributor struct {
	Locale string `json:"Locale,omitempty"`
	Name   string `json:"Name,omitempty"`
	Role   string `json:"Role,omitempty"`
}

// Classifications represents Classifications object in json response
type Classifications struct {
	Binding      SingleStringValuedAttribute `json:"Binding,omitempty"`
	ProductGroup SingleStringValuedAttribute `json:"ProductGroup,omitempty"`
}

// ContentInfo represents ContentInfo object in json response
type ContentInfo struct {
	Edition         SingleStringValuedAttribute  `json:"Edition,omitempty"`
	Languages       Languages                    `json:"Languages,omitempty"`
	PagesCount      SingleIntegerValuedAttribute `json:"PagesCount,omitempty"`
	PublicationDate SingleStringValuedAttribute  `json:"PublicationDate,omitempty"`
}

// Languages represents Languages object in json response
type Languages struct {
	DisplayValues []LanguageType `json:"DisplayValues,omitempty"`
	Label         string         `json:"Label,omitempty"`
	Locale        string         `json:"Locale,omitempty"`
}

// LanguageType represents LanguageType object in json response
type LanguageType struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Type         string `json:"Type,omitempty"`
}

// SingleIntegerValuedAttribute represents SingleIntegerValuedAttribute object in json response
type SingleIntegerValuedAttribute struct {
	DisplayValue int    `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

// ContentRating represents ContentRating object in json response
type ContentRating struct {
	AudienceRating SingleStringValuedAttribute `json:"AudienceRating,omitempty"`
}

// ExternalIds represents ExternalIds object in json response
type ExternalIds struct {
	EANs  MultiValuedAttribute `json:"EANs,omitempty"`
	ISBNs MultiValuedAttribute `json:"ISBNs,omitempty"`
	UPCs  MultiValuedAttribute `json:"UPCs,omitempty"`
}

// MultiValuedAttribute represents MultiValuedAttribute object in json response
type MultiValuedAttribute struct {
	DisplayValues []string `json:"DisplayValues,omitempty"`
	Label         string   `json:"Label,omitempty"`
	Locale        string   `json:"Locale,omitempty"`
}

// ManufactureInfo represents ManufactureInfo object in json response
type ManufactureInfo struct {
	ItemPartNumber SingleStringValuedAttribute `json:"ItemPartNumber,omitempty"`
	Model          SingleStringValuedAttribute `json:"Model,omitempty"`
	Warranty       SingleStringValuedAttribute `json:"Warranty,omitempty"`
}

// ProductInfo represents ProductInfo object in json response
type ProductInfo struct {
	Color          SingleStringValuedAttribute  `json:"Color,omitempty"`
	IsAdultProduct SingleBooleanValuedAttribute `json:"IsAdultProduct,omitempty"`
	ItemDimensions DimensionBasedAttribute      `json:"ItemDimensions,omitempty"`
	ReleaseDate    SingleStringValuedAttribute  `json:"ReleaseDate,omitempty"`
	Size           SingleStringValuedAttribute  `json:"Size,omitempty"`
	UnitCount      SingleIntegerValuedAttribute `json:"UnitCount,omitempty"`
}

// SingleBooleanValuedAttribute represents SingleBooleanValuedAttribute object in json response
type SingleBooleanValuedAttribute struct {
	DisplayValue bool   `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

// DimensionBasedAttribute represents DimensionBasedAttribute object in json response
type DimensionBasedAttribute struct {
	Height UnitBasedAttribute `json:"Height,omitempty"`
	Length UnitBasedAttribute `json:"Length,omitempty"`
	Weight UnitBasedAttribute `json:"Weight,omitempty"`
	Width  UnitBasedAttribute `json:"Width,omitempty"`
}

// UnitBasedAttribute represents UnitBasedAttribute object in json response
type UnitBasedAttribute struct {
	DisplayValue float32 `json:"DisplayValue,omitempty"`
	Label        string  `json:"Label,omitempty"`
	Locale       string  `json:"Locale,omitempty"`
	Unit         string  `json:"Unit,omitempty"`
}

// TechnicalInfo represents TechnicalInfo object in json response
type TechnicalInfo struct {
	Formats MultiValuedAttribute `json:"Formats,omitempty"`
}

// TradeInInfo represents TradeInInfo object in json response
type TradeInInfo struct {
	IsEligibleForTradeIn bool         `json:"IsEligibleForTradeIn,omitempty"`
	Price                TradeInPrice `json:"Price,omitempty"`
}

// TradeInPrice represents TradeInPrice object in json response
type TradeInPrice struct {
	Amount        float32 `json:"Amount,omitempty"`
	Currency      string  `json:"Currency,omitempty"`
	DisplayAmount string  `json:"DisplayAmount,omitempty"`
}

// Offers represents Offers object in json response
type Offers struct {
	Listings  []OfferListing `json:"Listings,omitempty"`
	Summaries []OfferSummary `json:"Summaries,omitempty"`
}

// OfferListing represents OfferListing object in json response
type OfferListing struct {
	Availability       OfferAvailability       `json:"Availability,omitempty"`
	Condition          OfferCondition          `json:"Condition,omitempty"`
	DeliveryInfo       OfferDeliveryInfo       `json:"DeliveryInfo,omitempty"`
	Id                 string                  `json:"Id,omitempty"`
	IsBuyBoxWinner     bool                    `json:"IsBuyBoxWinner,omitempty"`
	LoyaltyPoints      OfferLoyaltyPoints      `json:"LoyaltyPoints,omitempty"`
	MerchantInfo       OfferMerchantInfo       `json:"MerchantInfo,omitempty"`
	Price              OfferPrice              `json:"Price,omitempty"`
	ProgramEligibility OfferProgramEligibility `json:"ProgramEligibility,omitempty"`
	Promotions         []OfferPromotion        `json:"Promotions,omitempty"`
	SavingBasis        OfferPrice              `json:"SavingBasis,omitempty"`
	ViolatesMAP        bool                    `json:"ViolatesMAP,omitempty"`
}

// OfferAvailability represents OfferAvailability object in json response
type OfferAvailability struct {
	MaxOrderQuantity int    `json:"MaxOrderQuantity,omitempty"`
	Message          string `json:"Message,omitempty"`
	MinOrderQuantity int    `json:"MinOrderQuantity,omitempty"`
	Type             string `json:"Type,omitempty"`
}

// OfferDeliveryInfo represents OfferDeliveryInfo object in json response
type OfferDeliveryInfo struct {
	IsAmazonFulfilled      bool                  `json:"IsAmazonFulfilled,omitempty"`
	IsFreeShippingEligible bool                  `json:"IsFreeShippingEligible,omitempty"`
	IsPrimeEligible        bool                  `json:"IsPrimeEligible,omitempty"`
	ShippingCharges        []OfferShippingCharge `json:"ShippingCharges,omitempty"`
}

// OfferLoyaltyPoints represents OfferLoyaltyPoints object in json response
type OfferLoyaltyPoints struct {
	Points int `json:"Points,omitempty"`
}

// OfferMerchantInfo represents OfferMerchantInfo object in json response
type OfferMerchantInfo struct {
	DefaultShippingCountry string `json:"DefaultShippingCountry,omitempty"`
	Id                     string `json:"Id,omitempty"`
	Name                   string `json:"Name,omitempty"`
}

// OfferProgramEligibility represents OfferProgramEligibility object in json response
type OfferProgramEligibility struct {
	IsPrimeExclusive bool `json:"IsPrimeExclusive,omitempty"`
	IsPrimePantry    bool `json:"IsPrimePantry,omitempty"`
}

// OfferPromotion represents OfferPromotion object in json response
type OfferPromotion struct {
	Amount          float32 `json:"Amount,omitempty"`
	Currency        string  `json:"Currency,omitempty"`
	DiscountPercent int     `json:"DiscountPercent,omitempty"`
	DisplayAmount   string  `json:"DisplayAmount,omitempty"`
	PricePerUnit    float32 `json:"PricePerUnit,omitempty"`
	Type            string  `json:"Type,omitempty"`
}

// OfferShippingCharge represents OfferShippingCharge object in json response
type OfferShippingCharge struct {
	Amount             float32 `json:"Amount,omitempty"`
	Currency           string  `json:"Currency,omitempty"`
	DisplayAmount      string  `json:"DisplayAmount,omitempty"`
	IsRateTaxInclusive bool    `json:"IsRateTaxInclusive,omitempty"`
	Type               string  `json:"Type,omitempty"`
}

// OfferSummary represents OfferSummary object in json response
type OfferSummary struct {
	Condition    OfferCondition `json:"Condition,omitempty"`
	HighestPrice OfferPrice     `json:"HighestPrice,omitempty"`
	LowestPrice  OfferPrice     `json:"LowestPrice,omitempty"`
	OfferCount   int            `json:"OfferCount,omitempty"`
}

// OfferCondition represents OfferCondition object in json response
type OfferCondition struct {
	DisplayValue string            `json:"DisplayValue,omitempty"`
	Label        string            `json:"Label,omitempty"`
	Locale       string            `json:"Locale,omitempty"`
	Value        string            `json:"Value,omitempty"`
	SubCondition OfferSubCondition `json:"SubCondition,omitempty"`
}

// OfferSubCondition represents OfferSubCondition object in json response
type OfferSubCondition struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
	Value        string `json:"Value,omitempty"`
}

// OfferPrice represents OfferPrice object in json response
type OfferPrice struct {
	Amount        float32      `json:"Amount,omitempty"`
	Currency      string       `json:"Currency,omitempty"`
	DisplayAmount string       `json:"DisplayAmount,omitempty"`
	PricePerUnit  float32      `json:"PricePerUnit,omitempty"`
	Savings       OfferSavings `json:"Savings,omitempty"`
}

// OfferSavings represents OfferSavings object in json response
type OfferSavings struct {
	Amount        float32 `json:"Amount,omitempty"`
	Currency      string  `json:"Currency,omitempty"`
	DisplayAmount string  `json:"DisplayAmount,omitempty"`
	Percentage    int     `json:"Percentage,omitempty"`
	PricePerUnit  float32 `json:"PricePerUnit,omitempty"`
}
