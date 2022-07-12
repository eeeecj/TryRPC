package sim

import (
	"github.com/TryRpc/internal/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sim *Simulation) Get(c *gin.Context) {
	data := middlewares.DefaultLimiter.ReleaseLimiter()
	c.JSON(http.StatusOK, data)
}
