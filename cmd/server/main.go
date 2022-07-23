package main

import "github.com/TryRpc/internal/server"

func main() {
	server.NewApp("proxy-server").Run()
}
