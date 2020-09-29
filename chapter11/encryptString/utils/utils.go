package utils

import (
	"log"
	"encoding/base64"
	"crypto/cipher"
    "crypto/aes"
)

var initVector = []byte{100, 169, 67, 62, 174, 124, 204, 238, 226, 252, 14, 218, 87, 56, 98, 44}

func EncryptString(key []byte, text string) string{
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    plaintext := []byte(text)
    cfb := cipher.NewCFBEncrypter(block, initVector)
    ciphertext := make([]byte, len(plaintext))
    cfb.XORKeyStream(ciphertext, plaintext)
    return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptString(key []byte, text string) string{
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext, _ := base64.StdEncoding.DecodeString(text)
    cfb := cipher.NewCFBEncrypter(block, initVector)
    plaintext := make([]byte, len(ciphertext))
    cfb.XORKeyStream(plaintext, ciphertext)
    return string(plaintext)
}