package app

import mflag "github.com/TryRpc/component/pkg/cli/flag"

type CliOption interface {
	Flags() (fs mflag.NamedFlagSets)
	Validate() []error
}
