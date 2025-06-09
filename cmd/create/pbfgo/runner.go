package pbfgo

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/generator"
	"github.com/xh3b4sd/workflow/pkg/generator/pbfgo"
)

const (
	path = ".github/workflows/pbf-go.yaml"
)

type runner struct {
	flag   *flag
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(_ context.Context, _ *cobra.Command, _ []string) error {
	var err error

	var g generator.Interface
	{
		c := pbfgo.Config{
			Command:            strings.Join(os.Args, " "),
			FilePath:           path,
			GithubOrganization: r.flag.Github.Organization,
			GithubRepository:   r.flag.Github.Repository,
			VersionGolang:      r.flag.Version.Golang,
			VersionProtoc:      r.flag.Version.Protoc,
		}

		g, err = pbfgo.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var b []byte
	{
		b, err = g.Workflow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}

		err = os.WriteFile(path, b, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
