package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Limiter struct {
	maxConn int
	bucket  chan *GrpcData
}

var DefaultLimiter = NewLimiter(10)

func NewLimiter(c int) *Limiter {
	return &Limiter{
		maxConn: c,
		bucket:  make(chan *GrpcData, c),
	}
}

func (l *Limiter) GetConn(data *GrpcData) {
	l.bucket <- data
}

func (l *Limiter) ReleaseLimiter() *GrpcData {
	fmt.Println(len(l.bucket))
	c := <-l.bucket
	return c
}
func (l *Limiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(l.bucket) >= l.maxConn {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
