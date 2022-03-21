package common

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type Option struct {
	Timeout time.Duration
}

// SendReqPost send http post
func SendReqPost(url string, headers map[string]string, body []byte, opts ...Option) (int, []byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

// SendReqPut send http put`
func SendReqPut(url string, headers map[string]string, body []byte, opts ...Option) (int, []byte, error) {
	req, err := http.NewRequest("PUT`", url, bytes.NewBuffer(body))
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
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

// SendReqGet send http get
func SendReqGet(url string, headers map[string]string, opts ...Option) (int, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
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
func SendReqPostWithRetry(url string, headers map[string]string, body []byte) (int, []byte, error) {
	c := &http.Client{
		Transport: &Retry{http.DefaultTransport},
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
		return 0, nil, err
	}
	defer func() {
		req.Close = true
		resp.Body.Close()
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}
