package utils

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
	"log"
)

func GenHmacSha1(message, key string) []byte {
	k := []byte(key)
	h := hmac.New(sha1.New, k)
	if _, err := io.WriteString(h, message); err != nil {
		log.Fatalf("HMAC-SHA1 error: %s\n", err)
	}
	return h.Sum(nil)
}

func GenHmacSha256(message, key string) []byte {
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	if _, err := io.WriteString(h, message); err != nil {
		log.Fatalf("HMAC-SHA256 error: %s\n", err)
	}
	return h.Sum(nil)
}

func GenHmacSha384(message, key string) []byte {
	k := []byte(key)
	h := hmac.New(crypto.SHA384.New, k)
	if _, err := io.WriteString(h, message); err != nil {
		log.Fatalf("HMAC-SHA384 error: %s\n", err)
	}
	return h.Sum(nil)
}

func GenHmacSha512(message, key string) []byte {
	k := []byte(key)
	h := hmac.New(sha512.New, k)
	if _, err := io.WriteString(h, message); err != nil {
		log.Fatalf("HMAC-SHA512 error: %s\n", err)
	}
	return h.Sum(nil)
}
