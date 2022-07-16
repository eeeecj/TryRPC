package main

import "github.com/TryRpc/internal/local"

func main() {
	app := local.NewApp()
	app.Prepare().Run()
}
