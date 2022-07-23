package server

import (
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/internal/server/config"
	"github.com/TryRpc/internal/server/options"
	"github.com/TryRpc/pkg/app"
)

const commandDesc = `The server proxy the http request to the local server`

func NewApp(basename string) *app.App {
	opt := options.NewOptions()
	application := app.NewApp(basename, "Proxy Server",
		app.WithOptions(opt),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opt)))
	return application
}

func run(opt *options.Options) app.RunFunc {
	return func(basename string) error {
		cuszap.Init(opt.Log)
		defer cuszap.Flush()
		cfg := config.NewConfig(opt)
		return Run(cfg)
	}
}
