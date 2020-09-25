package golang

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Version string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Version, "version", "v", "1.15.2", "Golang version to set in, e.g. go.mod.")
}

func (f *flag) Validate() error {
	if f.Version == "" {
		return tracer.Maskf(invalidFlagError, "-v/--version must not be empty")
	}

	s := strings.Split(f.Version, ".")
	if len(s) != 3 {
		return tracer.Maskf(invalidFlagError, "-v/--version must have 3 parts like 1.15.2")
	}

	return nil
}
