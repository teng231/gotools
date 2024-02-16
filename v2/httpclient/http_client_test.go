package httpclient

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSendReq(t *testing.T) {
	code, body, err := New("https://webhook.site/7dac37f5-a041-437b-a14f-0b0efbcf9515",
		WithMethod("POST"),
		WithTimeout(time.Second),
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
