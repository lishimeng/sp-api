package seller

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/common"
	"github.com/lishimeng/sp-api/lwa"
	"github.com/lishimeng/sp-api/rest"
	"time"
)

//var AutoRefreshToken = false

type OptFunc func(*Client)

// WithLwa 注入lwa client
var WithLwa = func(conn *lwa.Client, accessTokenListener ...func(token lwa.AccessToken)) OptFunc {
	return func(c *Client) {
		c.auth = conn
		if len(accessTokenListener) > 0 {
			c.tokenListener = accessTokenListener[0]
		}
	}
}

var WithTokenErrorListener = func(listener func(clientId string, err error)) OptFunc {
	return func(c *Client) {
		c.tokenErrorListener = listener
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

		c.centralUrl = m
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
	centralUrl    common.CentralUrl
	ssl           bool
	marketPlaceId string
	userAgent     string

	auth *lwa.Client

	tokenTemp          lwa.AccessToken
	tokenListener      func(token lwa.AccessToken)
	tokenErrorListener func(clientId string, err error)
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
	accessToken, err := c.auth.GetAccessToken()
	if err != nil {
		log.Info(err)
		if c.tokenErrorListener != nil {
			c.tokenErrorListener(c.auth.GetClientId(), err)
		}
		return
	}
	if c.tokenListener != nil {
		c.tokenListener(accessToken)
	}

	c.tokenTemp = accessToken
	log.Info(accessToken)
}

func (c *Client) request(sellerApi bool) *rest.Request {
	if sellerApi {
		return c.sellerRequest()
	} else {
		return c.adsRequest()
	}
}

// request 创建请求, 如果token过期会自动刷新, 添加了通用header
func (c *Client) adsRequest() *rest.Request {
	c.refreshAccessToken()
	var host = string(c.centralUrl.AdHost)
	log.Info("rest request: %s[ssl:%t]", host, c.ssl)
	req := rest.NewRequest(host, c.ssl).
		Authorization(c.tokenTemp.AccessToken).
		Header(string(rest.HeaderUserAgent), c.userAgent).
		Header("Host", host).
		RequestTime(time.Now())
	return req
}

func (c *Client) sellerRequest() *rest.Request {
	c.refreshAccessToken()
	var host = string(c.centralUrl.Endpoint)
	log.Info("rest request: %s[ssl:%t]", host, c.ssl)
	req := rest.NewRequest(host, c.ssl).
		Authorization(c.tokenTemp.AccessToken).
		Header(string(rest.HeaderUserAgent), c.userAgent).
		Header("Host", host).
		RequestTime(time.Now())
	return req
}

func (c *Client) tokenValid() bool {

	return c.tokenTemp.Valid()
}
