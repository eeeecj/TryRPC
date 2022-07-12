package sim

import (
	"github.com/TryRpc/internal/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sim *Simulation) Create(c *gin.Context) {
	var r middlewares.GrpcData
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	middlewares.DefaultLimiter.GetConn(&r)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
