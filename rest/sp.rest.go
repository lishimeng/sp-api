package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lishimeng/go-log"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	HeaderUserAgent = "User-Agent"
)

func UserAgent(appid string, version string) string {
	return fmt.Sprintf("%s/%s (Language=Go/12.2; Platform=Ubuntu/22.04)", appid, version)
}

type SpClient struct {
	host string
	path string
	req  *http.Request

	accessToken string

	gzip bool // request gzip
}

var appid = ""
var version = ""

func Init(_appid, _version string) {
	appid = _appid
	version = _version
}

func New(host string) *SpClient {

	c := &SpClient{host: host}
	return c
}

func (c *SpClient) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *SpClient) Path(p string) {
	c.path = p
}

func (c *SpClient) Get(query map[string]string, resultPtr any) (err error) {

	fp, err := url.JoinPath(c.host, c.path)
	if err != nil {
		return
	}
	p := url.Values{}
	for k, v := range query {
		p.Add(k, v)
	}
	fp += "?" + p.Encode()
	log.Info("url:%s", fp)
	req, err := http.NewRequest("GET", fp, nil)
	if err != nil {
		return
	}
	req.Header.Set(HeaderUserAgent, UserAgent(appid, version))
	req.Header.Set("host", c.host)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-amz-date", time.Now().UTC().Format("20060102T150405Z"))
	req.Header.Set("x-amz-access-token", c.accessToken)

	for name, values := range req.Header {
		log.Info("%s:%s", name, values[0])
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("http status code: %d", resp.StatusCode))
		log.Info(err)
		bs, _ := io.ReadAll(resp.Body)

		log.Info(string(bs))
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, resultPtr)
	return
}

func (c *SpClient) Post(data []byte, resultPtr any) (err error) {

	fp, err := url.JoinPath(c.host, c.path)
	if err != nil {
		return
	}
	log.Info("post url:%s", fp)
	req, err := http.NewRequest("POST", fp, bytes.NewReader(data))
	if err != nil {
		return
	}
	c.fillCommonHeaders(req)
	c.setAuth(req)

	log.Info("post headers:")
	for name, values := range req.Header {
		log.Info("%s:%s", name, values[0])
	}
	log.Info("post data:%s", string(data))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Info(err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	log.Info("resp status code:%d", resp.StatusCode)

	log.Info("resp headers:")
	for name, values := range resp.Header {
		log.Info("%s:%s", name, values[0])
	}
	if resp.StatusCode != http.StatusAccepted {
		err = errors.New(fmt.Sprintf("http status code: %d", resp.StatusCode))
		log.Info(err)
		bs, _ := io.ReadAll(resp.Body)

		log.Info(string(bs))
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, resultPtr)

	return
}

func (c *SpClient) auth(req *http.Request) {
	// TODO login
}

func (c *SpClient) setAuth(req *http.Request) {
	req.Header.Set("x-amz-access-token", c.accessToken)
}

func (c *SpClient) fillCommonHeaders(req *http.Request) {
	// user agent
	req.Header.Set(HeaderUserAgent, UserAgent(appid, version))
	req.Header.Set("host", c.host)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-amz-date", time.Now().UTC().Format("20060102T150405Z"))
	req.Header.Set("Content-Type", "application/json")
}
