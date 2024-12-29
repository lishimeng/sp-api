package seller

import (
	"encoding/json"
	"testing"
)

func TestOrderList(t *testing.T) {

	apiType = SellerType
	c := createApi(t)
	payload, err := c.GetOrders()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(payload)
	bs, _ := json.Marshal(payload)
	t.Logf("%s", string(bs))
}
