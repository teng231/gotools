package common

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func Test_toNano(t *testing.T) {
	tbl := []int64{
		1591931629214,
		1591931629214 * 1e6,
		1591931652,
		26532196,
	}
	for _, item := range tbl {
		tm := time.Unix(0, ToNs(item))
		log.Print(tm.String())
	}
}

func Test_toMili(t *testing.T) {
	tbl := []int64{
		1591931629214,
		1591931629214 * 1e6,
		1591931652,
		26532196,
	}
	for _, item := range tbl {
		tm := time.Unix(0, ToNs(item))
		log.Print(tm.String())
	}
}

// http://pos.giftpop.com.vn:9901/interface/order/goodsListAll.m12

func Test_httpSend(t *testing.T) {
	giftpopSecret := "TUVESUFPTkU6UWpsU1RuUWxNa1lsTWtJM2JUaElSMWxaVWpkclMxaw=="
	limit := 50
	page := 1
	goodsSecret, _ := json.Marshal(map[string]interface{}{
		"authKey":    giftpopSecret,
		"sizeOfPage": limit,
		"page":       page,
	})
	header := map[string]string{
		"Content-Type": "apllication/json",
	}
	for i := 1; i < 100; i++ {
		code, _, err := SendReqPost("http://pos.giftpop.com.vn:9901/interface/order/goodsListAll.m12", header, goodsSecret)
		log.Print(code, err)
	}
	// code, body, err := SendReqPost("http://pos.giftpop.com.vn:9901/interface/order/goodsListAll.m12", header, goodsSecret)
}

func TestCreateHash(t *testing.T) {
	log.Print(HashRecord("tenguyen", 1, 5, "22"))
}
