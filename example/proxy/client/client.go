package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "120.24.250.251:7238")
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("hello"))
	b, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1"
			req.Host = "127.0.0.1:8080"
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
