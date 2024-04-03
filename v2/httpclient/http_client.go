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
	timeout   time.Duration
	transport http.RoundTripper
}
type Option func(*Req)

var (
	ErrTimeout = errors.New("request_timeout")
	// 30 số fibonacci đầu tiên
	fibonaccies = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269}
)

func isTimeout(err error) bool {
	return strings.Contains(err.Error(), "context deadline exceeded")
}

type Response struct {
	HttpCode int
	Body     []byte
	Header   http.Header
}

func (r Response) Bytes() []byte {
	data, err := json.MarshalIndent(r, "", "")
	if err != nil {
		return nil
	}
	return data
}

// Exec send request http
func Exec(url string, opts ...Option) (*Response, error) {
	newReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// return -1, nil, err
		return &Response{}, err
	}

	req := &Req{timeout: 10 * time.Second, Request: newReq}

	for _, opt := range opts {
		opt(req)
	}
	if req.transport == nil {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 100
		t.MaxConnsPerHost = 100
		t.MaxIdleConnsPerHost = 100
		req.transport = t
	}
	client := &http.Client{Timeout: req.timeout, Transport: req.transport}
	resp, err := client.Do(req.Request)
	if err != nil {
		if isTimeout(err) {
			return &Response{}, ErrTimeout
		}
		return &Response{}, err
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
		return &Response{
			HttpCode: resp.StatusCode,
			Body:     data,
			Header:   resp.Header,
		}, err
	}
	return &Response{
		HttpCode: resp.StatusCode,
		Body:     data,
		Header:   resp.Header,
	}, nil
}

// Deprecated: New send http request
// Now use httpclient.Exec(...)
// Changed: 3 param to 2 param
// httpcode (int), body ([]byte), ==> *Response{httpcode, body, header}
func New(url string, opts ...Option) (int, []byte, error) {
	newReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1, nil, err
	}

	req := &Req{timeout: 10 * time.Second, Request: newReq}

	for _, opt := range opts {
		opt(req)
	}
	if req.transport == nil {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 100
		t.MaxConnsPerHost = 100
		t.MaxIdleConnsPerHost = 100
		req.transport = t
	}
	client := &http.Client{Timeout: req.timeout, Transport: req.transport}
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

func WithTransport(MaxIdleConns, MaxConnsPerHost, MaxIdleConnsPerHost int) Option {
	return func(r *Req) {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = MaxIdleConns
		t.MaxConnsPerHost = MaxConnsPerHost
		t.MaxIdleConnsPerHost = MaxIdleConnsPerHost
		r.transport = t
	}
}

type retryableTransport struct {
	transport  http.RoundTripper
	retryCount int
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusBadGateway ||
		resp.StatusCode == http.StatusServiceUnavailable ||
		resp.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}
func drainBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
}
func backoff(retries int) time.Duration {
	return time.Duration(fibonaccies[retries]) * time.Second
}

// const RetryCount = 3

func (t *retryableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request body
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	// Send the request
	resp, err := t.transport.RoundTrip(req)
	// Retry logic
	retries := 0
	for shouldRetry(err, resp) && retries < t.retryCount {
		log.Print("RUN Retry")
		// Wait for the specified backoff period
		time.Sleep(backoff(retries))
		// We're going to retry, consume any response to reuse the connection.
		drainBody(resp)
		// Clone the request body again
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		// Retry the request
		resp, err = t.transport.RoundTrip(req)
		retries++
	}
	// Return the response
	return resp, err
}

// WithRetries max retryCount = 30
func WithRetries(retryCount int) Option {
	return func(r *Req) {
		if retryCount > 30 {
			retryCount = 30
		}
		transport := &retryableTransport{
			transport:  &http.Transport{},
			retryCount: retryCount,
		}
		r.transport = transport
	}

}
