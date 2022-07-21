package genericserver

import (
	"context"
	"github.com/TryRpc/internal/local/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type GenericServer struct {
	middlewares []string
	Engine      *gin.Engine
	Server      *http.Server
}

func NewGenericServer() *GenericServer {
	return &GenericServer{Engine: gin.Default(), middlewares: []string{}}
}

func (g *GenericServer) Run(addr string) error {
	g.Server = &http.Server{Addr: ":8080", Handler: g.Engine}
	if err := g.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
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

func (g *GenericServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := g.Server.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
