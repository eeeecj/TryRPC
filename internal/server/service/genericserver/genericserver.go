package genericserver

import (
	"github.com/TryRpc/internal/local/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

type GenericServer struct {
	Engine      *gin.Engine
	middlewares []string
}

func NewGenericServer() *GenericServer {
	return &GenericServer{Engine: gin.Default(), middlewares: []string{}}
}

func (g *GenericServer) Run(addr string) {
	g.Engine.Run(addr)
}

func (g *GenericServer) InstallMiddleWares() {
	for _, m := range g.middlewares {
		mw, ok := middlewares.DefaultMiddleWares[m]
		if !ok {
			log.Printf("can not find middleware")
			continue
		}
		g.Engine.Use(mw)
	}
}
