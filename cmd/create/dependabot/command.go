package dependabot

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "dependabot"
	short = "Create a dependabot workflow for e.g. golang and docker."
	long  = `Create a dependabot workflow for e.g. golang and docker. The dependabot
workflow for golang includes a separate action of executing "go mod tidy" due
to some dependabot limitations. This limitation requires automated fixing and
therefore dedicated push access to the configured repository. The explicitly
authorized push access is necessary in order to trigger another execution of
the "go build" workflow after we fixed the go.mod and go.sum files. Due to
these implementation details we setup deploy keys in each repository via the
"red" command line tool. More information about the tool can be found at
github.com/xh3b4sd/red. For each repository using the dependabot workflow the
following command must be used in order to generate deploy keys and the
associated GPG password.

    red generate keys -d .github/asset

The GPG encrypted private key must be put into the ".github/asset" directory.
During each build the GPG encrypted private key is decrypted within the build
container and used to setup the local SSH agent for the explicit push access.

    .github/asset/id_rsa.enc

The plain text public key must be added as deploy key with write access to
the configured repository. During builds of dependabot pull requests the
configured deploy key verifies that the configured private key is allowed to
push changes during builds.

    .github/asset/id_rsa.pub

During builds, a password is required for the decryption of the GPG encrypted
private key. This password gets generated together with the RSA public and
private key as shown above. The password must be set as secret to the
configured repository. The secret name must be as follows.

    RED_GPG_PASS
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
