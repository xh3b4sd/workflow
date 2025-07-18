package npm

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/runtime"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Version struct {
		Node string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Version.Node, "version-node", "n", version.Node, "Node version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Version.Node == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-n/--version-node must not be empty"})
		}

		s := strings.Split(f.Version.Node, ".")
		if len(s) != 3 {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-n/--version-node must have 3 parts like 1.15.2"})
		}
	}

	return nil
}
