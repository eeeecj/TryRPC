package server

import (
	"github.com/TryRpc/internal/local/genericServer"
	"github.com/TryRpc/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
)

type GenericConfig struct {
	MiddleWare []string
	BindAddr   string
	BindPort   int
	MaxConn    int
}

func NewGenericConfig() *GenericConfig {
	return &GenericConfig{
		MiddleWare: []string{},
	}
}
func (s *GenericConfig) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

func (g *GenericConfig) ToApply(s *genericServer.GenericServerOptions) error {
	g.BindAddr = s.BindAddr
	g.BindPort = s.BindPort
	g.MiddleWare = s.MiddleWares
	g.MaxConn = s.MaxConn
	return nil
}
func (s *GenericConfig) New() (*GenericServer, error) {
	server := &GenericServer{
		Engine:          gin.New(),
		Middlewares:     s.MiddleWare,
		InsecureServing: s,
		ShutdownTimeOut: 10,
	}
	initGenericServer(server)
	op := &genericServer.GenericServerOptions{
		BindAddr:    s.BindAddr,
		BindPort:    s.BindPort,
		MiddleWares: s.MiddleWare,
		MaxConn:     s.MaxConn,
	}
	middleware.GetLimiter(op)
	return server, nil
}
