package dependabot

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Reviewers []string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&f.Reviewers, "reviewers", "r", []string{}, "Reviewers assigned to dependabot PRs, e.g. xh3b4sd. Works with github usernames and teams.")
}

func (f *flag) Validate() error {
	if len(f.Reviewers) == 0 {
		return tracer.Maskf(invalidFlagError, "-r/--reviewers must not be empty")
	}

	return nil
}
