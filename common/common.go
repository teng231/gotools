package common

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/metadata"
)

// ContextValue func
func ContextValue(ctx context.Context, out proto.Message) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || md == nil || md["ctx"] == nil || len(md["ctx"]) == 0 {
		out = nil
		return nil
	}
	if err := json.Unmarshal([]byte(md["ctx"][0]), out); err != nil {
		out = nil
		return err
	}
	return nil
}
func MakeContext(sec int, claims proto.Message) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	if claims != nil {
		bin, err := json.Marshal(claims)
		if err != nil {
			log.Print(err)
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "ctx", string(bin))
		return ctx, cancel
	}
	return ctx, cancel
}

// ToNs convert time to nanosecond
func ToNs(timestamp int64) int64 {
	// ns
	if timestamp > 1e18 {
		return timestamp
	}
	// ms
	if timestamp > 1e12 {
		return timestamp * 1e6
	}
	// s
	if timestamp > 1e9 {
		return timestamp * 1e9
	}
	return -1
}

// ToMs convert time to milisecond
func ToMs(timestamp int64) int64 {
	// ns
	if timestamp > 1e18 {
		return timestamp / 1e6
	}
	// ms
	if timestamp > 1e12 {
		return timestamp
	}
	// s
	if timestamp > 1e9 {
		return timestamp * 1e3
	}
	return -1
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func HashRecord(items ...interface{}) string {
	lenItems := len(items)
	payload := ""
	for i := range items {
		if i == lenItems-1 {
			payload += "%v"
			break
		}
		payload += "%v_"
	}
	// log.Print(payload)
	has := md5.Sum([]byte(fmt.Sprintf(payload, items...)))
	return fmt.Sprintf("%x", has)
}

func UniqInt32(intSlice []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqInt64(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func MergeStructs(in ...interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	temp := make(map[string]interface{})
	for _, m := range in {
		bin, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(bin, &temp); err != nil {
			return nil, err
		}
		for k, v := range temp {
			result[k] = v
		}
	}
	return result, nil
}
