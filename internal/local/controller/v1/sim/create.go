package sim

import (
	"github.com/TryRpc/api/grpc"
	"github.com/TryRpc/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sim *Simulation) Create(c *gin.Context) {
	var r grpc.GrpcData
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	limiter, _ := middleware.GetLimiter(nil)
	limiter.GetConn(&r)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
