package pbfgo

import (
	"github.com/spf13/cobra"
)

const (
	use = "pbfgo"
	sho = "Create a protocol buffer workflow for golang code generation."
	lon = `Create a protocol buffer workflow for golang code generation. The workflow
generated here works in a setup of two Github repositories. Call them apischema
and apigocode. The workflow generated with the following command is added to the
apischema repository.

    workflow create pbfgo -o fancy-organization -r apigocode

In order to make the workflow function correctly a deploy key must be generated
and distributed. The public key is added as deploy key with write access to the
apigocode repository. The private key is added as Github Action Secret to the
apischema repository, call it SSH_DEPLOY_KEY_APIGOCODE. A new key pair can be
generated like shown below.

    ssh-keygen -t ed25519 -C your@email.com

More information about the Github Action used to push changes from one
repository to another can be found following the link below.

    https://github.com/cpina/github-action-push-to-another-repository
`
)

func New() *cobra.Command {
	var flg *flag
	{
		flg = &flag{}
	}

	var cmd *cobra.Command
	{
		cmd = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{flag: flg}).run,
		}
	}

	{
		flg.Init(cmd)
	}

	return cmd
}
