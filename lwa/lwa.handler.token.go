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
		data.Add("scope", ScopeNotification.String())
	} else if opt.Gt == AuthorizationCode {
		data.Add("code", opt.AuthorizeCode)
		data.Add("redirect_uri", opt.RedirectUri)
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
