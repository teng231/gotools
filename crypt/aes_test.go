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
	text := "SF3CXm++mlQyjB723QYWIWoaxJJRsNxvgSR8L3lB6L39hOedQwrZWLlYSj8hcWV9mZCKwqUmMPKIxNN9FWGrSwMNof9B9iMFCW8ONneeUwGeq2rkdwpA47vPgXdq+UQ53sTsPDnhuUyJhaxDuk/y0INiHX1m+XlklPlhYALN7WoYehV84zb/twA5FbLBh+y7ZQMp3V270nigaVWL6wXEMb09FWGGjCyIXr7WXGeZovPykSNCVnX5WRPWWCHGFvbKaGbmJyEx6j9ZuTX8QmfYcgJvlWU+P0Ffjs356A4M8L911CIdE4JpYqfMc616yS6mOKWrajf9j+uAhF60eaIJN/34lNXISXoACuML+fprR4Aic30aDOoPc95hHdPWKFWZWXda3fgYz3R/UvaHe/J7+GcnND/Njw2FGW740F6OQ5OOgEOQykpX0ECyoxoMUQon"

	bin, _ := AESDecrypt(key, text)
	log.Printf("%s => %s\n", text, bin)
}

func TestDescryptDev(t *testing.T) {
	key := []byte("6BgaZKcY6jj1vrwFUCmzzQP3uDdCPcxB")
	text := "IhpwMeETLCJwSl11WjnFhMJIAeYdf/7L18qBM6go7twHxIKAqTwbFrIvfMDYFxPLECTyH0MYTEIr0t3Bkc1uEUnycIW0vyiI+Uuzp52CcdEnTi9b0oBJJzGjThX2mAI2YEyTRj6N8tyL+brOewenT80dore+DTflvSGjzB+nZjaO1Sx5kH4IYFoj6UkVNCDrX8F5Z4mK3aaUnk5hUgbBxH8hmTtddtB0Te9pG7oz9YPGllIv6R3JgqgXY4xR1rqP4y0RZfuUE9EU/leBYE97MsVzpqsESnzqiZEassvWjBfc7D/4EikwCOl2jyBIjnbD"

	bin, _ := AESDecrypt(key, text)
	log.Printf("%s => %s\n", text, bin)
}
