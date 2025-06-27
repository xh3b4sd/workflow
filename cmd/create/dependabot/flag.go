package dependabot

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Branch    string
	Reviewers []string
	Version   struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Branch, "branch", "b", "main", "Dependabort target branch to merge pull requests into.")
	cmd.Flags().StringSliceVarP(&f.Reviewers, "reviewers", "r", []string{}, "Reviewers assigned to dependabot PRs, e.g. @xh3b4sd.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Branch != "main" && f.Branch != "master" {
			return tracer.Maskf(invalidFlagError, "-b/--branch must either be main or master")
		}
	}

	{
		if len(f.Reviewers) == 0 {
			return tracer.Maskf(invalidFlagError, "-r/--reviewers must not be empty")
		}
		for _, r := range f.Reviewers {
			if !strings.HasPrefix(r, "@") {
				return tracer.Maskf(invalidFlagError, "-r/--reviewers must start with @")
			}
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
