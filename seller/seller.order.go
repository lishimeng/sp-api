package seller

import (
	"github.com/lishimeng/sp-api/rest"
	"net/http"
)

type OrderResp map[string]any

func (c *Client) GetOrders() (payload OrderResp, err error) {
	var action = "/orders/v0/orders"
	var result OrderResp
	err = c.sellerRequest().Path(action).
		Query("MarketplaceIds", "ATVPDKIKX0DER").
		Query("CreatedAfter", "2024-12-01").
		Query("MaxResultPerPage", "50").
		Accept("application/json").
		Expect(http.StatusOK).Response(&result).Json(rest.GET)

	if err != nil {
		return
	}
	payload = result
	return
}
