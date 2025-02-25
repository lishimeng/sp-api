package advertising

import "github.com/lishimeng/sp-api/rest"

type ProfileAccountInfo struct {
	MarketplaceId      string `json:"marketplaceStringId"`
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	ValidPaymentMethod string `json:"validPaymentMethod"`
}

type Profile struct {
	ProfileId    string             `json:"profileId"`
	CountryCode  string             `json:"countryCode"`
	CurrencyCode string             `json:"currencyCode"`
	Timezone     string             `json:"timezone"`
	AccountInfo  ProfileAccountInfo `json:"accountInfo"`
}

func (c *Client) GetProfile() (err error) {
	var action = "/v2/profiles"
	var resp Profile
	c.profileId = "" // 清空profile
	err = c.request().Path(action).Response(&resp).Json(rest.POST)
	if err == nil {
		if c.onProfile != nil {
			c.onProfile(resp)
		}
		c.profileId = resp.ProfileId
	}
	return
}
