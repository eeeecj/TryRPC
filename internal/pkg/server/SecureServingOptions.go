package server

import (
	"fmt"
	"github.com/TryRpc/internal/pkg"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type SecureServingOptions struct {
	BindAddr   string      `json:"bind-addr" mapstructure:"bind-address"`
	BindPort   int         `json:"bind-port" mapstructure:"bind-port"`
	ServerCert pkg.CertKey `json:"tls" mapstructure:"tls"`
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddr:   "0.0.0.0",
		BindPort:   8080,
		ServerCert: pkg.CertKey{},
	}
}
func (s *SecureServingOptions) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

func (s *SecureServingOptions) Validate() []error {
	var errs []error
	if s.BindPort < 0 || s.BindPort > 2<<16 {
		errs = append(errs, fmt.Errorf("--insecure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port", s.BindPort))
	}
	return errs
}

func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddr, "secure.bind-address", s.BindAddr, ""+
		"The IP address on which to listen for the --secure.bind-port port. The "+
		"associated interface(s) must be reachable by the rest of the engine, and by CLI/web "+
		"clients. If blank, all interfaces will be used (0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&s.BindPort, "secure.bind-port", s.BindPort, "The port for https server")
	fs.StringVar(&s.ServerCert.CertFile, "secure.tls.cert-key.cert-file", s.ServerCert.CertFile, ""+
		"File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
		"after server cert).")
	fs.StringVar(&s.ServerCert.KeyFile, "secure.tls.cert-key.private-key-file", s.ServerCert.KeyFile, ""+
		"File containing the default x509 private key matching --secure.tls.cert-key.cert-file.")
	fs.StringVar(&s.ServerCert.CaFile, "secure.tls.cert-key.ca-key", s.ServerCert.CaFile, ""+
		"File containing the default x509 private key matching --secure.tls.cert-key.ca-key.")
}
