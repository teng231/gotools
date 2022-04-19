package crypt

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	val := []byte("tenguyen pro")

	cirfer := Encrypt(key, val)
	log.Printf("%s => %s\n", val, cirfer)
}

func TestDescrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	text := "70013cfda80785e6-50efd8631272ce94b9b12d36-f5e67d18751057aa86a35a1fbc7e84a94addae48dc299cc792e08aeb"

	bin := Decrypt(key, []byte(text))
	log.Printf("%s => %s\n", text, bin)
}
