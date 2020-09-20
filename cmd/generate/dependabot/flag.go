package dependabot

import (
	"github.com/spf13/cobra"
)

type flag struct {
	Reviewers []string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&f.Reviewers, "reviewers", "r", []string{}, "Reviewers assigned to dependabot PRs, e.g. xh3b4sd. Works with github usernames and teams.")
}

func (f *flag) Validate() error {
	return nil
}
