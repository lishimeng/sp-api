package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lishimeng/go-log"
	"io"
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
	if r.expectedHttpCode == 0 {
		r.expectedHttpCode = http.StatusOK // 默认200
	}
	err = r.handleRespBody(resp)
	if resp.StatusCode != r.expectedHttpCode {
		err = errors.New("unexpected http code: " + resp.Status)
		return
	}

	return
}

func (r *Request) handleRespBody(resp *http.Response) (err error) {
	if r.accept == "application/json" {
		err = r.jsonResp(resp)
	} else if r.accept == "application/vnd.createasyncreportrequest.v3+json" {
		err = r.jsonAmazonResp(resp)
	}
	return
}

func (r *Request) jsonAmazonResp(resp *http.Response) (err error) {
	var buf bytes.Buffer
	n, err := io.Copy(&buf, resp.Body)
	if err != nil {
		return
	}
	log.Info("resp body: %d bytes", n)
	log.Info(string(buf.Bytes()))
	err = json.NewDecoder(&buf).Decode(r.respPtr)
	return
}

func (r *Request) jsonResp(resp *http.Response) (err error) {
	var buf bytes.Buffer
	n, err := io.Copy(&buf, resp.Body)
	if err != nil {
		return
	}
	log.Info("resp body: %d bytes", n)
	log.Info(string(buf.Bytes()))
	err = json.NewDecoder(&buf).Decode(r.respPtr)
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
