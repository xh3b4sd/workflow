package dependabot

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/pkg/file"
	"github.com/xh3b4sd/workflow/pkg/runtime"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Branch    string
	Combine   []string
	Reviewers []string
	Version   struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Branch, "branch", "b", "main", "Dependabort target branch to merge pull requests into.")
	cmd.Flags().StringSliceVarP(&f.Combine, "combine", "c", []string{}, "Combine dependency updates of a kind, e.g. gomod:github.com/aws/aws-sdk-go-v2.")
	cmd.Flags().StringSliceVarP(&f.Reviewers, "reviewers", "r", []string{}, "Reviewers assigned to dependabot PRs, e.g. @xh3b4sd.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Branch == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-b/--branch must not be empty"})
		}
	}

	if !file.Exists(".github/CODEOWNERS") {
		if len(f.Reviewers) == 0 {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-r/--reviewers must not be empty"})
		}
		for _, x := range f.Reviewers {
			if !strings.HasPrefix(x, "@") {
				return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-r/--reviewers must start with @"})
			}
		}
	}

	{
		for _, x := range f.Combine {
			if !strings.HasPrefix(x, "gomod:") {
				return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-c/--combine must start with gomod:"})
			}
		}
	}

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
