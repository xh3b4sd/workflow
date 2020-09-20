package version

import (
	"context"
	"fmt"
	"io"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/project"
)

type runner struct {
	logger logger.Interface
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Fprintf(r.stdout, "Git Commit    %s\n", project.GitSHA())
	fmt.Fprintf(r.stdout, "Go Version    %s\n", runtime.Version())
	fmt.Fprintf(r.stdout, "Go Arch       %s\n", runtime.GOARCH)
	fmt.Fprintf(r.stdout, "Go OS         %s\n", runtime.GOOS)
	fmt.Fprintf(r.stdout, "Source        %s\n", project.Source())
	fmt.Fprintf(r.stdout, "Version       %s\n", project.Version())

	return nil
}
