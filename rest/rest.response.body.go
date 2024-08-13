package rest

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (r *Request) Expect(httpCode int) *Request {
	r.expectedHttpCode = httpCode
	return r
}

func (r *Request) Response(ptr any) *Request {
	r.respPtr = ptr
	return r
}

func (r *Request) onResponse(resp *http.Response) (err error) {
	if resp.StatusCode != r.expectedHttpCode {
		err = errors.New("unexpected http code: " + resp.Status)
		// TODO 如果有通用的error response结构体，可以在这里解析
		return
	}
	if r.accept == "application/json" {
		err = r.jsonResp(resp)
	}
	return
}

func (r *Request) jsonResp(resp *http.Response) (err error) {
	err = json.NewDecoder(resp.Body).Decode(r.respPtr)
	return
}

func (r *Request) textResp() (err error) {
	return
}

func (r *Request) xmlResp() (err error) {
	return
}

func (r *Request) htmlResp() (err error) {
	return
}

func (r *Request) rawResp() (err error) {
	return
}
