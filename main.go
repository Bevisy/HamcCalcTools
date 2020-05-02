package main

import (
	"HmacCalcTools/utils"
	"encoding/base64"
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var certFile = *flag.String("cert", "tls/test.crt", "TLS certificate")
	var keyFile = *flag.String("key", "tls/test.key", "TLS key")
	flag.Parse()

	r := gin.Default()

	// 计算 hmac 值，返回值类型为 []byte
	r.POST("/hmac", func(c *gin.Context) {
		algorithm := c.DefaultQuery("algorithm", "hmac-sha1")
		key := c.DefaultQuery("key", "")
		message := c.Query("msg")

		switch algorithm {
		case "hmac-sha1":
			c.JSON(http.StatusOK, gin.H{
				"message":   message,
				"key":       key,
				"algorithm": algorithm,
				"HMAC-SHA1": utils.GenHmacSha1(message, key),
			})
		case "hmac-sha256":
			c.JSON(http.StatusOK, gin.H{
				"message":     message,
				"key":         key,
				"algorithm":   algorithm,
				"HMAC-SHA256": utils.GenHmacSha256(message, key),
			})
		case "hmac-sha384":
			c.JSON(http.StatusOK, gin.H{
				"message":     message,
				"key":         key,
				"algorithm":   algorithm,
				"HMAC-SHA384": utils.GenHmacSha384(message, key),
			})
		case "hmac-sha512":
			c.JSON(http.StatusOK, gin.H{
				"message":     message,
				"key":         key,
				"algorithm":   algorithm,
				"HMAC-SHA512": utils.GenHmacSha512(message, key),
			})
		default:
			c.String(http.StatusBadRequest, "Unknown algorithm type: %s\n", algorithm)
		}

	})
	// 计算 base64
	r.POST("/base64", func(c *gin.Context) {
		msg := c.Query("message")
		if msg == "" {
			c.String(http.StatusBadRequest, "The param:message can't be nil")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
				"base64":  base64.StdEncoding.EncodeToString([]byte(msg)),
			})
		}

	})

	r.RunTLS("127.0.0.1:8080", certFile, keyFile)
}
