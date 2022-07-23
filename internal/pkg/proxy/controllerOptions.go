package proxy

import (
	"fmt"
	"github.com/TryRpc/internal/pkg"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type ControllerOptions struct {
	BindAddr string      `json:"bind-addr" mapstructure:"bind-addr"`
	BindPort int         `json:"bind-port" mapstructure:"bind-port"`
	CertKey  pkg.CertKey `json:"tls" mapstructure:"tls"`
}

func NewControllerOptions() *ControllerOptions {
	return &ControllerOptions{
		BindAddr: "0.0.0.0",
		BindPort: 8083,
		CertKey:  pkg.CertKey{},
	}
}

func (s *ControllerOptions) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

func (s *ControllerOptions) Validate() []error {
	var errs []error
	if s.BindPort < 0 || s.BindPort > 2<<16 {
		errs = append(errs, fmt.Errorf("--insecure-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port", s.BindPort))
	}
	return errs
}

func (s *ControllerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddr, "controller.secure.bind-address", s.BindAddr,
		`The address for remote tunnel secure server`)
	fs.IntVar(&s.BindPort, "controller.secure.bind-port", s.BindPort,
		"The port for remote tunnel secure server")
	fs.StringVar(&s.CertKey.CertFile, "controller.secure.cert-file", s.CertKey.CertFile,
		"File containing the default x509 Certificate for HTTPS")
	fs.StringVar(&s.CertKey.KeyFile, "controller.secure.cert-key", s.CertKey.KeyFile,
		"File containing the default x509 private key matching")
	fs.StringVar(&s.CertKey.CaFile, "controller.secure.Ca-key", s.CertKey.CaFile,
		"File containing the default x509 CA matching")
}
