package generate

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/cmd/generate/dependabot"
	"github.com/xh3b4sd/workflow/cmd/generate/dockergo"
	"github.com/xh3b4sd/workflow/cmd/generate/dockerts"
	"github.com/xh3b4sd/workflow/cmd/generate/golang"
	"github.com/xh3b4sd/workflow/cmd/generate/grpcgo"
	"github.com/xh3b4sd/workflow/cmd/generate/grpcts"
)

const (
	name        = "generate"
	description = "Generate github workflows and config files."
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var dependabotCmd *cobra.Command
	{
		c := dependabot.Config{
			Logger: config.Logger,
		}

		dependabotCmd, err = dependabot.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var dockergoCmd *cobra.Command
	{
		c := dockergo.Config{
			Logger: config.Logger,
		}

		dockergoCmd, err = dockergo.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var dockertsCmd *cobra.Command
	{
		c := dockerts.Config{
			Logger: config.Logger,
		}

		dockertsCmd, err = dockerts.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var golangCmd *cobra.Command
	{
		c := golang.Config{
			Logger: config.Logger,
		}

		golangCmd, err = golang.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var grpcgoCmd *cobra.Command
	{
		c := grpcgo.Config{
			Logger: config.Logger,
		}

		grpcgoCmd, err = grpcgo.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var grpctsCmd *cobra.Command
	{
		c := grpcts.Config{
			Logger: config.Logger,
		}

		grpctsCmd, err = grpcts.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var c *cobra.Command
	{
		r := &runner{
			logger: config.Logger,
		}

		c = &cobra.Command{
			Use:   name,
			Short: description,
			Long:  description,
			RunE:  r.Run,
		}

		c.AddCommand(dependabotCmd)
		c.AddCommand(dockergoCmd)
		c.AddCommand(dockertsCmd)
		c.AddCommand(golangCmd)
		c.AddCommand(grpcgoCmd)
		c.AddCommand(grpctsCmd)
	}

	return c, nil
}
