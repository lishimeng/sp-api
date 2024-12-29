package common

// ItemClassification Classification type associated with the Amazon catalog item.
type ItemClassification string

const (
	BaseProduct     ItemClassification = "BASE_PRODUCT"
	Other           ItemClassification = "OTHER"
	ProductBundle   ItemClassification = "PRODUCT_BUNDLE"
	VariationParent ItemClassification = "VARIATION_PARENT"
)
