package options

import (
	"encoding/json"
	mflag "github.com/TryRpc/component/pkg/cli/flag"
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/internal/pkg/proxy"
	"github.com/TryRpc/internal/pkg/server"
	"github.com/TryRpc/internal/server/service/genericserver"
)

type Options struct {
	GenericServerRunOptions *server.ServerRunOptions `json:"server" mapstructure:"server"`
	ProxyServer             *proxy.ProxyServerOption `json:"proxy" mapstructure:"proxy"`
	Log                     *cuszap.Options          `json:"log" mapstructure:"log"`
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.GenericServerRunOptions.Validate()...)
	errs = append(errs, o.ProxyServer.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return errs
}

func (o *Options) ApplyTo(c *genericserver.GenericConfig) error {
	return nil
}

func NewOptions() *Options {
	o := &Options{
		GenericServerRunOptions: server.NewServerRunOptions(),
		ProxyServer:             proxy.NewProxyServerOption(),
		Log:                     cuszap.NewOptions(),
	}
	return o
}

func (o *Options) Flags() (fs mflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fs.FlagSet("generic"))
	o.ProxyServer.AddFlags(fs.FlagSet("proxy"))
	o.Log.AddFlags(fs.FlagSet("log"))
	return fs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
