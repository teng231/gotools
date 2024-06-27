package httpclient

import (
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	urlWebhook = "https://webhook.site/2c65c12d-198f-41d2-bc83-a03f13bf0a40"
)

func TestSendReq(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("POST"),
		WithTimeout(2*time.Second),
		WithBasicAuth("te", "manh"),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		WithBody(map[string]any{
			"userId": 1,
			"id":     1,
			"title":  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		}),
	)

	log.Print(code, string(body), err)
}

func TestSendGetBenchmark(t *testing.T) {
	for i := 1; i < 100; i++ {
		code, body, err := New(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i),
			// WithMethod("POST"),
			WithTimeout(2*time.Second),
			// WithBasicAuth("te", "manh"),
			WithHeader(map[string]string{
				"content-type": "application/json",
			}),
		)

		log.Print(code, string(body), err)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)
}

func TestSendReqWithPerformace(t *testing.T) {
	code, body, err := New("https://jsonplaceholder.typicode.com/todos/1",
		WithMethod("GET"),
		WithTimeout(20*time.Second),
		WithBasicAuth("te", "manh"),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		WithTransport(100, 100, 100),
	)

	log.Print(code, string(body), err)
}

func TestSendReqWithRetry(t *testing.T) {
	code, body, err := New("https://dummyjson.com/http/503",
		WithMethod("GET"),
		WithTimeout(20*time.Second),
		WithBasicAuth("te", "manh"),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		WithRetries(3),
	)

	log.Print(code, string(body), err)
}

func TestSendReqWithTimeout(t *testing.T) {
	code, body, err := New("https://dummyjson.com/products/1?delay=5000",
		WithMethod("GET"),
		WithTimeout(4*time.Second),
		// WithBasicAuth("te", "manh"),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		// WithRetries(3),
	)

	log.Print(code, string(body), err)
}

func TestPutRequest(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("PUT"),
		WithTimeout(100*time.Second),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		WithPutFile("../../README.MD"),
	)

	log.Print(code, string(body), err)
}

func TestWithURLEncode(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("POST"),
		WithTimeout(10*time.Second),
		WithHeader(map[string]string{
			"content-type": "application/json",
		}),
		WithUrlEncode(map[string]string{
			"mot": "1", "hai": "2",
		}),
	)

	log.Print(code, string(body), err)
}
func TestPutRequest2(t *testing.T) {
	resp, _ := Exec("https://image.lexica.art/full_webp/01707d2a-8b3f-4ac9-b679-b6ff97e7c360")
	resp2, err := Exec("https://leia-storage-service-production.s3.us-east-1.amazonaws.com/timed/D001/95802150-215f-40f3-bc71-b00a4a72da61/e5e3d5a3-a35d-4c2f-9f9e-5324f5f402ec/c138ad49-6a50-4831-b5d6-b2a8f7d6a3d8/e8eefce0860d97b021739d3b7209ea22.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIASC7ECGJVHARKLZ6E%2F20240419%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240419T083548Z&X-Amz-Expires=86400&X-Amz-Signature=b744ab11d8665872cef139d5114038ad832929236654e969bea08e9da0d5810a&X-Amz-SignedHeaders=host&x-id=PutObject",
		WithMethod("PUT"),
		WithTimeout(15*time.Second),
		WithHeader(map[string]string{
			"content-type": "application/octet-stream",
		}),
		WithPutBytes(resp.Body),
	)

	log.Print(resp2.HttpCode, string(resp2.Body), err)
}

func TestWithFormData(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("POST"),
		WithTimeout(10*time.Second),
		WithFormData(map[string]string{
			"mot": "1", "hai": "2",
		}),
	)

	log.Print(code, string(body), err)
}
func TestWithWithNilBody(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("POST"),
		WithTimeout(10*time.Second),
		WithBody(nil),
	)

	log.Print(code, string(body), err)
}

func TestWithWithNilPutContent(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("PUT"),
		WithTimeout(10*time.Second),
		WithPutBytes(nil),
	)

	log.Print(code, string(body), err)
}
func TestWithWithNilEncoding(t *testing.T) {
	code, body, err := New(urlWebhook,
		WithMethod("POST"),
		WithTimeout(10*time.Second),
		WithUrlEncode(nil),
	)

	log.Print(code, string(body), err)
}
