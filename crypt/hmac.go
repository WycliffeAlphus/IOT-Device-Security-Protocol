package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyHMAC (message, key, expectedMAC string) bool {
	actualMAC := GenerateHMAC(message, key)
	return hmac.Equal([]byte(actualMAC), []byte(expectedMAC))
}
