package middleware

import (
	"fmt"
	"github.com/TryRpc/api/grpc"
	"github.com/TryRpc/internal/local/genericServer"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	defaultLimiter *Limiter
	once           sync.Once
)

type Limiter struct {
	maxConn int
	bucket  chan *grpc.GrpcData
}

func GetLimiter(opt *genericServer.GenericServerOptions) (*Limiter, error) {
	if opt == nil && defaultLimiter == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}
	var err error
	once.Do(func() {
		defaultLimiter = &Limiter{
			maxConn: opt.MaxConn,
			bucket:  make(chan *grpc.GrpcData, opt.MaxConn),
		}
	})
	if defaultLimiter == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, errors: %w", defaultLimiter, err)
	}
	return defaultLimiter, nil
}

func (l *Limiter) GetConn(data *grpc.GrpcData) {
	l.bucket <- data
}

func (l *Limiter) ReleaseLimiter() *grpc.GrpcData {
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
