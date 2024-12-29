package lwa

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	apiHost = "https://api.amazon.com"
)

const (
	AuthorizeHost = "https://www.amazon.com"
)

type Client struct {
	host                string // api host
	clientId            string // 客户端ID
	clientSecret        string // 密钥
	defaultRefreshToken string // 默认RefreshToken
	scopes              []Scope
}

type GrantType string

const (
	AuthorizationCode GrantType = "authorization_code"
	ClientCredentials GrantType = "client_credentials"
	RefreshToken      GrantType = "refresh_token"
)

type Scope string

func (s Scope) String() string {
	return string(s)
}

const (
	ScopeNotification         Scope = "sellingpartnerapi::notifications"
	ScopeAppManagement        Scope = "sellingpartnerapi::client_credential:rotation"
	ScopeMigrationApi         Scope = "sellingpartnerapi::migration"
	ScopeAdvertisingCampaign  Scope = "advertising::campaign_management"
	ScopeAdvertisingAudiences Scope = "advertising::audiences"
)

type TokenOption struct {
	Gt            GrantType
	Scope         string
	RefreshToken  string
	AuthorizeCode string
	RedirectUri   string
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
var WithScope = func(scope ...Scope) TokenOptionFunc {
	return func(c *TokenOption) {
		if len(scope) > 0 {
			c.Scope = scope[0].String()
		}
	}
}
var WithAuthorizeCode = func(code, redirectUri string) TokenOptionFunc {
	return func(c *TokenOption) {
		c.RedirectUri = redirectUri
		c.AuthorizeCode = code
	}
}
var WithRefreshToken = func(rt string) TokenOptionFunc {
	return func(c *TokenOption) {
		c.RefreshToken = rt
	}
}

func (c *Client) Code(redirectUri string, scopes ...Scope) (authorizeUrl string, err error) {
	p, err := url.JoinPath(AuthorizeHost, "/ap/oa")
	if err != nil {
		return
	}
	var scope Scope

	authorizeUrl = fmt.Sprintf("client_id=%s&response_type=code&redirect_uri=%s",
		c.clientId, redirectUri)
	if len(scopes) > 0 {
		scope = scopes[0]
		authorizeUrl = fmt.Sprintf("%s&scope=%s", authorizeUrl, scope.String())
	}
	authorizeUrl = url.PathEscape(authorizeUrl)
	authorizeUrl = fmt.Sprintf("%s?%s", p, authorizeUrl)
	return
}

func (c *Client) GetAccessTokenWithCode(authorizeCode string, redirectUri string) (at AccessToken, err error) {

	at, err = c.Token(WithGrantType(AuthorizationCode),
		WithAuthorizeCode(authorizeCode, redirectUri))
	if err != nil {
		return
	}
	at.RefreshTime() // 刷新时间戳,valid函数生效
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

func (c *Client) GetClientId() string {
	return c.clientId
}

func httpGet(url string, params map[string]string, result *AccessToken) {

}
