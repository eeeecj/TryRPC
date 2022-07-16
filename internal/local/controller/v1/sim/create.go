package sim

import (
	"github.com/TryRpc/pkg/Limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sim *Simulation) Create(c *gin.Context) {
	var r Limiter.GrpcData
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	Limiter.DefaultLimiter.GetConn(&r)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
