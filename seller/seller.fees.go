package seller

import "fmt"

type FeesEstimateParams struct {
	FeesEstimateRequest FeesEstimateRequest `json:"feesEstimateRequest"`
}

type FeesEstimateRequest struct {
	FeesEstimate
}

type FeesEstimate struct {
	MarketplaceId     string `json:"MarketplaceId"`
	IsAmazonFulfilled string `json:"IsAmazonFulfilled"`
	Identifier        string `json:"Identifier"`
}

type PriceToEstimateFees struct {
	ListingPrice MoneyType `json:"ListingPrice"`
	Shipping     MoneyType `json:"Shipping,omitempty"`
	Points       Points    `json:"Points,omitempty"`
}

func GetFeesEstimate(sku string) {
	var action = fmt.Sprintf("/products/fees/v0/listings/%s/feesEstimate", sku)
	fmt.Println(action)
	return
}
