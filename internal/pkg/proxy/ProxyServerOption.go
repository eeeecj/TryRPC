package proxy

import (
	"github.com/spf13/pflag"
)

type ProxyServerOption struct {
	RemoteServer     *RemoteServerOption `json:"remote" mapstructure:"remote"`
	LocalServer      *LocalServerOption  `json:"local" mapstructure:"local"`
	ControllerServer *ControllerOptions  `json:"controller" mapstructure:"controller"`
}

func NewProxyServerOption() *ProxyServerOption {
	return &ProxyServerOption{
		RemoteServer:     NewRemoteServerOption(),
		LocalServer:      NewLocalServerOption(),
		ControllerServer: NewControllerOptions(),
	}
}

func (s *ProxyServerOption) Validate() []error {
	var errs []error
	errs = append(errs, s.LocalServer.Validate()...)
	errs = append(errs, s.RemoteServer.Validate()...)
	errs = append(errs, s.ControllerServer.Validate()...)
	return errs
}

func (s *ProxyServerOption) AddFlags(fs *pflag.FlagSet) {
	s.LocalServer.AddFlags(fs)
	s.RemoteServer.AddFlags(fs)
}
