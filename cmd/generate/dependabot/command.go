package dependabot

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "dependabot"
	short = "Generate a dependabot workflow for e.g. golang and docker."
	long  = `Generate a dependabot workflow for e.g. golang and docker. The dependabot
workflow for golang includes a separate action of executing "go mod tidy" due
to some dependabot limitations. This limitation requires automated fixing and
therefore dedicated push access. The explicitly authorized push access is
necessary in order to trigger another execution of the "go build" workflow
after we fixed the go.mod and go.sum files. Due to these implementation
details we setup deploy keys in each repository via the "red" command line
tool. More information about the tool can be found at github.com/xh3b4sd/red.
For each repository using the dependabot workflow the following command must
be used in order to generate deploy keys and the associated GPG password.

    red generate keys -d .github/asset

The GPG encrypted private key must be put into the ".github/asset" directory.
During each build the GPG encrypted private key is decrypted within the build
container and used to setup the local SSH agent for the explicit push access.

For the decryption of the GPG encrypted private key a password is required.
This password gets generated together with the public and private key. The
password must be made available to each repository as secret using the name
"RED_GPG_PASS".
	`
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var c *cobra.Command
	{
		f := &flag{}

		r := &runner{
			flag:   f,
			logger: config.Logger,
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
