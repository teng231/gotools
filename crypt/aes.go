package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func AESEncrypt(key []byte, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Padding văn bản
	text = pad(text)
	// log.Print(aes.BlockSize + len(text))
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	// log.Print(string(iv))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], text)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AESDecrypt(key []byte, encryptedText string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("Ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// log.Print(string(iv))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad văn bản
	plaintext, err := unpad(ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func pad(text []byte) []byte {
	padding := aes.BlockSize - len(text)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func unpad(text []byte) ([]byte, error) {
	length := len(text)
	unpadding := int(text[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("Invalid padding")
	}
	return text[:(length - unpadding)], nil
}
