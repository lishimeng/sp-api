package seller

import (
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/rest"
	"net/http"
)

type ItemSearchOption struct {
	identifiersEnable bool
	identifiersType   common.IdentifiersType
	identifiers       string

	sellerIdEnable bool
	sellerId       string

	keywordsEnable bool
	keywords       string
}
type ItemSearchOptionFunc func(iso *ItemSearchOption)

var WithIdentify = func(category common.IdentifiersType, value ...string) ItemSearchOptionFunc {
	return func(iso *ItemSearchOption) {
		if len(value) == 0 {
			return
		}
		iso.identifiersEnable = true
		iso.identifiersType = category
		for i, v := range value {
			if i > 0 {
				iso.identifiers += "," + v
			} else {
				iso.identifiers = v
			}
		}
	}
}

var WithKeywords = func(value ...string) ItemSearchOptionFunc {
	return func(iso *ItemSearchOption) {
		if len(value) == 0 {
			return
		}
		iso.keywordsEnable = true
		for i, v := range value {
			if i > 0 {
				iso.keywords += "," + v
			} else {
				iso.keywords = v
			}
		}
	}
}

func (c *Client) GetItems(searchOpt ...ItemSearchOptionFunc) (result ItemSearchResults, err error) {
	var action = "/catalog/2022-04-01/items"

	opt := &ItemSearchOption{}
	for _, f := range searchOpt {
		f(opt)
	}

	// 检查内容
	if opt.identifiersType == common.SKU {
		if !opt.sellerIdEnable {
			panic("sellerId is required") // TODO
		}
	}

	err = c.sellerRequest().Path(action).
		Accept("application/json").
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

func (c *Client) GetItem(asin string) (err error) {
	var action = "/catalog/2022-04-01/items"

	err = c.sellerRequest().Path(action).Path(asin).
		Accept("application/json").
		ContentType(rest.ApplicationJson).
		Expect(http.StatusOK).
		Query("marketplaceIds", c.marketPlaceId).
		Query("identifiers", "dog").
		Query("identifiersType", "ASIN").
		//Response(&result).
		Get()
	return
}
