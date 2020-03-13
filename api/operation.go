package api

import (
	"fmt"
)

// Operation custom type for PA API operations
type Operation string

// ErrInvalidResource is thrown when a resource is not compatible with an operation
type ErrInvalidResource struct {
	resource  Resource
	operation Operation
}

func (e ErrInvalidResource) Error() string {
	return fmt.Sprintf(`Invalid resource "%s" for operation "%s"`, e.resource, e.operation)
}

const (
	// GetBrowseNodes constant for GetBrowseNodes operation
	GetBrowseNodes Operation = "GetBrowseNodes"
	// GetItems constant for GetItems operation
	GetItems Operation = "GetItems"
	// GetVariations constant for GetVariations operation
	GetVariations Operation = "GetVariations"
	// SearchItems constant for SearchItems operation
	SearchItems Operation = "SearchItems"
)

// resourceOperationsMap maps resources to their valid operations
var resourceOperationsMap = map[Resource][]Operation{
	BrowseNodesAncestor:                                    {GetBrowseNodes},
	BrowseNodesChildren:                                    {GetBrowseNodes},
	BrowseNodeInfoBrowseNodes:                              {GetItems, GetVariations, SearchItems},
	BrowseNodeInfoBrowseNodesAncestor:                      {GetItems, GetVariations, SearchItems},
	BrowseNodeInfoBrowseNodesSalesRank:                     {GetItems, GetVariations, SearchItems},
	BrowseNodeInfoWebsiteSalesRank:                         {GetItems, GetVariations, SearchItems},
	CustomerReviewsCount:                                   {GetItems, GetVariations, SearchItems},
	CustomerReviewsStarRating:                              {GetItems, GetVariations, SearchItems},
	ImagesPrimarySmall:                                     {GetItems, GetVariations, SearchItems},
	ImagesPrimaryMedium:                                    {GetItems, GetVariations, SearchItems},
	ImagesPrimaryLarge:                                     {GetItems, GetVariations, SearchItems},
	ImagesVariantsSmall:                                    {GetItems, GetVariations, SearchItems},
	ImagesVariantsMedium:                                   {GetItems, GetVariations, SearchItems},
	ImagesVariantsLarge:                                    {GetItems, GetVariations, SearchItems},
	ItemInfoByLineInfo:                                     {GetItems, GetVariations, SearchItems},
	ItemInfoContentInfo:                                    {GetItems, GetVariations, SearchItems},
	ItemInfoContentRating:                                  {GetItems, GetVariations, SearchItems},
	ItemInfoClassifications:                                {GetItems, GetVariations, SearchItems},
	ItemInfoExternalIds:                                    {GetItems, GetVariations, SearchItems},
	ItemInfoFeatures:                                       {GetItems, GetVariations, SearchItems},
	ItemInfoManufactureInfo:                                {GetItems, GetVariations, SearchItems},
	ItemInfoProductInfo:                                    {GetItems, GetVariations, SearchItems},
	ItemInfoTechnicalInfo:                                  {GetItems, GetVariations, SearchItems},
	ItemInfoTitle:                                          {GetItems, GetVariations, SearchItems},
	ItemInfoTradeInInfo:                                    {GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMaxOrderQuantity:             {GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMessage:                      {GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMinOrderQuantity:             {GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityType:                         {GetItems, GetVariations, SearchItems},
	OffersListingsCondition:                                {GetItems, GetVariations, SearchItems},
	OffersListingsConditionSubCondition:                    {GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsAmazonFulfilled:            {GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsFreeShippingEligible:       {GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsPrimeEligible:              {GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoShippingCharges:              {GetItems, GetVariations, SearchItems},
	OffersListingsIsBuyBoxWinner:                           {GetItems, GetVariations, SearchItems},
	OffersListingsLoyaltyPointsPoints:                      {GetItems, GetVariations, SearchItems},
	OffersListingsMerchantInfo:                             {GetItems, GetVariations, SearchItems},
	OffersListingsPrice:                                    {GetItems, GetVariations, SearchItems},
	OffersListingsProgramEligibilityIsPrimeExclusive:       {GetItems, GetVariations, SearchItems},
	OffersListingsProgramEligibilityIsPrimePantry:          {GetItems, GetVariations, SearchItems},
	OffersListingsPromotions:                               {GetItems, GetVariations, SearchItems},
	OffersListingsSavingBasis:                              {GetItems, GetVariations, SearchItems},
	OffersSummariesHighestPrice:                            {GetItems, GetVariations, SearchItems},
	OffersSummariesLowestPrice:                             {GetItems, GetVariations, SearchItems},
	OffersSummariesOfferCount:                              {GetItems, GetVariations, SearchItems},
	ParentASIN:                                             {GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMaxOrderQuantity:       {GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMessage:                {GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMinOrderQuantity:       {GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityType:                   {GetItems, GetVariations, SearchItems},
	RentalOffersListingsBasePrice:                          {GetItems, GetVariations, SearchItems},
	RentalOffersListingsCondition:                          {GetItems, GetVariations, SearchItems},
	RentalOffersListingsConditionSubCondition:              {GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsAmazonFulfilled:      {GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsFreeShippingEligible: {GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsPrimeEligible:        {GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoShippingCharges:        {GetItems, GetVariations, SearchItems},
	RentalOffersListingsMerchantInfo:                       {GetItems, GetVariations, SearchItems},
	VariationSummaryPriceHighestPrice:                      {GetVariations},
	VariationSummaryPriceLowestPrice:                       {GetVariations},
	VariationSummaryVariationDimension:                     {GetVariations},
	SearchRefinements:                                      {SearchItems},
}

// Validate validates resources for an operation
func (o Operation) Validate(resources []Resource) error {
	for _, r := range resources {
		validOperations := resourceOperationsMap[r]
		if !existsInOperations(o, validOperations) {
			return ErrInvalidResource{r, o}
		}
	}

	return nil
}

// existsInOperations checks if an Operation exists in a slice of Operations
func existsInOperations(operation Operation, operations []Operation) bool {
	for _, o := range operations {
		if o == operation {
			return true
		}
	}

	return false
}
