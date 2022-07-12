package main

import "github.com/TryRpc/internal/server"

func main() {
	app := server.New()
	app.InstallMiddleWares()
	server.InitRouter(app.Engine)
	app.Run()
}
