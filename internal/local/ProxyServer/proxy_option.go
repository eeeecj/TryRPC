package ProxyServer

import (
	"fmt"
	"github.com/TryRpc/internal/pkg/proxy"
	"github.com/spf13/pflag"
)

type ProxyOption struct {
	LocalServerServing      *proxy.LocalServerOption  `json:"local" mapstructure:"local"`
	RemoteServerServing     *proxy.RemoteServerOption `json:"remote" mapstructure:"remote"`
	ControllerServerServing *proxy.ControllerOptions  `json:"controller" mapstructure:"controller"`
	PoolSize                int                       `json:"pool-size" mapstructure:"pool-size"`
}

func NewProxyOption() *ProxyOption {
	return &ProxyOption{
		LocalServerServing:      proxy.NewLocalServerOption(),
		RemoteServerServing:     proxy.NewRemoteServerOption(),
		ControllerServerServing: proxy.NewControllerOptions(),
		PoolSize:                50,
	}
}

func (o *ProxyOption) Validate() []error {
	var errs []error
	if o.PoolSize <= 0 {
		errs = append(errs, fmt.Errorf("--pool-size must greater than 0."))
	}
	errs = append(errs, o.LocalServerServing.Validate()...)
	errs = append(errs, o.RemoteServerServing.Validate()...)
	errs = append(errs, o.ControllerServerServing.Validate()...)
	return errs
}
func (o *ProxyOption) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&o.PoolSize, "pool-size", o.PoolSize, "The max number of proxy connection")
	o.LocalServerServing.AddFlags(fs)
	o.RemoteServerServing.AddFlags(fs)
	o.ControllerServerServing.AddFlags(fs)
}
