package rest

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HeaderValue string

type Request struct {
	expectedHttpCode int
	respPtr          any
	accept           HeaderValue
	contentType      ContentType

	actionPath  []string
	header      http.Header
	endpoint    string
	requestBody any
	ssl         bool
	query       *url.Values
}

func NewRequest(endpoint string, ssl bool) *Request {
	q := &Request{}
	q.header = make(http.Header)
	q.query = &url.Values{}
	q.endpoint = endpoint
	q.ssl = ssl
	return q
}

func (r *Request) Path(p ...string) *Request {
	r.actionPath = append(r.actionPath, p...)
	return r
}

func (r *Request) Body(b any) *Request {
	r.requestBody = b
	return r
}

func (r *Request) Query(key, value string) *Request {
	r.query.Add(key, value)
	return r
}

func (r *Request) Do(method string) (err error) {

	var host string
	if !strings.HasPrefix(r.endpoint, "http") {
		var schema string
		if r.ssl {
			schema = "https"
		} else {
			schema = "http"
		}
		host = fmt.Sprintf("%s://%s", schema, r.endpoint)
	} else {
		host = r.endpoint
	}
	fullPath, err := url.JoinPath(host, r.actionPath...)
	if err != nil {
		return
	}

	// 拼接query
	query := r.query.Encode()
	if len(query) > 0 {
		fullPath = fmt.Sprintf("%s?%s", fullPath, query)
	}

	// 处理request body
	reader, err := r.bodyReader()
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, fullPath, reader)
	if err != nil {
		return
	}

	for key, value := range r.header {
		req.Header[key] = value
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	log.Info("http response code: [%d]%s", resp.StatusCode, resp.Status)
	defer func() {
		_ = resp.Body.Close()
	}()
	err = r.onResponse(resp)
	return
}

func (r *Request) Get() (err error) {
	err = r.Do("GET")
	return
}

// Post 发送Post请求, 默认不携带ContentType, 需要按需设置
func (r *Request) Post() (err error) {
	err = r.Do("POST")
	return
}

// FormUrlencoded 发送Post FormUrlencoded请求, 默认携带ContentType: application/x-www-form-urlencoded
func (r *Request) FormUrlencoded() (err error) {
	r.ContentType(FormUrlencoded)
	err = r.Do("POST")
	return
}

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

// Json 发送Post Json请求, 默认携带ContentType: application/json
func (r *Request) Json(method Method, contentType ...ContentType) (err error) {
	if len(contentType) > 0 {
		r.ContentType(contentType[0])
	} else {
		r.ContentType(ApplicationJson) // default content type
	}
	err = r.Do(string(method))
	return
}

func (r *Request) Download(readFunc func(totalSize int64, reader io.Reader)) (err error) {
	resp, err := http.DefaultClient.Get(r.endpoint)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http response code [%d]", resp.StatusCode)
		return
	}

	if readFunc == nil {
		return
	}

	readFunc(resp.ContentLength, resp.Body)

	return
}
