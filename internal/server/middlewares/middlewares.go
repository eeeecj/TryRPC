package middlewares

import "github.com/gin-gonic/gin"

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"limiter": DefaultLimiter.Limit(),
	}
}

var DefaultMiddleWares = defaultMiddlewares()
