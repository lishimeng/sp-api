package lwa

import (
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiHost = "https://api.amazon.com"
)

type Client struct {
	host                string // api host
	clientId            string // 客户端ID
	clientSecret        string // 密钥
	defaultRefreshToken string // 默认RefreshToken
}

type GrantType string

const (
	ClientCredentials GrantType = "client_credentials"
	RefreshToken      GrantType = "refresh_token"
)

type Scope string

const (
	ScopeNotification  = "sellingpartnerapi::notifications"
	ScopeAppManagement = "sellingpartnerapi::client_credential:rotation"
	ScopeMigrationApi  = "sellingpartnerapi::migration"
)

type TokenOption struct {
	Gt           GrantType
	Scope        string
	RefreshToken string
}

// New 创建客户端. defaultRefreshToken 可选
func New(clientId, secret string, defaultRefreshToken ...string) *Client {
	c := &Client{host: apiHost, clientSecret: secret, clientId: clientId}
	if len(defaultRefreshToken) > 0 {
		c.defaultRefreshToken = defaultRefreshToken[0] // 默认refreshToken
	}
	return c
}

type TokenOptionFunc func(c *TokenOption)

var WithGrantType = func(gt GrantType) TokenOptionFunc {
	return func(c *TokenOption) {
		c.Gt = gt
	}
}
var WithScope = func(scope string) TokenOptionFunc {
	return func(c *TokenOption) {
		c.Scope = scope
	}
}
var WithRefreshToken = func(rt string) TokenOptionFunc {
	return func(c *TokenOption) {
		c.RefreshToken = rt
	}
}

// Token 申请token
func (c *Client) Token(opts ...TokenOptionFunc) (at AccessToken, err error) {

	var action = "/auth/o2/token"
	fullPath, err := url.JoinPath(c.host, action)
	if err != nil {
		return
	}
	var opt TokenOption
	for _, optFn := range opts {
		optFn(&opt)
	}

	data := url.Values{}
	data.Add("grant_type", string(opt.Gt))
	data.Add("client_id", c.clientId)
	data.Add("client_secret", c.clientSecret)
	if opt.Gt == RefreshToken {
		data.Add("refresh_token", opt.RefreshToken)
	} else if opt.Gt == ClientCredentials {
		data.Add("scope", ScopeNotification)
	}

	log.Info("lwa token:")
	log.Info("url: %s", fullPath)
	log.Info("data: %s", data.Encode())
	req, err := http.NewRequest("POST", fullPath, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Info(err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		log.Info(err)
		bs, _ := io.ReadAll(resp.Body)
		log.Info(string(bs))
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Info(err)
		return
	}

	err = json.Unmarshal(bs, &at)
	if err != nil {
		log.Info(err)
		return
	}

	// 返回token
	return
}

func (c *Client) GetAccessToken(rt ...string) (at AccessToken, err error) {
	var t string
	if len(rt) > 0 {
		t = rt[0]
	} else {
		t = c.defaultRefreshToken
	}
	if len(t) == 0 {
		err = errors.New("refresh_token empty")
		return
	}
	at, err = c.Token(WithGrantType(RefreshToken), WithRefreshToken(t))
	if err != nil {
		return
	}
	at.RefreshTime() // 刷新时间戳,valid函数生效
	return
}

// ClientCredentials 客户端模式
func (c *Client) ClientCredentials() (at AccessToken, err error) {
	at, err = c.Token(WithGrantType(ClientCredentials), WithScope(ScopeNotification))
	return
}

func httpGet(url string, params map[string]string, result *AccessToken) {

}
