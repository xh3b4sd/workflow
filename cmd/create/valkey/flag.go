package valkey

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/runtime"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Version struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Version.Golang == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-g/--version-golang must not be empty"})
		}

		s := strings.Split(f.Version.Golang, ".")
		if len(s) != 3 {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-g/--version-golang must have 3 parts like 1.15.2"})
		}
	}

	return nil
}
