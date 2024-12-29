package common

type Marketplace string

type CentralUrl struct {
	Endpoint ApiEndpoint
	AdHost   AdvertisingHost
	Id       string
}

type ApiEndpoint string

const (
	NorthAmerica ApiEndpoint = "sellingpartnerapi-na.amazon.com"
	Europe       ApiEndpoint = "sellingpartnerapi-eu.amazon.com"
	FarEast      ApiEndpoint = "sellingpartnerapi-fe.amazon.com"
)

type AdvertisingHost string

const (
	AdsNorthAmerica AdvertisingHost = "advertising-api.amazon.com"
	AdsEurope       AdvertisingHost = "advertising-api-eu.amazon.com"
	AdsFarEast      AdvertisingHost = "advertising-api-fe.amazon.com"
)

const ( // North America
	Canada Marketplace = "Canada"
	Us     Marketplace = "US"
	Mexico Marketplace = "Mexico"
	Brazil Marketplace = "Brazil"
)

const ( // Europe
	Spain       Marketplace = "Spain"
	UK          Marketplace = "UK"
	France      Marketplace = "France"
	Netherlands Marketplace = "Netherlands" // 荷兰 Holland
	Germany     Marketplace = "Germany"
	Italy       Marketplace = "Italy"
	Sweden      Marketplace = "Sweden"
	SouthAfrica Marketplace = "SouthAfrica"
	Poland      Marketplace = "Poland"
	Egypt       Marketplace = "Egypt" // 埃及
	Turkey      Marketplace = "Turkey"
	SaudiArabia Marketplace = "SaudiArabia" // 沙特
	UAE         Marketplace = "UAE"         // 阿联酋
	India       Marketplace = "India"
	Belgium     Marketplace = "Belgium" // 比利时
)

const ( // Far East
	Singapore Marketplace = "Singapore"
	Australia Marketplace = "Australia"
	Japan     Marketplace = "Japan"
)

type MarketplaceRepo map[Marketplace]CentralUrl

func (r MarketplaceRepo) add(marketplace Marketplace, id string, endpoint ApiEndpoint, adHost AdvertisingHost) {
	r[marketplace] = CentralUrl{
		Id:       id,
		Endpoint: endpoint,
		AdHost:   adHost,
	}
}

var marketplaces = make(MarketplaceRepo)

func init() {

	// North America
	marketplaces.add(Canada, "A2EUQ1WTGCTBG2", NorthAmerica, AdsNorthAmerica)
	marketplaces.add(Us, "ATVPDKIKX0DER", NorthAmerica, AdsNorthAmerica)
	marketplaces.add(Mexico, "A1AM78C64UM0Y8", NorthAmerica, AdsNorthAmerica)
	marketplaces.add(Brazil, "A2Q3Y263D00KWC", NorthAmerica, AdsNorthAmerica)

	// Europe
	marketplaces.add(Spain, "A1RKKUPIHCS9HS", Europe, AdsEurope)
	marketplaces.add(UK, "A1F83G8C2ARO7P", Europe, AdsEurope)
	marketplaces.add(France, "A13V1IB3VIYZZH", Europe, AdsEurope)
	marketplaces.add(Netherlands, "A1805IZSGTT6HS", Europe, AdsEurope)
	marketplaces.add(Germany, "A1PA6795UKMFR9", Europe, AdsEurope)
	marketplaces.add(Italy, "APJ6JRA9NG5V4", Europe, AdsEurope)
	marketplaces.add(Sweden, "A2NODRKZP88ZB9", Europe, AdsEurope)
	marketplaces.add(SouthAfrica, "AE08WJ6YKNBMC", Europe, AdsEurope)
	marketplaces.add(Poland, "A1C3SOZRARQ6R3", Europe, AdsEurope)
	marketplaces.add(Egypt, "ARBP9OOSHTCHU", Europe, AdsEurope)
	marketplaces.add(Turkey, "A33AVAJ2PDY3EV", Europe, AdsEurope)
	marketplaces.add(SaudiArabia, "A17E79C6D8DWNP", Europe, AdsEurope)
	marketplaces.add(UAE, "A2VIGQ35RCS4UG", Europe, AdsEurope)
	marketplaces.add(India, "A21TJRUUN4KGV", Europe, AdsEurope)
	marketplaces.add(Belgium, "AMEN7PMS3EDWL", Europe, AdsEurope)

	// Far East
	marketplaces.add(Singapore, "A19VAU5U5O7RUS", FarEast, AdsFarEast)
	marketplaces.add(Australia, "A39IBJ37TRP1C6", FarEast, AdsFarEast)
	marketplaces.add(Japan, "A1VC38T7YXB528", FarEast, AdsFarEast)

}

func GetCentralURL(marketplace Marketplace) (ok bool, c CentralUrl) {
	c, ok = marketplaces[marketplace]
	return
}
