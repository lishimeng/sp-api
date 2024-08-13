package seller

import (
	"github.com/lishimeng/sp-api/rest"
	"net/http"
	"time"
)

func (c *Client) GetItems() (result ItemSearchResults, err error) {
	var action = "/catalog/2022-04-01/items"

	c.refreshAccessToken() // get一次就可以了, 刷新token
	err = rest.NewRequest(c.endPoint, c.ssl).
		Path(action).
		Authorization(c.tokenTemp.AccessToken).
		Header(rest.HeaderUserAgent, c.userAgent).
		Header("Host", c.endPoint).
		Accept("application/json").
		RequestTime(time.Now()).
		Expect(http.StatusOK).
		Query("marketplaceIds", c.marketPlaceId).
		Query("identifiers", "dog").
		Query("identifiersType", "ASIN").
		Response(&result).Get()

	if err != nil {
		return
	}
	return
}
