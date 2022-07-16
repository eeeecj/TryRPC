package middlewares

import (
	"github.com/TryRpc/pkg/Limiter"
	"github.com/gin-gonic/gin"
)

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"limiter": Limiter.DefaultLimiter.Limit(),
	}
}

var DefaultMiddleWares = defaultMiddlewares()
