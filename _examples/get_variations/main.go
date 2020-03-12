package main

import (
	"context"
	"fmt"
	"os"

	"github.com/utekaravinash/gopaapi5"
	"github.com/utekaravinash/gopaapi5/api"
)

const tmpl = `Title: %s
URL: %s
----------------------
`

func main() {

	// Get Access Key, Secret Key and Associate Tag from the environment variables
	accessKey := os.Getenv("PA_ACCESS_KEY")
	secretKey := os.Getenv("PA_SECRET_KEY")
	associateTag := os.Getenv("PA_ASSOCIATE_TAG")

	// Initiate gopaapi5 Client
	client, err := gopaapi5.NewClient(accessKey, secretKey, associateTag, api.UnitedStates)
	if err != nil {
		panic(err)
	}

	// Construct request parameters for GetVariations operation
	params := api.GetVariationsParams{
		ASIN: "B01MD19OI2",
		Resources: []api.Resource{
			api.BrowseNodeInfoBrowseNodes,
			api.BrowseNodeInfoBrowseNodesAncestor,
			api.BrowseNodeInfoBrowseNodesSalesRank,
			api.BrowseNodeInfoWebsiteSalesRank,
			api.CustomerReviewsCount,
			api.CustomerReviewsStarRating,
			api.ImagesPrimarySmall,
			api.ImagesPrimaryMedium,
			api.ImagesPrimaryLarge,
			api.ImagesVariantsSmall,
			api.ImagesVariantsMedium,
			api.ImagesVariantsLarge,
			api.ItemInfoByLineInfo,
			api.ItemInfoContentInfo,
			api.ItemInfoContentRating,
			api.ItemInfoClassifications,
			api.ItemInfoExternalIds,
			api.ItemInfoFeatures,
			api.ItemInfoManufactureInfo,
			api.ItemInfoProductInfo,
			api.ItemInfoTechnicalInfo,
			api.ItemInfoTitle,
			api.ItemInfoTradeInInfo,
			api.OffersListingsAvailabilityMaxOrderQuantity,
			api.OffersListingsAvailabilityMessage,
			api.OffersListingsAvailabilityMinOrderQuantity,
			api.OffersListingsAvailabilityType,
			api.OffersListingsCondition,
			api.OffersListingsConditionSubCondition,
			api.OffersListingsDeliveryInfoIsAmazonFulfilled,
			api.OffersListingsDeliveryInfoIsFreeShippingEligible,
			api.OffersListingsDeliveryInfoIsPrimeEligible,
			api.OffersListingsDeliveryInfoShippingCharges,
			api.OffersListingsIsBuyBoxWinner,
			api.OffersListingsLoyaltyPointsPoints,
			api.OffersListingsMerchantInfo,
			api.OffersListingsPrice,
			api.OffersListingsProgramEligibilityIsPrimeExclusive,
			api.OffersListingsProgramEligibilityIsPrimePantry,
			api.OffersListingsPromotions,
			api.OffersListingsSavingBasis,
			api.OffersSummariesHighestPrice,
			api.OffersSummariesLowestPrice,
			api.OffersSummariesOfferCount,
			api.ParentASIN,
			api.RentalOffersListingsAvailabilityMaxOrderQuantity,
			api.RentalOffersListingsAvailabilityMessage,
			api.RentalOffersListingsAvailabilityMinOrderQuantity,
			api.RentalOffersListingsAvailabilityType,
			api.RentalOffersListingsBasePrice,
			api.RentalOffersListingsCondition,
			api.RentalOffersListingsConditionSubCondition,
			api.RentalOffersListingsDeliveryInfoIsAmazonFulfilled,
			api.RentalOffersListingsDeliveryInfoIsFreeShippingEligible,
			api.RentalOffersListingsDeliveryInfoIsPrimeEligible,
			api.RentalOffersListingsDeliveryInfoShippingCharges,
			api.RentalOffersListingsMerchantInfo,
			api.VariationSummaryPriceHighestPrice,
			api.VariationSummaryPriceLowestPrice,
			api.VariationSummaryVariationDimension,
		},
	}

	// Call GetVariations operation
	response, err := client.GetVariations(context.Background(), &params)
	if err != nil {
		panic(err)
	}

	// Loop over items in response
	for _, item := range response.VariationsResult.Items {
		fmt.Printf(tmpl, item.ItemInfo.Title.DisplayValue, item.DetailPageURL)
	}
}
