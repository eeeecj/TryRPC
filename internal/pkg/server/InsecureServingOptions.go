package server

import (
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type InsecureServingOptions struct {
	BindAddr string `json:"bind-address" mapstructure:"bind-address"`
	BindPort int    `json:"bind-port" mapstructure:"bind-port"`
}

func NewInsecureServingOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		BindAddr: "127.0.0.1",
		BindPort: 8080,
	}
}
func (s *InsecureServingOptions) Address() string {
	return net.JoinHostPort(s.BindAddr, strconv.Itoa(s.BindPort))
}

func (s *InsecureServingOptions) Validate() []error {
	var errs []error
	if s.BindPort < 0 || s.BindPort > 2<<16 {
		errs = append(errs, fmt.Errorf("--insecure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port", s.BindPort))
	}
	return errs
}

func (s *InsecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddr, "insecure.bind-address", s.BindAddr, "The IP address on which to serve the --insecure.bind-port "+
		"(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&s.BindPort, "insecure.bind-port", s.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")
}
