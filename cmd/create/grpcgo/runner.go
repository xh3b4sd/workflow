package grpcgo

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/generator"
	"github.com/xh3b4sd/workflow/pkg/generator/grpcgo"
	"github.com/xh3b4sd/workflow/pkg/repo"
)

const (
	path = ".github/workflows/grpc-go.yaml"
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

	// We compute the command actually being used to create the gRPC workflow
	// for golang code generation. The command is written to the header of the
	// workflow file. This enables us to re-generated workflows using the update
	// command without detailed knowledge about the specific flags being used
	// initially. Since the grpcgo command is a little verbose to support the
	// user in the first place, we do not want to print further usage
	// information upon updating the workflow. Thus we add the silence flag if
	// it is not already set.
	//
	//     -s/--silent
	//
	var command string
	{
		args := os.Args

		if !r.flag.Silent {
			args = append(args, "-s")
		}

		command = strings.Join(args, " ")
	}

	var g generator.Interface
	{
		c := grpcgo.Config{
			Command:            command,
			FilePath:           path,
			GithubCurrent:      repo.Current(),
			GithubOrganization: r.flag.Github.Organization,
			GithubRepository:   r.flag.Github.Repository,
			VersionGolang:      r.flag.Version.Golang,
			VersionProtoc:      r.flag.Version.Protoc,
		}

		g, err = grpcgo.New(c)
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

	if !r.flag.Silent {
		b, err = g.Usage()
		if err != nil {
			return tracer.Mask(err)
		}

		fmt.Printf("%s", b)
	}

	return nil
}
