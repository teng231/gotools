package crypt

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	val := "tenguyen pro"

	cirfer, _ := AESEncryptToBase64(key, val)
	log.Printf("%s => %s\n", val, cirfer)
}

func TestDescrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	text := "eqJ8mMZXPCTWLD7ZkWlvqUmT1KKDwTlBwYYAzg"

	bin, _ := AESDecryptToBase64(key, text)
	log.Printf("%s => %s\n", text, bin)
}
