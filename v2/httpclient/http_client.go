package httpclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Req struct {
	*http.Request
	timeout time.Duration
}
type Option func(*Req)

var (
	ErrTimeout = errors.New("request_timeout")
)

func isTimeout(err error) bool {
	return strings.Contains(err.Error(), "context deadline exceeded")
}

// Send send http post
func New(url string, opts ...Option) (int, []byte, error) {
	newReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1, nil, err
	}

	req := &Req{timeout: 10 * time.Second, Request: newReq}

	for _, opt := range opts {
		opt(req)
	}
	client := &http.Client{Timeout: req.timeout}
	resp, err := client.Do(req.Request)
	if err != nil {
		if isTimeout(err) {
			return 0, nil, ErrTimeout
		}
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		log.Print(err)
		return resp.StatusCode, data, err
	}
	return resp.StatusCode, data, nil
}

func WithBody(body any) Option {
	return func(r *Req) {
		out, _ := json.Marshal(body)
		var rd io.Reader = bytes.NewBuffer(out)
		rc, ok := rd.(io.ReadCloser)
		if !ok && body != nil {
			rc = io.NopCloser(rd)
		}
		r.Request.Body = rc
	}
}

func WithMethod(method string) Option {
	return func(r *Req) {
		r.Request.Method = strings.ToUpper(method)
	}
}

func WithHeader(headers map[string]string) Option {
	return func(r *Req) {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
	}
}

func WithTimeout(dur time.Duration) Option {
	return func(r *Req) {
		r.timeout = dur
	}
}

func WithBasicAuth(username, password string) Option {
	return func(r *Req) {
		r.SetBasicAuth(username, password)
	}
}
