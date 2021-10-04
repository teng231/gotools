package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Hash using sha1 to encoding
func Hash(keys ...interface{}) string {
	kStr := []string{}
	for _, key := range keys {
		kStr = append(kStr, fmt.Sprintf("%v", key))
	}
	payload := strings.Join(kStr, "")
	h := sha256.New()
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
