package pbfgo

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Github struct {
		Organization string
		Repository   string
	}
	Version struct {
		Golang string
		Protoc string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Github.Organization, "github-organization", "o", "", "Github organization to generate code for.")
	cmd.Flags().StringVarP(&f.Github.Repository, "github-repository", "r", "", "Github repository to generate code for.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
	cmd.Flags().StringVarP(&f.Version.Protoc, "version-protoc", "p", version.Protoc, "Protoc version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Github.Organization == "" {
			return tracer.Maskf(invalidFlagError, "-o/--github-organization must not be empty")
		}
		if f.Github.Repository == "" {
			return tracer.Maskf(invalidFlagError, "-r/--github-repository must not be empty")
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

	{
		if f.Version.Protoc == "" {
			return tracer.Maskf(invalidFlagError, "-p/--version-protoc must not be empty")
		}
	}

	return nil
}
