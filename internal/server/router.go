package server

import (
	"github.com/TryRpc/component/pkg/core"
	"github.com/TryRpc/component/pkg/errors"
	"github.com/TryRpc/internal/pkg/code"
	sim2 "github.com/TryRpc/internal/server/controller/v1/Sim"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine, Address string) {
	installController(g, Address)
}

func installController(g *gin.Engine, Address string) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})
	v1 := g.Group("/v1")
	{
		simv1 := v1.Group("/sims")
		{
			sim := sim2.NewSim(Address)
			simv1.POST("/create", sim.Create)
			//simv1.POST("/get", simController.Get)
			//simv1.POST("/create", func(c *gin.Context) {
			//	c.JSON(http.StatusOK, gin.H{
			//		"message": "ok",
			//	})
			//})
		}
	}
}
