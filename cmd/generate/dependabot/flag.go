package dependabot

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/golang"
)

type flag struct {
	Reviewers []string
	Version   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&f.Reviewers, "reviewers", "r", []string{}, "Reviewers assigned to dependabot PRs, e.g. xh3b4sd. Works with github usernames and teams.")
	cmd.Flags().StringVarP(&f.Version, "version", "v", golang.Version, "Golang version to set in, e.g. go.mod.")
}

func (f *flag) Validate() error {
	if len(f.Reviewers) == 0 {
		return tracer.Maskf(invalidFlagError, "-r/--reviewers must not be empty")
	}

	if f.Version == "" {
		return tracer.Maskf(invalidFlagError, "-v/--version must not be empty")
	}

	s := strings.Split(f.Version, ".")
	if len(s) != 3 {
		return tracer.Maskf(invalidFlagError, "-v/--version must have 3 parts like 1.15.2")
	}

	return nil
}
