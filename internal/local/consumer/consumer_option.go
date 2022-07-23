package consumer

import (
	"fmt"
	"github.com/spf13/pflag"
)

type ConsumerOption struct {
	MaxConsumer int    `json:"max-consumer" mapstructure:"max-consumer"`
	BindAddr    string `json:"bind-addr" mapstructure:"bind-addr"`
	BindPort    int    `json:"bind-port" mapstructure:"bind-port"`
}

func NewConsumerOption() *ConsumerOption {
	return &ConsumerOption{MaxConsumer: 2}
}
func (o *ConsumerOption) Validate() []error {
	var errs []error
	if o.MaxConsumer <= 0 {
		errs = append(errs, fmt.Errorf("--consumer.maxconsumer must greater than 0"))
	}
	return errs
}

func (o *ConsumerOption) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&o.MaxConsumer, "consumer.maxconsumer", o.MaxConsumer, "The number of simulations run concurrently")
}
