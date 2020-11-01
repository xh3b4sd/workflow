package grpcts

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "grpcts"
	short = "Create a grpc workflow for typescript code generation."
	long  = `Create a grpc workflow for typescript code generation. The workflow
generated here works in a setup of two Github repositories. Call them
apischema and tscode. The workflow generated with the following command is
added to the apischema repository.

    workflow create grpcts -o xh3b4sd -r tscode

In order to make the workflow function correctly a deploy key is generated
and distributed as follows. The public key and the encrypted private key
files are added to the apischema repository. The public key is added as
deploy key with write access to the tscode repository. The GPG password and
the deploy keys are generated with the red command line tool. For more
information see https://github.com/xh3b4sd/red.

    red generate keys -d .github/asset/xh3b4sd/tscode

Generating the deploy keys also generates a GPG password which is used to
decrypt the encrypted private key within the build container of the workflow.
Considering the example described above, the GPG password needs to be added
to the apischema Github repository secrets using the following secret name.

    RED_GPG_PASS_XH3B4SD_TSCODE

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
