package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func main() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"xiaoming": "hello",
		})
	})
	g.Run(":7810")
}
func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ssasa")
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1:7809"
			req.Host = "127.0.0.1:7809"
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
