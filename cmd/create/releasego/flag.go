package releasego

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/runtime"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	Linker struct {
		Git struct {
			Sha string
			Tag string
		}
		Path string
	}
	Release struct {
		Assets string
	}
	Version struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Linker.Git.Sha, "linker-git-sha", "s", "sha", "The linker variable receiving the binary's git sha.")
	cmd.Flags().StringVarP(&f.Linker.Git.Tag, "linker-git-tag", "t", "tag", "The linker variable receiving the binary's git tag.")
	cmd.Flags().StringVarP(&f.Linker.Path, "linker-path", "p", "pkg/runtime", "The repository path to compile linker flags into.")
	cmd.Flags().StringVarP(&f.Release.Assets, "release-assets", "a", "darwin/amd64,linux/amd64", "Binary architectures to compile.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.Linker.Git.Sha == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-s/--linker-git-sha must not be empty"})
		}
		if f.Linker.Git.Tag == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-t/--linker-git-tag must not be empty"})
		}
		if f.Linker.Path == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-p/--linker-path must not be empty"})
		}
	}

	{
		if f.Release.Assets == "" {
			return tracer.Mask(runtime.InvalidFlagError, tracer.Context{Key: "reason", Value: "-a/--release-assets must not be empty"})
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
