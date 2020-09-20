package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/cmd/completion"
	"github.com/xh3b4sd/workflow/cmd/generate"
	"github.com/xh3b4sd/workflow/cmd/version"
	"github.com/xh3b4sd/workflow/pkg/project"
)

var (
	name        = project.Name()
	description = project.Description()
)

type Config struct {
	Logger logger.Interface
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	r := &runner{
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	m := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
		// We slience errors because we do not want to see spf13/cobra printing.
		// The errors returned by the commands will be propagated to the main.go
		// anyway, where we have custom error printing for the command line
		// tool.
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	var err error

	var completionCmd *cobra.Command
	{
		c := completion.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		completionCmd, err = completion.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var generateCmd *cobra.Command
	{
		c := generate.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		generateCmd, err = generate.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var versionCmd *cobra.Command
	{
		c := version.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		versionCmd, err = version.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m.AddCommand(completionCmd)
	m.AddCommand(generateCmd)
	m.AddCommand(versionCmd)

	return m, nil
}
