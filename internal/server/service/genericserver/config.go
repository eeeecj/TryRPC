package genericserver

import (
	"github.com/TryRpc/internal/pkg"
	"github.com/TryRpc/internal/pkg/proxy"
	"github.com/TryRpc/internal/pkg/server"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
)

type GenericConfig struct {
	MiddleWare      []string
	SecureServing   *SecureServing
	InsecureServing *InsecureServing
	LocalServing    *LocalServing
}

type SecureServing struct {
	BindAddr   string
	BindPort   int
	ServerCert pkg.CertKey
}

func (s *SecureServing) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

type InsecureServing struct {
	BindAddr string
	BindPort int
}

func (s *InsecureServing) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

type LocalServing struct {
	BindAddr string
	BindPort int
}

func (s *LocalServing) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}
func NewGenericConfig() *GenericConfig {
	return &GenericConfig{
		MiddleWare: []string{},
	}
}

func (s *GenericConfig) New() (*GenericServer, error) {
	server := &GenericServer{
		Engine:          gin.New(),
		Middlewares:     s.MiddleWare,
		SecureServing:   s.SecureServing,
		InsecureServing: s.InsecureServing,
		ShutdownTimeOut: 10,
		LocalServing:    s.LocalServing,
	}
	initGenericServer(server)
	return server, nil
}

func (cfg *GenericConfig) LocalToApply(s *proxy.LocalServerOption) error {
	cfg.LocalServing = &LocalServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
	}
	return nil
}

func (cfg *GenericConfig) SecureToApply(s *server.SecureServingOptions) error {
	cfg.SecureServing = &SecureServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
		ServerCert: pkg.CertKey{
			s.ServerCert.CertFile,
			s.ServerCert.KeyFile,
			s.ServerCert.CaFile,
		},
	}
	return nil
}

func (cfg *GenericConfig) InsecureToApply(s *server.InsecureServingOptions) error {
	cfg.InsecureServing = &InsecureServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
	}
	return nil
}

func (cfg *GenericConfig) ToApply(s *server.ServerRunOptions) error {
	cfg.MiddleWare = s.MiddleWares
	cfg.SecureToApply(s.SecureServingOptions)
	cfg.InsecureToApply(s.InsecureServingOptions)
	return nil
}
