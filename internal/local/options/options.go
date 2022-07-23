package options

import (
	"encoding/json"
	mflag "github.com/TryRpc/component/pkg/cli/flag"
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/internal/local/ProxyServer"
	"github.com/TryRpc/internal/local/consumer"
	"github.com/TryRpc/internal/local/genericServer"
)

type Options struct {
	GenericServerServing *genericServer.GenericServerOptions `json:"server" mapstructure:"server"`
	ProxyServerServing   *ProxyServer.ProxyOption            `json:"proxy" mapstructure:"proxy"`
	ConsumerOptions      *consumer.ConsumerOption            `json:"consumer" mapstructure:"consumer"`
	Log                  *cuszap.Options                     `json:"log" mapstructure:"log"`
}

func NewOptions() *Options {
	return &Options{
		GenericServerServing: genericServer.NewGenericServerOptions(),
		ProxyServerServing:   ProxyServer.NewProxyOption(),
		ConsumerOptions:      consumer.NewConsumerOption(),
		Log:                  cuszap.NewOptions(),
	}
}
func (o *Options) Flags() (fss mflag.NamedFlagSets) {
	o.GenericServerServing.AddFlags(fss.FlagSet("generic"))
	o.ProxyServerServing.AddFlags(fss.FlagSet("proxy"))
	o.ConsumerOptions.AddFlags(fss.FlagSet("consumer"))
	o.Log.AddFlags(fss.FlagSet("log"))
	return
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.GenericServerServing.Validate()...)
	errs = append(errs, o.ProxyServerServing.Validate()...)
	errs = append(errs, o.ConsumerOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return errs
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
