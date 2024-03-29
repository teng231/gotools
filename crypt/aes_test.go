package crypt

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	val := []byte("chao ban golang")

	cirfer, err := AESEncrypt(key, val)
	log.Printf("%s => %s\n %v", val, cirfer, err)
}

func TestDescrypt(t *testing.T) {
	key := []byte("1234567890qwedavsfrtyjncvt67uuhs")
	text := "FYJQzdWuFT3V06EigANhyUjjiVADdNU6gegHI8KS5Uo="

	bin, _ := AESDecrypt(key, text)
	log.Printf("%s => %s\n", text, bin)
}
