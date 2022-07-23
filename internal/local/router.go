package local

import (
	"github.com/TryRpc/internal/local/controller/v1/sim"
	"github.com/TryRpc/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) {
	//pprof.Register(g)
	v1 := g.Group("/v1")
	{
		simv1 := v1.Group("/sims")
		{
			limiter, _ := middleware.GetLimiter(nil)
			simController := sim.NewSimulationController()
			simv1.POST("/create", limiter.Limit(), simController.Create)
		}
	}
}
