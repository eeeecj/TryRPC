package proxy

import (
	proxyServer "github.com/TryRpc/internal/local/ProxyServer"
	"github.com/TryRpc/internal/pkg"
	"github.com/TryRpc/internal/pkg/proxy"
	"net"
	"strconv"
)

type ProxyClientConfig struct {
	LocalClient      *LocalClientServing
	RemoteClient     *RemoteClientServing
	ControllerClient *ControllerClientServing
	PoolSize         int
}

func NewProxyClientConfig() *ProxyClientConfig {
	return &ProxyClientConfig{}
}

func (p *ProxyClientConfig) New() (*ProxyServer, error) {
	server := &ProxyServer{
		RemoteServing:     p.RemoteClient,
		LocalServing:      p.LocalClient,
		ControllerServing: p.ControllerClient,
		closing:           make(chan struct{}, 1),
		PoolSize:          p.PoolSize,
	}
	return server, nil
}

type LocalClientServing struct {
	BindAddr string
	BindPort int
	MaxConn  int
}

func (l *LocalClientServing) Address() string {
	return net.JoinHostPort(l.BindAddr, strconv.Itoa(l.BindPort))
}

type RemoteClientServing struct {
	BindAddr string
	BindPort int
	MaxConn  int
	CertKey  pkg.CertKey
}

func (r *RemoteClientServing) Address() string {
	return net.JoinHostPort(r.BindAddr, strconv.Itoa(r.BindPort))
}

type ControllerClientServing struct {
	BindAddr string
	BindPort int
	CertKey  pkg.CertKey
}

func (c *ControllerClientServing) Address() string {
	return net.JoinHostPort(c.BindAddr, strconv.Itoa(c.BindPort))
}

func (p *ProxyClientConfig) LocalToApply(s *proxy.LocalServerOption) error {
	p.LocalClient = &LocalClientServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
	}
	return nil
}
func (p *ProxyClientConfig) RemoteToApply(s *proxy.RemoteServerOption) error {
	p.RemoteClient = &RemoteClientServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
		CertKey: pkg.CertKey{
			s.CertKey.CertFile,
			s.CertKey.KeyFile,
			s.CertKey.CaFile,
		},
	}
	return nil
}

func (p *ProxyClientConfig) ControllerToApply(s *proxy.ControllerOptions) error {
	p.ControllerClient = &ControllerClientServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
		CertKey: pkg.CertKey{
			s.CertKey.CertFile,
			s.CertKey.KeyFile,
			s.CertKey.CaFile,
		},
	}
	return nil
}

func (p *ProxyClientConfig) ToApply(s *proxyServer.ProxyOption) error {
	p.PoolSize = s.PoolSize
	p.LocalToApply(s.LocalServerServing)
	p.RemoteToApply(s.RemoteServerServing)
	p.ControllerToApply(s.ControllerServerServing)
	return nil
}
