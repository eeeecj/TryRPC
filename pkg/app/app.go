package app

import (
	"github.com/TryRpc/internal/server/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Engine      *gin.Engine
	Middlewares []string
}

func (app *App) Run() {
	app.Engine.Run(":8080")
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
