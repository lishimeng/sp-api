package common

type RelationshipType string //Type of relationship.

const (
	// Variation
	// The Amazon catalog item in the request is a variation parent
	// or variation child of the related item(s) indicated by ASIN.
	Variation RelationshipType = "VARIATION"
	// PackageHierarchy
	// The Amazon catalog item in the request is a package container
	// or is contained by the related item(s) indicated by ASIN.
	PackageHierarchy RelationshipType = "PACKAGE_HIERARCHY" //
)
