package proxy

import (
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type LocalServerOption struct {
	BindAddr string `json:"bind-addr" mapstructure:"bind-addr"`
	BindPort int    `json:"bind-port" mapstructure:"bind-port"`
}

func NewLocalServerOption() *LocalServerOption {
	return &LocalServerOption{
		BindAddr: "127.0.0.1",
		BindPort: 8082,
	}
}

func (s *LocalServerOption) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}
func (s *LocalServerOption) Validate() []error {
	var errs []error
	if s.BindPort < 0 || s.BindPort > 2<<16 {
		errs = append(errs, fmt.Errorf("--insecure-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port", s.BindPort))
	}
	return errs
}

func (s *LocalServerOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddr, "local.proxy.bind-addr", s.BindAddr,
		`The address for local proxy server`)
	fs.IntVar(&s.BindPort, "local.proxy.bind-port", s.BindPort,
		`the port for local proxy server`)
}
