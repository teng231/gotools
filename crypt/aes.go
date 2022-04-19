package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/xdg-go/pbkdf2"
)

// source here: https://gist.github.com/enyachoke/5c60f5eebed693d9b4bacddcad693b47

func deriveKey(passphrase, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func Encrypt(passphrase, plaintext []byte) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, plaintext, nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func Decrypt(passphrase, ciphertext []byte) []byte {
	arr := strings.Split(string(ciphertext), "-")
	if len(arr) < 3 {
		return ciphertext
	}
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return data
}
