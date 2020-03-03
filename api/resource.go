package api

// Resource determine what information will be returned in the API response
type Resource string

const (
	// BrowseNodesAncestor constant for BrowseNodes.Ancestor resource
	BrowseNodesAncestor Resource = "BrowseNodes.Ancestor"
	// BrowseNodesChildren constant for BrowseNodes.Children resource
	BrowseNodesChildren Resource = "BrowseNodes.Children"
	// BrowseNodeInfoBrowseNodes constant for BrowseNodeInfo.BrowseNodes resource
	BrowseNodeInfoBrowseNodes Resource = "BrowseNodeInfo.BrowseNodes"
	// BrowseNodeInfoBrowseNodesAncestor constant for BrowseNodeInfo.BrowseNodes.Ancestor resource
	BrowseNodeInfoBrowseNodesAncestor Resource = "BrowseNodeInfo.BrowseNodes.Ancestor"
	// BrowseNodeInfoBrowseNodesSalesRank constant for BrowseNodeInfo.BrowseNodes.SalesRank resource
	BrowseNodeInfoBrowseNodesSalesRank Resource = "BrowseNodeInfo.BrowseNodes.SalesRank"
	// BrowseNodeInfoWebsiteSalesRank constant for BrowseNodeInfo.WebsiteSalesRank resource
	BrowseNodeInfoWebsiteSalesRank Resource = "BrowseNodeInfo.WebsiteSalesRank"
	// CustomerReviewsCount constant for CustomerReviews.Count resource
	CustomerReviewsCount Resource = "CustomerReviews.Count"
	// CustomerReviewsStarRating constant for CustomerReviews.StarRating resource
	CustomerReviewsStarRating Resource = "CustomerReviews.StarRating"
	// ImagesPrimarySmall constant for Images.Primary.Small resource
	ImagesPrimarySmall Resource = "Images.Primary.Small"
	// ImagesPrimaryMedium constant for Images.Primary.Medium resource
	ImagesPrimaryMedium Resource = "Images.Primary.Medium"
	// ImagesPrimaryLarge constant for Images.Primary.Large resource
	ImagesPrimaryLarge Resource = "Images.Primary.Large"
	// ImagesVariantsSmall constant for Images.Variants.Small resource
	ImagesVariantsSmall Resource = "Images.Variants.Small"
	// ImagesVariantsMedium constant for Images.Variants.Medium resource
	ImagesVariantsMedium Resource = "Images.Variants.Medium"
	// ImagesVariantsLarge constant for Images.Variants.Large resource
	ImagesVariantsLarge Resource = "Images.Variants.Large"
	// ItemInfoByLineInfo constant for ItemInfo.ByLineInfo resource
	ItemInfoByLineInfo Resource = "ItemInfo.ByLineInfo"
	// ItemInfoContentInfo constant for ItemInfo.ContentInfo resource
	ItemInfoContentInfo Resource = "ItemInfo.ContentInfo"
	// ItemInfoContentRating constant for ItemInfo.ContentRating resource
	ItemInfoContentRating Resource = "ItemInfo.ContentRating"
	// ItemInfoClassifications constant for ItemInfo.Classifications resource
	ItemInfoClassifications Resource = "ItemInfo.Classifications"
	// ItemInfoExternalIds constant for ItemInfo.ExternalIds resource
	ItemInfoExternalIds Resource = "ItemInfo.ExternalIds"
	// ItemInfoFeatures constant for ItemInfo.Features resource
	ItemInfoFeatures Resource = "ItemInfo.Features"
	// ItemInfoManufactureInfo constant for ItemInfo.ManufactureInfo resource
	ItemInfoManufactureInfo Resource = "ItemInfo.ManufactureInfo"
	// ItemInfoProductInfo constant for ItemInfo.ProductInfo resource
	ItemInfoProductInfo Resource = "ItemInfo.ProductInfo"
	// ItemInfoTechnicalInfo constant for ItemInfo.TechnicalInfo resource
	ItemInfoTechnicalInfo Resource = "ItemInfo.TechnicalInfo"
	// ItemInfoTitle constant for ItemInfo.Title resource
	ItemInfoTitle Resource = "ItemInfo.Title"
	// ItemInfoTradeInInfo constant for ItemInfo.TradeInInfo resource
	ItemInfoTradeInInfo Resource = "ItemInfo.TradeInInfo"
	// OffersListingsAvailabilityMaxOrderQuantity constant for Offers.Listings.Availability.MaxOrderQuantity resource
	OffersListingsAvailabilityMaxOrderQuantity Resource = "Offers.Listings.Availability.MaxOrderQuantity"
	// OffersListingsAvailabilityMessage constant for Offers.Listings.Availability.Message resource
	OffersListingsAvailabilityMessage Resource = "Offers.Listings.Availability.Message"
	// OffersListingsAvailabilityMinOrderQuantity constant for Offers.Listings.Availability.MinOrderQuantity resource
	OffersListingsAvailabilityMinOrderQuantity Resource = "Offers.Listings.Availability.MinOrderQuantity"
	// OffersListingsAvailabilityType constant for Offers.Listings.Availability.Type resource
	OffersListingsAvailabilityType Resource = "Offers.Listings.Availability.Type"
	// OffersListingsCondition constant for Offers.Listings.Condition resource
	OffersListingsCondition Resource = "Offers.Listings.Condition"
	// OffersListingsConditionSubCondition constant for Offers.Listings.Condition.SubCondition resource
	OffersListingsConditionSubCondition Resource = "Offers.Listings.Condition.SubCondition"
	// OffersListingsDeliveryInfoIsAmazonFulfilled constant for Offers.Listings.DeliveryInfo.IsAmazonFulfilled resource
	OffersListingsDeliveryInfoIsAmazonFulfilled Resource = "Offers.Listings.DeliveryInfo.IsAmazonFulfilled"
	// OffersListingsDeliveryInfoIsFreeShippingEligible constant for Offers.Listings.DeliveryInfo.IsFreeShippingEligible resource
	OffersListingsDeliveryInfoIsFreeShippingEligible Resource = "Offers.Listings.DeliveryInfo.IsFreeShippingEligible"
	// OffersListingsDeliveryInfoIsPrimeEligible constant for Offers.Listings.DeliveryInfo.IsPrimeEligible resource
	OffersListingsDeliveryInfoIsPrimeEligible Resource = "Offers.Listings.DeliveryInfo.IsPrimeEligible"
	// OffersListingsDeliveryInfoShippingCharges constant for Offers.Listings.DeliveryInfo.ShippingCharges resource
	OffersListingsDeliveryInfoShippingCharges Resource = "Offers.Listings.DeliveryInfo.ShippingCharges"
	// OffersListingsIsBuyBoxWinner constant for Offers.Listings.IsBuyBoxWinner resource
	OffersListingsIsBuyBoxWinner Resource = "Offers.Listings.IsBuyBoxWinner"
	// OffersListingsLoyaltyPointsPoints constant for Offers.Listings.LoyaltyPoints.Points resource
	OffersListingsLoyaltyPointsPoints Resource = "Offers.Listings.LoyaltyPoints.Points"
	// OffersListingsMerchantInfo constant for Offers.Listings.MerchantInfo resource
	OffersListingsMerchantInfo Resource = "Offers.Listings.MerchantInfo"
	// OffersListingsPrice constant for Offers.Listings.Price resource
	OffersListingsPrice Resource = "Offers.Listings.Price"
	// OffersListingsProgramEligibilityIsPrimeExclusive constant for Offers.Listings.ProgramEligibility.IsPrimeExclusive resource
	OffersListingsProgramEligibilityIsPrimeExclusive Resource = "Offers.Listings.ProgramEligibility.IsPrimeExclusive"
	// OffersListingsProgramEligibilityIsPrimePantry constant for Offers.Listings.ProgramEligibility.IsPrimePantry resource
	OffersListingsProgramEligibilityIsPrimePantry Resource = "Offers.Listings.ProgramEligibility.IsPrimePantry"
	// OffersListingsPromotions constant for Offers.Listings.Promotions resource
	OffersListingsPromotions Resource = "Offers.Listings.Promotions"
	// OffersListingsSavingBasis constant for Offers.Listings.SavingBasis resource
	OffersListingsSavingBasis Resource = "Offers.Listings.SavingBasis"
	// OffersSummariesHighestPrice constant for Offers.Summaries.HighestPrice resource
	OffersSummariesHighestPrice Resource = "Offers.Summaries.HighestPrice"
	// OffersSummariesLowestPrice constant for Offers.Summaries.LowestPrice resource
	OffersSummariesLowestPrice Resource = "Offers.Summaries.LowestPrice"
	// OffersSummariesOfferCount constant for Offers.Summaries.OfferCount resource
	OffersSummariesOfferCount Resource = "Offers.Summaries.OfferCount"
	// ParentASIN constant for ParentASIN resource
	ParentASIN Resource = "ParentASIN"
	// RentalOffersListingsAvailabilityMaxOrderQuantity constant for RentalOffers.Listings.Availability.MaxOrderQuantity resource
	RentalOffersListingsAvailabilityMaxOrderQuantity Resource = "RentalOffers.Listings.Availability.MaxOrderQuantity"
	// RentalOffersListingsAvailabilityMessage constant for RentalOffers.Listings.Availability.Message resource
	RentalOffersListingsAvailabilityMessage Resource = "RentalOffers.Listings.Availability.Message"
	// RentalOffersListingsAvailabilityMinOrderQuantity constant for RentalOffers.Listings.Availability.MinOrderQuantity resource
	RentalOffersListingsAvailabilityMinOrderQuantity Resource = "RentalOffers.Listings.Availability.MinOrderQuantity"
	// RentalOffersListingsAvailabilityType constant for RentalOffers.Listings.Availability.Type resource
	RentalOffersListingsAvailabilityType Resource = "RentalOffers.Listings.Availability.Type"
	// RentalOffersListingsBasePrice constant for RentalOffers.Listings.BasePrice resource
	RentalOffersListingsBasePrice Resource = "RentalOffers.Listings.BasePrice"
	// RentalOffersListingsCondition constant for RentalOffers.Listings.Condition resource
	RentalOffersListingsCondition Resource = "RentalOffers.Listings.Condition"
	// RentalOffersListingsConditionSubCondition constant for RentalOffers.Listings.Condition.SubCondition resource
	RentalOffersListingsConditionSubCondition Resource = "RentalOffers.Listings.Condition.SubCondition"
	// RentalOffersListingsDeliveryInfoIsAmazonFulfilled constant for RentalOffers.Listings.DeliveryInfo.IsAmazonFulfilled resource
	RentalOffersListingsDeliveryInfoIsAmazonFulfilled Resource = "RentalOffers.Listings.DeliveryInfo.IsAmazonFulfilled"
	// RentalOffersListingsDeliveryInfoIsFreeShippingEligible constant for RentalOffers.Listings.DeliveryInfo.IsFreeShippingEligible resource
	RentalOffersListingsDeliveryInfoIsFreeShippingEligible Resource = "RentalOffers.Listings.DeliveryInfo.IsFreeShippingEligible"
	// RentalOffersListingsDeliveryInfoIsPrimeEligible constant for RentalOffers.Listings.DeliveryInfo.IsPrimeEligible resource
	RentalOffersListingsDeliveryInfoIsPrimeEligible Resource = "RentalOffers.Listings.DeliveryInfo.IsPrimeEligible"
	// RentalOffersListingsDeliveryInfoShippingCharges constant for RentalOffers.Listings.DeliveryInfo.ShippingCharges resource
	RentalOffersListingsDeliveryInfoShippingCharges Resource = "RentalOffers.Listings.DeliveryInfo.ShippingCharges"
	// RentalOffersListingsMerchantInfo constant for RentalOffers.Listings.MerchantInfo resource
	RentalOffersListingsMerchantInfo Resource = "RentalOffers.Listings.MerchantInfo"
	// VariationSummaryPriceHighestPrice constant for VariationSummary.Price.HighestPrice resource
	VariationSummaryPriceHighestPrice Resource = "VariationSummary.Price.HighestPrice"
	// VariationSummaryPriceLowestPrice constant for VariationSummary.Price.LowestPrice resource
	VariationSummaryPriceLowestPrice Resource = "VariationSummary.Price.LowestPrice"
	// VariationSummaryVariationDimension constant for VariationSummary.VariationDimension resource
	VariationSummaryVariationDimension Resource = "VariationSummary.VariationDimension"
	// SearchRefinements constant for SearchRefinements resource
	SearchRefinements Resource = "SearchRefinements"
)
