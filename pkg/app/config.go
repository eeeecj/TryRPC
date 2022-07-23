package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const configName = "config"

var config string

func init() {
	pflag.StringVarP(&config, configName, "c", config,
		"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

func readConfig(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configName))

	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		if config != "" {
			viper.SetConfigFile(config)
		} else {
			viper.AddConfigPath(".")
			viper.AddConfigPath("./config/")
			viper.AddConfigPath("./config/server")
			viper.SetConfigName(basename)
		}
		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", config, err)
			os.Exit(1)
		}

	})
}
