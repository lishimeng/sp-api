package seller

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/lwa"
)

//var AutoRefreshToken = false

type OptFunc func(*Client)

// WithLwa 注入lwa client
var WithLwa = func(conn *lwa.Client, accessTokenListener ...func(token lwa.AccessToken)) OptFunc {
	return func(c *Client) {
		c.connector = conn
		if len(accessTokenListener) > 0 {
			c.tokenListener = accessTokenListener[0]
		}
	}
}

// WithLwaToken 注入缓存token
var WithLwaToken = func(token lwa.AccessToken) OptFunc {
	return func(c *Client) {
		c.tokenTemp = token
	}
}

// WithMarketplace 设置marketplace
var WithMarketplace = func(marketPlace common.Marketplace, ssl bool) OptFunc {
	return func(c *Client) {
		ok, m := common.GetCentralURL(marketPlace)
		if !ok {
			panic(fmt.Sprintf("marketplace not support: %s", marketPlace))
		}

		c.endPoint = string(m.Endpoint)
		c.ssl = ssl
		c.marketPlaceId = m.Id
	}
}

var WithUserAgent = func(ua string) OptFunc {
	return func(c *Client) {
		c.userAgent = ua
	}
}

func UserAgent(appid string, version string) string {
	return fmt.Sprintf("%s/%s (Language=Go/12.2; Platform=Ubuntu/22.04)", appid, version)
}

type Client struct {
	endPoint      string
	ssl           bool
	marketPlaceId string
	userAgent     string

	connector *lwa.Client

	//logic *rest.SpClient

	tokenTemp     lwa.AccessToken
	tokenListener func(token lwa.AccessToken)
}

func New(opts ...OptFunc) *Client {
	c := &Client{}

	c._init(opts...)
	return c
}

func (c *Client) _init(opts ...OptFunc) {

	for _, opt := range opts {
		opt(c)
	}
}

// GetCurrentToken 查看当前token
func (c *Client) GetCurrentToken() lwa.AccessToken {
	return c.tokenTemp
}

func (c *Client) refreshAccessToken() {
	if c.tokenValid() {
		return
	}
	log.Info("refresh access token...")
	accessToken, err := c.connector.GetAccessToken()
	if err != nil {
		log.Info(err)
		return
	}
	if c.tokenListener != nil {
		c.tokenListener(accessToken)
	}

	c.tokenTemp = accessToken
	log.Info(accessToken)
}

func (c *Client) tokenValid() bool {

	return c.tokenTemp.Valid()
}
