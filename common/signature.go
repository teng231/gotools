package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// Hash using sha1 to encoding
func Hash(keys ...interface{}) string {
	kStr := []string{}
	for _, key := range keys {
		kStr = append(kStr, fmt.Sprintf("%v", key))
	}
	payload := strings.Join(kStr, "")
	log.Print(payload)
	h := sha256.New()
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
