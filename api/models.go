package api

type Error struct {
	Type    string `json:"__type,omitempty"`
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

type ItemsResult struct {
	Items []Item `json:"Items,omitempty"`
}

type BrowseNodesResult struct {
	BrowseNodes []BrowseNode `json:"BrowseNodes,omitempty"`
}

type VariationsResult struct {
	Items            []Item           `json:"Items,omitempty"`
	VariationSummary VariationSummary `json:"VariationSummary,omitempty"`
}

type SearchResult struct {
	TotalResultCount  int
	SearchURL         string
	Items             []Item
	SearchRefinements SearchRefinementsModel
}

type SearchRefinementsModel struct {
	BrowseNode       Refinement
	OtherRefinements []Refinement
	SearchIndex      Refinement
}

type Refinement struct {
	Bins        []RefinementBin
	DisplayName string
	Id          string
}

type RefinementBin struct {
	DisplayName string
	Id          string
}

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

type VariationSummary struct {
	PageCount           int
	Price               Price
	VariationCount      int
	VariationDimensions []VariationDimension
}

type Price struct {
	HighestPrice OfferPrice
	LowestPrice  OfferPrice
}

type VariationDimension struct {
	DisplayName string
	Locale      string
	Name        string
	Values      []string
}

type BrowseNodeInfo struct {
	BrowseNodes      []BrowseNode     `json:"BrowseNodes,omitempty"`
	WebsiteSalesRank WebsiteSalesRank `json:"WebsiteSalesRank,omitempty"`
}

type RentalOffers struct {
	Listings []RentalOfferListing `json:"Listings,omitempty"`
}

type RentalOfferListing struct {
	Availability OfferAvailability `json:"Availability,omitempty"`
	BasePrice    DurationPrice     `json:"BasePrice,omitempty"`
	Condition    OfferCondition    `json:"Condition,omitempty"`
	DeliveryInfo OfferDeliveryInfo `json:"DeliveryInfo,omitempty"`
	Id           string            `json:"Id,omitempty"`
	MerchantInfo OfferMerchantInfo `json:"MerchantInfo,omitempty"`
}

type DurationPrice struct {
	Price    OfferPrice         `json:"Price,omitempty"`
	Duration UnitBasedAttribute `json:"Duration,omitempty"`
}

type VariationAttribute struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}

type BrowseNode struct {
	Ancestor        BrowseNodeAncestor `json:"Ancestor,omitempty"`
	Children        []BrowseNodeChild  `json:"Children,omitempty"`
	ContextFreeName string             `json:"ContextFreeName,omitempty"`
	DisplayName     string             `json:"DisplayName,omitempty"`
	Id              string             `json:"Id,omitempty"`
	IsRoot          bool               `json:"IsRoot,omitempty"`
	SalesRank       int                `json:"SalesRank,omitempty"`
}

type BrowseNodeAncestor struct {
	Ancestor        *BrowseNodeAncestor `json:"Ancestor,omitempty"`
	ContextFreeName string              `json:"ContextFreeName,omitempty"`
	DisplayName     string              `json:"DisplayName,omitempty"`
	Id              string              `json:"Id,omitempty"`
}

type BrowseNodeChild struct {
	ContextFreeName string `json:"ContextFreeName,omitempty"`
	DisplayName     string `json:"DisplayName,omitempty"`
	Id              string `json:"Id,omitempty"`
}

type WebsiteSalesRank struct {
	ContextFreeName string `json:"ContextFreeName,omitempty"`
	DisplayName     string `json:"DisplayName,omitempty"`
	Id              string `json:"Id,omitempty"`
	SalesRank       int    `json:"SalesRank,omitempty"`
}

type Images struct {
	Primary  ImageType   `json:"Primary,omitempty"`
	Variants []ImageType `json:"Variants,omitempty"`
}

type ImageType struct {
	Small  ImageSize `json:"Small,omitempty"`
	Medium ImageSize `json:"Medium,omitempty"`
	Large  ImageSize `json:"Large,omitempty"`
}

type ImageSize struct {
	URL    string `json:"URL,omitempty"`
	Height int    `json:"Height,omitempty"`
	Width  int    `json:"Width,omitempty"`
}

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

type ByLineInfo struct {
	Brand        SingleStringValuedAttribute `json:"Brand,omitempty"`
	Contributors []Contributor               `json:"Contributors,omitempty"`
	Manufacturer SingleStringValuedAttribute `json:"Manufacturer,omitempty"`
}

type SingleStringValuedAttribute struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

type Contributor struct {
	Locale string `json:"Locale,omitempty"`
	Name   string `json:"Name,omitempty"`
	Role   string `json:"Role,omitempty"`
}

type Classifications struct {
	Binding      SingleStringValuedAttribute `json:"Binding,omitempty"`
	ProductGroup SingleStringValuedAttribute `json:"ProductGroup,omitempty"`
}

type ContentInfo struct {
	Edition         SingleStringValuedAttribute  `json:"Edition,omitempty"`
	Languages       Languages                    `json:"Languages,omitempty"`
	PagesCount      SingleIntegerValuedAttribute `json:"PagesCount,omitempty"`
	PublicationDate SingleStringValuedAttribute  `json:"PublicationDate,omitempty"`
}

type Languages struct {
	DisplayValues []LanguageType `json:"DisplayValues,omitempty"`
	Label         string         `json:"Label,omitempty"`
	Locale        string         `json:"Locale,omitempty"`
}

type LanguageType struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Type         string `json:"Type,omitempty"`
}

type SingleIntegerValuedAttribute struct {
	DisplayValue int    `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

type ContentRating struct {
	AudienceRating SingleStringValuedAttribute `json:"AudienceRating,omitempty"`
}

type ExternalIds struct {
	EANs  MultiValuedAttribute `json:"EANs,omitempty"`
	ISBNs MultiValuedAttribute `json:"ISBNs,omitempty"`
	UPCs  MultiValuedAttribute `json:"UPCs,omitempty"`
}

type MultiValuedAttribute struct {
	DisplayValues []string `json:"DisplayValues,omitempty"`
	Label         string   `json:"Label,omitempty"`
	Locale        string   `json:"Locale,omitempty"`
}

type ManufactureInfo struct {
	ItemPartNumber SingleStringValuedAttribute `json:"ItemPartNumber,omitempty"`
	Model          SingleStringValuedAttribute `json:"Model,omitempty"`
	Warranty       SingleStringValuedAttribute `json:"Warranty,omitempty"`
}

type ProductInfo struct {
	Color          SingleStringValuedAttribute  `json:"Color,omitempty"`
	IsAdultProduct SingleBooleanValuedAttribute `json:"IsAdultProduct,omitempty"`
	ItemDimensions DimensionBasedAttribute      `json:"ItemDimensions,omitempty"`
	ReleaseDate    SingleStringValuedAttribute  `json:"ReleaseDate,omitempty"`
	Size           SingleStringValuedAttribute  `json:"Size,omitempty"`
	UnitCount      SingleIntegerValuedAttribute `json:"UnitCount,omitempty"`
}

type SingleBooleanValuedAttribute struct {
	DisplayValue bool   `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
}

type DimensionBasedAttribute struct {
	Height UnitBasedAttribute `json:"Height,omitempty"`
	Length UnitBasedAttribute `json:"Length,omitempty"`
	Weight UnitBasedAttribute `json:"Weight,omitempty"`
	Width  UnitBasedAttribute `json:"Width,omitempty"`
}

type UnitBasedAttribute struct {
	DisplayValue float32 `json:"DisplayValue,omitempty"`
	Label        string  `json:"Label,omitempty"`
	Locale       string  `json:"Locale,omitempty"`
	Unit         string  `json:"Unit,omitempty"`
}

type TechnicalInfo struct {
	Formats MultiValuedAttribute `json:"Formats,omitempty"`
}

type TradeInInfo struct {
	IsEligibleForTradeIn bool         `json:"IsEligibleForTradeIn,omitempty"`
	Price                TradeInPrice `json:"Price,omitempty"`
}

type TradeInPrice struct {
	Amount        float32 `json:"Amount,omitempty"`
	Currency      string  `json:"Currency,omitempty"`
	DisplayAmount string  `json:"DisplayAmount,omitempty"`
}

type Offers struct {
	Listings  []OfferListing `json:"Listings,omitempty"`
	Summaries []OfferSummary `json:"Summaries,omitempty"`
}

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

type OfferAvailability struct {
	MaxOrderQuantity int    `json:"MaxOrderQuantity,omitempty"`
	Message          string `json:"Message,omitempty"`
	MinOrderQuantity int    `json:"MinOrderQuantity,omitempty"`
	Type             string `json:"Type,omitempty"`
}

type OfferDeliveryInfo struct {
	IsAmazonFulfilled      bool                  `json:"IsAmazonFulfilled,omitempty"`
	IsFreeShippingEligible bool                  `json:"IsFreeShippingEligible,omitempty"`
	IsPrimeEligible        bool                  `json:"IsPrimeEligible,omitempty"`
	ShippingCharges        []OfferShippingCharge `json:"ShippingCharges,omitempty"`
}

type OfferLoyaltyPoints struct {
	Points int `json:"Points,omitempty"`
}

type OfferMerchantInfo struct {
	DefaultShippingCountry string `json:"DefaultShippingCountry,omitempty"`
	Id                     string `json:"Id,omitempty"`
	Name                   string `json:"Name,omitempty"`
}

type OfferProgramEligibility struct {
	IsPrimeExclusive bool `json:"IsPrimeExclusive,omitempty"`
	IsPrimePantry    bool `json:"IsPrimePantry,omitempty"`
}

type OfferPromotion struct {
	Amount          float32 `json:"Amount,omitempty"`
	Currency        string  `json:"Currency,omitempty"`
	DiscountPercent int     `json:"DiscountPercent,omitempty"`
	DisplayAmount   string  `json:"DisplayAmount,omitempty"`
	PricePerUnit    float32 `json:"PricePerUnit,omitempty"`
	Type            string  `json:"Type,omitempty"`
}

type OfferShippingCharge struct {
	Amount             float32 `json:"Amount,omitempty"`
	Currency           string  `json:"Currency,omitempty"`
	DisplayAmount      string  `json:"DisplayAmount,omitempty"`
	IsRateTaxInclusive bool    `json:"IsRateTaxInclusive,omitempty"`
	Type               string  `json:"Type,omitempty"`
}

type OfferSummary struct {
	Condition    OfferCondition `json:"Condition,omitempty"`
	HighestPrice OfferPrice     `json:"HighestPrice,omitempty"`
	LowestPrice  OfferPrice     `json:"LowestPrice,omitempty"`
	OfferCount   int            `json:"OfferCount,omitempty"`
}

type OfferCondition struct {
	DisplayValue string            `json:"DisplayValue,omitempty"`
	Label        string            `json:"Label,omitempty"`
	Locale       string            `json:"Locale,omitempty"`
	Value        string            `json:"Value,omitempty"`
	SubCondition OfferSubCondition `json:"SubCondition,omitempty"`
}

type OfferSubCondition struct {
	DisplayValue string `json:"DisplayValue,omitempty"`
	Label        string `json:"Label,omitempty"`
	Locale       string `json:"Locale,omitempty"`
	Value        string `json:"Value,omitempty"`
}

type OfferPrice struct {
	Amount        float32      `json:"Amount,omitempty"`
	Currency      string       `json:"Currency,omitempty"`
	DisplayAmount string       `json:"DisplayAmount,omitempty"`
	PricePerUnit  float32      `json:"PricePerUnit,omitempty"`
	Savings       OfferSavings `json:"Savings,omitempty"`
}

type OfferSavings struct {
	Amount        float32 `json:"Amount,omitempty"`
	Currency      string  `json:"Currency,omitempty"`
	DisplayAmount string  `json:"DisplayAmount,omitempty"`
	Percentage    int     `json:"Percentage,omitempty"`
	PricePerUnit  float32 `json:"PricePerUnit,omitempty"`
}
