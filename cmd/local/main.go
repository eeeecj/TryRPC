package main

import "github.com/TryRpc/internal/local"

func main() {
	local.NewApp("proxy-client").Run()
}
