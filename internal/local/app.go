package local

import (
	"fmt"
	"github.com/TryRpc/internal/local/service/Proxy"
	"github.com/TryRpc/internal/local/service/Worker"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
)

type App struct {
	GinServer *gin.Engine
	Proxy     *Proxy.Proxy
	Worker    *Worker.Worker
}

func NewApp() *App {
	return &App{
		GinServer: gin.Default(),
		Proxy:     Proxy.NewProxy(),
		Worker:    Worker.NewWorker(),
	}
}

func (app *App) Prepare() *App {
	InstallRouter(app.GinServer)
	return app
}

func (app *App) Run() {
	fmt.Println(app.Proxy)
	var e errgroup.Group
	e.Go(func() error {
		return app.GinServer.Run(":7887")
	})
	e.Go(func() error {
		app.Proxy.Start()
		return nil
	})
	e.Go(func() error {
		app.Worker.Start()
		return nil
	})
	if err := e.Wait(); err != nil {
		log.Fatal(err)
	}
}
