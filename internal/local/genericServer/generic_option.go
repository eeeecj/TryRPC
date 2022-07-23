package genericServer

import (
	"fmt"
	"github.com/spf13/pflag"
)

type GenericServerOptions struct {
	BindAddr    string   `json:"bind-addr" mapstructure:"bind-addr"`
	BindPort    int      `json:"bind-port" mapstructure:"bind-port"`
	MiddleWares []string `json:"middlewares" mapstructure:"middlewares"`
	MaxConn     int      `json:"max-connection" mapstructure:"max-connection"`
}

func NewGenericServerOptions() *GenericServerOptions {
	return &GenericServerOptions{
		BindAddr: "0.0.0.0",
		BindPort: 8085,
		MaxConn:  5,
	}
}

func (o *GenericServerOptions) Validate() []error {
	if o == nil {
		return nil
	}
	var errs []error
	if o.BindPort < 1 || o.BindPort >= 2<<16 {
		errs = append(errs, fmt.Errorf("--secure.bind-port %v must be between 1 and 65535, inclusive. It cannot be turned off with 0", o.BindPort))
	} else if o.BindPort < 0 || o.BindPort >= 2<<16 {
		errs = append(errs, fmt.Errorf("--secure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off secure port", o.BindPort))
	}
	return errs
}

func (o *GenericServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddr, "server.bind-address", o.BindAddr, ""+
		"The IP address on which to listen for the --secure.bind-port port. The "+
		"associated interface(s) must be reachable by the rest of the engine, and by CLI/web "+
		"clients. If blank, all interfaces will be used (0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&o.BindPort, "server.bind-port", o.BindPort, "The port for https server")
	fs.StringSliceVar(&o.MiddleWares, "server.middlewares", o.MiddleWares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
	fs.IntVar(&o.MaxConn, "server.max-connection", o.MaxConn, "The limiter of the tasks in queue")
}
