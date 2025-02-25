package client

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/sp-api/lwa"
)

type BaseClient struct {
	auth *lwa.Client

	tokenTemp lwa.AccessToken

	TokenListener      func(token lwa.AccessToken)
	TokenErrorListener func(clientId string, err error)
}

func (c *BaseClient) SetConn(auth *lwa.Client) {
	c.auth = auth
}

func (c *BaseClient) GetConn() *lwa.Client {
	return c.auth
}

func (c *BaseClient) SetToken(accessToken lwa.AccessToken) {
	c.tokenTemp = accessToken
}

func (c *BaseClient) GetCurrentToken() lwa.AccessToken {
	return c.tokenTemp
}

func (c *BaseClient) RefreshAccessToken() {
	if c.tokenValid() {
		return
	}
	log.Info("refresh access token...")
	accessToken, err := c.auth.GetAccessToken()
	if err != nil {
		log.Info(err)
		if c.TokenErrorListener != nil {
			c.TokenErrorListener(c.auth.GetClientId(), err)
		}
		return
	}
	if c.TokenListener != nil {
		c.TokenListener(accessToken)
	}

	c.tokenTemp = accessToken
	log.Info(accessToken)
}
func (c *BaseClient) tokenValid() bool {

	return c.tokenTemp.Valid()
}
