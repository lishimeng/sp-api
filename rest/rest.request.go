package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HeaderValue string

type Request struct {
	//req              *http.Request
	//client           *SpClient
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
		fullPath = fmt.Sprintf("%s?=%s", fullPath, query)
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

func (r *Request) Post() (err error) {
	err = r.Do("POST")
	return
}
func (r *Request) FormUrlencoded() (err error) {
	r.ContentType(FormUrlencoded)
	err = r.Do("POST")
	return
}

func (r *Request) Json() (err error) {
	r.ContentType(ApplicationJson)
	err = r.Do("POST")
	return
}
