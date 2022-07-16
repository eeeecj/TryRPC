package app

import (
	"github.com/TryRpc/internal/server/middlewares"
	"github.com/TryRpc/internal/server/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
)

type App struct {
	Engine      *gin.Engine
	Middlewares []string
	Proxy       *service.Proxy
}

func (app *App) Run() {
	var e errgroup.Group
	e.Go(func() error {
		return app.Engine.Run(":8080")
	})
	e.Go(func() error {
		app.Proxy.Run()
		return nil
	})

	if err := e.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (app *App) InstallMiddleWares() {
	for _, m := range app.Middlewares {
		mw, ok := middlewares.DefaultMiddleWares[m]
		if !ok {
			log.Printf("can not find middleware")
			continue
		}
		app.Engine.Use(mw)
	}
}
