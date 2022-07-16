package Sim

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (sim *Sim) Create(c *gin.Context) {
	ReverseProxy()(c)
}

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ssasa")
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1:8081"
			req.Host = "127.0.0.1:8081"
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
