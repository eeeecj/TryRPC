package Sim

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (sim *Sim) Create(c *gin.Context) {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = sim.proxyAddress
		req.Host = sim.proxyAddress
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
