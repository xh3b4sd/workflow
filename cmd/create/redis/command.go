package redis

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "redis"
	short = "Create a redis workflow for e.g. running conformance tests."
	long  = "Create a redis workflow for e.g. running conformance tests."
)

type Config struct {
	Logger logger.Interface
}

func New(con Config) (*cobra.Command, error) {
	if con.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", con)
	}

	var c *cobra.Command
	{
		f := &flag{}

		r := &runner{
			fla: f,
			log: con.Logger,
		}

		c = &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}

		f.Init(c)
	}

	return c, nil
}
