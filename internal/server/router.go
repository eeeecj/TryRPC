package server

import (
	sim2 "github.com/TryRpc/internal/server/controller/v1/Sim"
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
			sim := sim2.NewSim()
			simv1.POST("/create", sim.Create)
			//simv1.POST("/get", simController.Get)
		}
	}
}
