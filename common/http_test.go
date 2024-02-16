package common

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_httpRequestwithTimeout(t *testing.T) {
	code, body, err := SendReqGet("https://apib.urbox.vn/v1/api/user/me", map[string]string{"content-type": "application/json"}, HttpOption{
		Timeout: time.Millisecond,
	})
	log.Print(code, string(body), ":", err)
}
func Test_httpSendRetry(t *testing.T) {
	code, body, err := SendReqPostWithRetry("http://localhost:3000", map[string]string{"content-type": "application/json"}, []byte(""))
	log.Print(code, string(body), err)
}

func Test_httpSendRetry2(t *testing.T) {
	code, body, err := SendReqPostWithRetry("https://reqbin.com/echo/post/json", map[string]string{"content-type": "application/json", "Authorization": "Bearer mt0dgHmLJMVQhvjpNXDyA83vA_PxH23Y"}, []byte(`{
		"Id": 12345,
		"Customer": "John Smith",
		"Quantity": 1,
		"Price": 10.00
	  }`))
	log.Print(code, string(body), err)
}

func TestSendMultipleRequestPost(t *testing.T) {
	buf := make(chan map[string]interface{}, 2000)
	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		go func() {
			for {
				data := <-buf
				bin, _ := json.Marshal(data)
				hcode, body, err := SendReqPost("https://jsonplaceholder.typicode.com/posts", map[string]string{"content-type": "application/json"}, bin)
				if hcode > 202 {
					log.Print(hcode, string(body), err)
				}
				log.Print("done")
				wg.Done()
			}
		}()
	}
	for i := 0; i < 5000; i++ {
		buf <- map[string]interface{}{
			"userId": 1,
			"id":     0,
			"title":  "eum et est occaecati",
			"body":   "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque velit",
		}
		wg.Add(1)
	}
	wg.Wait()
}

func TestSendGetBenchmark(t *testing.T) {
	for i := 1; i < 100; i++ {
		code, body, err := SendReqGet(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i), map[string]string{"content-type": "application/json"})

		log.Print(code, string(body), err)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)
}
