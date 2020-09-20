package dependabot

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/file"
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
	var ecosystems []string
	{
		if file.Exists("Dockerfile") {
			ecosystems = append(ecosystems, "docker")
		}

		{
			ecosystems = append(ecosystems, "github-actions")
		}

		if file.Exists("go.mod") && file.Exists("go.sum") {
			ecosystems = append(ecosystems, "go")
		}
	}

	return nil
}
