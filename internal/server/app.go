package server

import (
	"github.com/TryRpc/internal/server/service/Proxy"
	"github.com/TryRpc/internal/server/service/genericserver"
	"golang.org/x/sync/errgroup"
	"log"
)

func New() *App {
	return &App{
		genericserver.NewGenericServer(),
		[]string{},
		Proxy.NewProxy(),
	}
}

type App struct {
	GenericServer *genericserver.GenericServer
	Middlewares   []string
	Proxy         *Proxy.Proxy
}

func (app *App) Prepare() *App {
	app.GenericServer.InstallMiddleWares()
	InitRouter(app.GenericServer.Engine)
	return app
}

func (app *App) Run() {
	var e errgroup.Group
	e.Go(func() error {
		app.GenericServer.Run(":8080")
		return nil
	})
	e.Go(func() error {
		app.Proxy.Start()
		return nil
	})

	if err := e.Wait(); err != nil {
		log.Fatal(err)
	}
}
