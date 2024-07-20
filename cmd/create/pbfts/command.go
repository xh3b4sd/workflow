package pbfts

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "pbfts"
	short = "Create a protocol buffer workflow for typescript code generation."
	long  = `Create a protocol buffer workflow for typescript code generation. The workflow
generated here works in a setup of two Github repositories. Call them apischema
and apitscode. The workflow generated with the following command is added to the
apischema repository.

    workflow create pbfts -o fancy-organization -r apitscode

In order to make the workflow function correctly a deploy key must be generated
and distributed. The public key is added as deploy key with write access to the
apitscode repository. The private key is added as Github Action Secret to the
apischema repository, call it SSH_DEPLOY_KEY_APIGOCODE. A new key pair can be
generated like shown below.

    ssh-keygen -t ed25519 -C your@email.com

More information about the Github Action used to push changes from one
repository to another can be found following the link below.

    https://github.com/cpina/github-action-push-to-another-repository
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
