package common

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	E_timeout = errors.New("request_timeout")
)

type HttpOption struct {
	Timeout time.Duration
}

func isTimeout(err error) bool {
	return strings.Contains(err.Error(), "context deadline exceeded")
}

// SendReqPost send http post
func SendReqPost(url string, headers map[string]string, body []byte, opts ...HttpOption) (int, []byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Close = true
	if err != nil {
		return 0, nil, err
	}
	if len(headers) > 0 {
		for k, val := range headers {
			req.Header.Set(k, val)
		}
	}
	client := &http.Client{Timeout: 10 * time.Second}
	if len(opts) > 0 && opts[0].Timeout != 0 {
		client.Timeout = opts[0].Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		if isTimeout(err) {
			return 0, nil, E_timeout
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

	body, _ = io.ReadAll(reader)
	return resp.StatusCode, body, nil
}

// SendReqPut send http put`
func SendReqPut(url string, headers map[string]string, body []byte, opts ...HttpOption) (int, []byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Close = true
	if err != nil {
		return 0, nil, err
	}
	if len(headers) > 0 {
		for k, val := range headers {
			req.Header.Set(k, val)
		}
	}
	client := &http.Client{Timeout: 10 * time.Second}
	if len(opts) > 0 && opts[0].Timeout != 0 {
		client.Timeout = opts[0].Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		if isTimeout(err) {
			return 0, nil, E_timeout
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
	body, _ = io.ReadAll(reader)
	return resp.StatusCode, body, nil
}

// SendReqGet send http get
func SendReqGet(url string, headers map[string]string, opts ...HttpOption) (int, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	if err != nil {
		return 0, nil, err
	}
	if len(headers) > 0 {
		for k, val := range headers {
			req.Header.Set(k, val)
		}
	}

	client := &http.Client{Timeout: 20 * time.Second}
	if len(opts) > 0 && opts[0].Timeout != 0 {
		client.Timeout = opts[0].Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		if isTimeout(err) {
			return 0, nil, E_timeout
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
	body, _ := io.ReadAll(reader)
	return resp.StatusCode, body, nil
}

// Retry - http client with retry support
type Retry struct {
	http.RoundTripper
}

// RoundTrip Naive Retry - every 2 seconds
func (r *Retry) RoundTrip(req *http.Request) (*http.Response, error) {
	for {
		resp, err := r.RoundTripper.RoundTrip(req)

		// just an example
		// we potentially could retry on 429 for example
		if err == nil && resp.StatusCode < 500 {
			return resp, err
		}

		select {
		// check if canceled or timed-out
		case <-req.Context().Done():
			return resp, req.Context().Err()
		case <-time.After(5 * time.Second):
		}
	}
}

// SendReqPostWithRetry send http post
func SendReqPostWithRetry(url string, headers map[string]string, body []byte, opts ...HttpOption) (int, []byte, error) {
	c := &http.Client{
		Transport: &Retry{http.DefaultTransport},
		Timeout:   10 * time.Second,
	}
	if len(opts) > 0 && opts[0].Timeout != 0 {
		c.Timeout = opts[0].Timeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}
	if len(headers) > 0 {
		for k, val := range headers {
			req.Header.Set(k, val)
		}
	}
	resp, err := c.Do(req)
	if err != nil {
		if isTimeout(err) {
			return 0, nil, E_timeout
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
	body, _ = io.ReadAll(reader)
	return resp.StatusCode, body, nil
}
