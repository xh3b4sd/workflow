package create

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/cmd/create/cfmtest"
	"github.com/xh3b4sd/workflow/cmd/create/dependabot"
	"github.com/xh3b4sd/workflow/cmd/create/dockergo"
	"github.com/xh3b4sd/workflow/cmd/create/dockerts"
	"github.com/xh3b4sd/workflow/cmd/create/dsmupdate"
	"github.com/xh3b4sd/workflow/cmd/create/dsmverify"
	"github.com/xh3b4sd/workflow/cmd/create/golang"
	"github.com/xh3b4sd/workflow/cmd/create/grpcgo"
	"github.com/xh3b4sd/workflow/cmd/create/grpcts"
	"github.com/xh3b4sd/workflow/cmd/create/npm"
	"github.com/xh3b4sd/workflow/cmd/create/redigo"
	"github.com/xh3b4sd/workflow/cmd/create/typescript"
)

const (
	name  = "create"
	short = "Create github workflows and config files."
	long  = "Create github workflows and config files."
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var cfmTestCmd *cobra.Command
	{
		c := cfmtest.Config{
			Logger: config.Logger,
		}

		cfmTestCmd, err = cfmtest.New(c)
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

	var dsmUpdateCmd *cobra.Command
	{
		c := dsmupdate.Config{
			Logger: config.Logger,
		}

		dsmUpdateCmd, err = dsmupdate.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var dsmVerifyCmd *cobra.Command
	{
		c := dsmverify.Config{
			Logger: config.Logger,
		}

		dsmVerifyCmd, err = dsmverify.New(c)
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

	var npmCmd *cobra.Command
	{
		c := npm.Config{
			Logger: config.Logger,
		}

		npmCmd, err = npm.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var redigoCmd *cobra.Command
	{
		c := redigo.Config{
			Logger: config.Logger,
		}

		redigoCmd, err = redigo.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var typescriptCmd *cobra.Command
	{
		c := typescript.Config{
			Logger: config.Logger,
		}

		typescriptCmd, err = typescript.New(c)
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
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}

		c.AddCommand(cfmTestCmd)
		c.AddCommand(dependabotCmd)
		c.AddCommand(dockergoCmd)
		c.AddCommand(dockertsCmd)
		c.AddCommand(dsmUpdateCmd)
		c.AddCommand(dsmVerifyCmd)
		c.AddCommand(golangCmd)
		c.AddCommand(grpcgoCmd)
		c.AddCommand(grpctsCmd)
		c.AddCommand(npmCmd)
		c.AddCommand(redigoCmd)
		c.AddCommand(typescriptCmd)
	}

	return c, nil
}
