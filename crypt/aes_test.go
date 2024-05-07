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
	key := []byte("mpDYXe9lnez8xeAAU3Op5V5mwn1bpHoo")
	text := "W72bkVHEGfHc2FQVZl2WjTJ2e3dIixBQAYRYVtyV3jQIPJDDLhcl2XasJskk82mn+vTXX2JqSJf5cUInu0qXTvO/EU9UU46sJvvb0/3g4TMmlW4OPVPYi74R+5Pk8Et29iO1FOGnch1TjmHjxVa5yRB/pliCjJXCPnfB5JTK/4T5NePwaULZG2CVS89a0XvdyXi95CCrv0bYe878nypoYSA9OJeSGyUASq2pF9Gf5wLB7eM/9qgGpqpguvd21G92ifYrloMSwKgGpuWvGaYfO1WE1bR6eFtoJ/TdeOehasaoPVkiqEch5vesdxSgqEmNEw97M3vT5MkQEfGTJRTf9w=="

	bin, _ := AESDecrypt(key, text)
	log.Printf("%s => %s\n", text, bin)
}
