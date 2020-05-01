package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var certFile = *flag.String("cert", "tls/test.crt", "TLS certificate")
	var keyFile = *flag.String("key", "tls/test.key", "TLS key")
	flag.Parse()

	r := gin.Default()

	r.GET("/hmac", func(c *gin.Context) {
		c.String(http.StatusOK, "ok.")
	})

	r.RunTLS("127.0.0.1:8080", certFile, keyFile)
}
