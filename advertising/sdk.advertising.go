package advertising

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/client"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/lwa"
	"github.com/lishimeng/sp-api/rest"
)

type Client struct {
	client.BaseClient
	centralUrl    common.CentralUrl
	ssl           bool
	marketPlaceId string
	userAgent     string

	profileId string
	onProfile func(profile Profile)
}

type OptFunc func(*Client)

// WithLwa 注入lwa client
var WithLwa = func(conn *lwa.Client, accessTokenListener ...func(token lwa.AccessToken)) OptFunc {
	return func(c *Client) {
		c.BaseClient.SetConn(conn)
		if len(accessTokenListener) > 0 {
			c.TokenListener = accessTokenListener[0]
		}
	}
}

var WithTokenErrorListener = func(listener func(clientId string, err error)) OptFunc {
	return func(c *Client) {
		c.TokenErrorListener = listener
	}
}

// WithLwaToken 注入缓存token
var WithLwaToken = func(token lwa.AccessToken) OptFunc {
	return func(c *Client) {
		c.SetToken(token)
	}
}

var WithProfile = func(profileId string) OptFunc {
	return func(c *Client) {
		c.profileId = profileId
	}
}

// WithProfileIdCallback 获取profile的时候回调
var WithProfileIdCallback = func(h func(profile Profile)) OptFunc {
	return func(c *Client) {
		c.onProfile = h
	}
}

// WithMarketplace 设置marketplace
var WithMarketplace = func(marketPlace common.Marketplace, ssl bool) OptFunc {
	return func(c *Client) {
		ok, m := common.GetCentralURL(marketPlace)
		if !ok {
			panic(fmt.Sprintf("marketplace not support: %s", marketPlace))
		}

		c.centralUrl = m
		c.ssl = ssl
		c.marketPlaceId = m.Id
	}
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

// request 创建请求, 如果token过期会自动刷新, 添加了通用header
func (c *Client) request() *rest.Request {
	c.RefreshAccessToken()
	var host = string(c.centralUrl.AdHost)
	log.Info("rest request: %s[ssl:%t]", host, c.ssl)
	req := rest.NewRequest(host, c.ssl).
		Header("Authorization", "Bearer "+c.GetCurrentToken().AccessToken).
		Header(string(rest.HeaderUserAgent), c.userAgent).
		Header("Amazon-Advertising-API-ClientId", c.GetConn().GetClientId())
	if len(c.profileId) > 0 {
		req = req.Header("WithProfileIdCallback", c.profileId)
	}
	return req
}
