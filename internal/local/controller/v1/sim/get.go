package sim

import (
	"github.com/TryRpc/pkg/Limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sim *Simulation) Get(c *gin.Context) {
	data := Limiter.DefaultLimiter.ReleaseLimiter()
	c.JSON(http.StatusOK, data)
}
