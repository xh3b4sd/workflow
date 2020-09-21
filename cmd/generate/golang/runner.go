package golang

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type runner struct {
	logger logger.Interface
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
	{
		p := ".github/workflows/"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/go-build.yaml"

		err := ioutil.WriteFile(p, []byte(templateGolang), 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
