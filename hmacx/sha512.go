package hmacx

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha512(key []byte, data []byte) []byte {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha512Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha512(key, data))
}
