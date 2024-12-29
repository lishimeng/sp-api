package common

// IdentifiersType Type of product identifiers to search the Amazon catalog for.
// Note: Required when identifiers are provided.
type IdentifiersType string

const (
	ASIN   IdentifiersType = "ASIN"   // Amazon Standard Identification Number.
	EAN    IdentifiersType = "EAN"    // European Article Number.
	GTIN   IdentifiersType = "GTIN"   // Global Trade Item Number.
	ISBN   IdentifiersType = "ISBN"   // International Standard Book Number.
	JAN    IdentifiersType = "JAN"    // Japanese Article Number.
	MINSAN IdentifiersType = "MINSAN" // Minsan Code.
	SKU    IdentifiersType = "SKU"    // Stock Keeping Unit, a seller-specified identifier for an Amazon listing. Note: Must be accompanied by sellerId.
	UPC    IdentifiersType = "UPC"    // Universal Product Code.
)
