package rest

import "time"

type ContentType string

const (
	ApplicationJson ContentType = "application/json"
	FormUrlencoded  ContentType = "application/x-www-form-urlencoded"
)

func (r *Request) Accept(h HeaderValue) *Request {
	r.header.Set("Accept", string(h))
	r.accept = h
	return r
}

func (r *Request) ContentType(h ContentType) *Request {
	r.header.Set("Content-Type", string(h))
	r.contentType = h
	return r
}

func (r *Request) Authorization(h string) *Request {
	r.header.Set("x-amz-access-token", h)
	return r
}

func (r *Request) RequestTime(t time.Time) *Request {
	r.header.Set("x-amz-date", t.UTC().Format("20060102T150405Z"))
	return r
}

func (r *Request) Header(key, value string) *Request {
	r.header.Set(key, value)
	return r
}
