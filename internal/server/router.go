package server

import (
	"github.com/TryRpc/internal/server/controller/v1/sim"
	"github.com/TryRpc/internal/server/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) {
	pprof.Register(g)
	v1 := g.Group("/v1")
	{
		simv1 := v1.Group("/sims")
		{
			simController := sim.NewSimulationController()
			simv1.POST("/create", middlewares.DefaultMiddleWares["limiter"], simController.Create)
			//simv1.POST("/get", simController.Get)
		}
	}
}
