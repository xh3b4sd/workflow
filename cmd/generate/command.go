package generate

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/cmd/generate/credentials"
	"github.com/xh3b4sd/workflow/cmd/generate/dependabot"
	"github.com/xh3b4sd/workflow/cmd/generate/golang"
)

const (
	name        = "generate"
	description = "Generate workflows and required assets."
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var credentialsCmd *cobra.Command
	{
		c := credentials.Config{
			Logger: config.Logger,
		}

		credentialsCmd, err = credentials.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

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

		c.AddCommand(credentialsCmd)
		c.AddCommand(dependabotCmd)
		c.AddCommand(golangCmd)
	}

	return c, nil
}
