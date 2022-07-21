package Proxy

import (
	"github.com/TryRpc/pkg/service"
	"github.com/fatih/pool"
	"net"
)

var localpool, _ = pool.NewChannelPool(0, 50, localFactory)

func localFactory() (net.Conn, error) {
	return service.CreateConn(option.Local)
}

var remotePool, _ = pool.NewChannelPool(0, 50, remoteFactory)

func remoteFactory() (net.Conn, error) {
	return service.CreateConn(option.Remote)
}
