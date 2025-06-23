package golang

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Binary  bool
	Private string
	User    string
	Version struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.Binary, "binary", "b", true, "Whether to add the 'go build .' or 'go build ./...' build step.")
	cmd.Flags().StringVarP(&f.Private, "private", "p", "", "GOPRIVATE string, e.g. github.com/org/rep.")
	cmd.Flags().StringVarP(&f.User, "user", "u", "", "Github user for GOPRIVATE authentication, e.g. xh3b4sd.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Private != "" && f.User == "" {
			return tracer.Maskf(invalidFlagError, "-p/--private and -u/--user must be given together")
		}
		if f.Private == "" && f.User != "" {
			return tracer.Maskf(invalidFlagError, "-p/--private and -u/--user must be given together")
		}
	}

	{
		if f.Version.Golang == "" {
			return tracer.Maskf(invalidFlagError, "-g/--version-golang must not be empty")
		}

		s := strings.Split(f.Version.Golang, ".")
		if len(s) != 3 {
			return tracer.Maskf(invalidFlagError, "-g/--version-golang must have 3 parts like 1.15.2")
		}
	}

	return nil
}
