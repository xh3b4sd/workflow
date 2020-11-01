package grpcts

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/generator"
	"github.com/xh3b4sd/workflow/pkg/generator/grpcts"
	"github.com/xh3b4sd/workflow/pkg/repo"
)

const (
	path = ".github/workflows/grpc-ts.yaml"
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

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var g generator.Interface
	{
		c := grpcts.Config{
			FilePath:           path,
			GithubCurrent:      repo.Current(),
			GithubOrganization: r.flag.Github.Organization,
			GithubRepository:   r.flag.Github.Repository,
			VersionGolang:      r.flag.Version.Golang,
			VersionGrpcWeb:     r.flag.Version.GrpcWeb,
			VersionProtoc:      r.flag.Version.Protoc,
		}

		g, err = grpcts.New(c)
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

		err = ioutil.WriteFile(path, b, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		b, err = g.Usage()
		if err != nil {
			return tracer.Mask(err)
		}

		fmt.Printf("%s", b)
	}

	return nil
}
