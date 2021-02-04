package cfmtest

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
)

var (
	repositories = []string{
		"apiserver",
		"apiworker",
		"cfm",
		"flux",
	}
)

type flag struct {
	Repository struct {
		Name string
	}
	Version struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Repository.Name, "repository-name", "r", "", "Repository name to generate the workflow for, e.g. flux or apiworker.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Repository.Name == "" {
			return tracer.Maskf(invalidFlagError, "-r/--repository-name must not be empty")
		}
		if !contains(repositories, f.Repository.Name) {
			return tracer.Maskf(invalidFlagError, "-r/--repository-name must not be one of %s", strings.Join(repositories, ", "))
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

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}

	return false
}
