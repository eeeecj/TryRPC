package local

import (
	"github.com/TryRpc/internal/local/controller/v1/sim"
	"github.com/TryRpc/internal/local/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InstallRouter(g *gin.Engine) {
	pprof.Register(g)

	v1 := g.Group("/v1")
	{
		simv1 := v1.Group("/sims")
		{
			simController := sim.NewSimulationController()
			simv1.POST("/create", middlewares.DefaultMiddleWares["limiter"], simController.Create)
		}
	}
}
