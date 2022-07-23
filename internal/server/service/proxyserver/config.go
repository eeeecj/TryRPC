package proxyserver

import (
	"github.com/TryRpc/internal/pkg"
	"github.com/TryRpc/internal/pkg/proxy"
	"net"
	"strconv"
)

type ProxyConfig struct {
	ControllerServing *ControllerServing
	RemoteServing     *RemoteServing
	LocalServing      *LocalServing
}

type ControllerServing struct {
	BindAddr string
	BindPort int
	CertKey  pkg.CertKey
}

func (c *ControllerServing) Address() string {
	return net.JoinHostPort(c.BindAddr, strconv.Itoa(c.BindPort))
}

type RemoteServing struct {
	BindAddr string
	BindPort int
	CertKey  pkg.CertKey
}

func (r *RemoteServing) Address() string {
	return net.JoinHostPort(r.BindAddr, strconv.Itoa(r.BindPort))
}

type LocalServing struct {
	BindAddr string
	BindPort int
}

func (l *LocalServing) Address() string {
	return net.JoinHostPort(l.BindAddr, strconv.Itoa(l.BindPort))
}

func NewProxyConfig() *ProxyConfig {
	return &ProxyConfig{}
}

func (p *ProxyConfig) New() (*ProxyServer, error) {
	server := &ProxyServer{
		RemoteServing:     p.RemoteServing,
		LocalServing:      p.LocalServing,
		ControllerServing: p.ControllerServing,
	}
	return server, nil
}

func (p *ProxyConfig) LocalToApply(s *proxy.LocalServerOption) error {
	p.LocalServing = &LocalServing{
		BindAddr: s.BindAddr,
		BindPort: s.BindPort,
	}
	return nil
}

func (p *ProxyConfig) RemoteToApply(s *proxy.RemoteServerOption) error {
	p.RemoteServing = &RemoteServing{
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

func (p *ProxyConfig) ControllerToApply(s *proxy.ControllerOptions) error {
	p.ControllerServing = &ControllerServing{
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

func (p *ProxyConfig) ToApply(s *proxy.ProxyServerOption) error {
	p.LocalToApply(s.LocalServer)
	p.RemoteToApply(s.RemoteServer)
	p.ControllerToApply(s.ControllerServer)
	return nil
}
