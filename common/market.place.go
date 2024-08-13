package common

type Marketplace string

type CentralUrl struct {
	Endpoint ApiEndpoint
	Id       string
}

type ApiEndpoint string

const (
	NorthAmerica ApiEndpoint = "sellingpartnerapi-na.amazon.com"
	Europe       ApiEndpoint = "sellingpartnerapi-eu.amazon.com"
	FarEast      ApiEndpoint = "sellingpartnerapi-fe.amazon.com"
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

func (r MarketplaceRepo) add(marketplace Marketplace, id string, endpoint ApiEndpoint) {
	r[marketplace] = CentralUrl{
		Id:       id,
		Endpoint: endpoint,
	}
}

var marketplaces = make(MarketplaceRepo)

func init() {

	// North America
	marketplaces.add(Canada, "A2EUQ1WTGCTBG2", NorthAmerica)
	marketplaces.add(Us, "ATVPDKIKX0DER", NorthAmerica)
	marketplaces.add(Mexico, "A1AM78C64UM0Y8", NorthAmerica)
	marketplaces.add(Brazil, "A2Q3Y263D00KWC", NorthAmerica)

	// Europe
	marketplaces.add(Spain, "A1RKKUPIHCS9HS", Europe)
	marketplaces.add(UK, "A1F83G8C2ARO7P", Europe)
	marketplaces.add(France, "A13V1IB3VIYZZH", Europe)
	marketplaces.add(Netherlands, "A1805IZSGTT6HS", Europe)
	marketplaces.add(Germany, "A1PA6795UKMFR9", Europe)
	marketplaces.add(Italy, "APJ6JRA9NG5V4", Europe)
	marketplaces.add(Sweden, "A2NODRKZP88ZB9", Europe)
	marketplaces.add(SouthAfrica, "AE08WJ6YKNBMC", Europe)
	marketplaces.add(Poland, "A1C3SOZRARQ6R3", Europe)
	marketplaces.add(Egypt, "ARBP9OOSHTCHU", Europe)
	marketplaces.add(Turkey, "A33AVAJ2PDY3EV", Europe)
	marketplaces.add(SaudiArabia, "A17E79C6D8DWNP", Europe)
	marketplaces.add(UAE, "A2VIGQ35RCS4UG", Europe)
	marketplaces.add(India, "A21TJRUUN4KGV", Europe)
	marketplaces.add(Belgium, "AMEN7PMS3EDWL", Europe)

	// Far East
	marketplaces.add(Singapore, "A19VAU5U5O7RUS", FarEast)
	marketplaces.add(Australia, "A39IBJ37TRP1C6", FarEast)
	marketplaces.add(Japan, "A1VC38T7YXB528", FarEast)

}

func GetCentralURL(marketplace Marketplace) (ok bool, c CentralUrl) {
	c, ok = marketplaces[marketplace]
	return
}
