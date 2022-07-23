package server

import (
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	SecureServingOptions   *SecureServingOptions   `json:"secure" mapstructure:"secure"`
	InsecureServingOptions *InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	MiddleWares            []string                `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		MiddleWares:            []string{},
		SecureServingOptions:   NewSecureServingOptions(),
		InsecureServingOptions: NewInsecureServingOptions(),
	}
}

func (s *ServerRunOptions) Validate() []error {
	var errs []error
	errs = append(errs, s.SecureServingOptions.Validate()...)
	errs = append(errs, s.InsecureServingOptions.Validate()...)
	return errs
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringSliceVar(&s.MiddleWares, "server.middlewares", s.MiddleWares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
	s.SecureServingOptions.AddFlags(fs)
	s.InsecureServingOptions.AddFlags(fs)
}
