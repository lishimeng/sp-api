package rest

import (
	"encoding/json"
	"io"
	"strings"
)

func (r *Request) bodyReader() (reader io.Reader, err error) {
	if r.requestBody == nil {
		return
	}
	var headers []string
	for key, value := range r.header {
		if strings.ToLower(key) == "content-type" {
			headers = value
		}
	}
	if len(headers) == 0 {
		return
	}
	switch headers[0] {
	case "application/json":
		reader, err = jsonReader(r.requestBody)
	}
	return
}

func jsonReader(b any) (r io.Reader, err error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return
	}
	r = strings.NewReader(string(bs))
	return
}
