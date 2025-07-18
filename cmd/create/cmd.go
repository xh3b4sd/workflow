package create

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/workflow/cmd/create/dependabot"
	"github.com/xh3b4sd/workflow/cmd/create/dockergo"
	"github.com/xh3b4sd/workflow/cmd/create/dockerts"
	"github.com/xh3b4sd/workflow/cmd/create/golang"
	"github.com/xh3b4sd/workflow/cmd/create/npm"
	"github.com/xh3b4sd/workflow/cmd/create/pbfgo"
	"github.com/xh3b4sd/workflow/cmd/create/pbflint"
	"github.com/xh3b4sd/workflow/cmd/create/pbfts"
	"github.com/xh3b4sd/workflow/cmd/create/redigo"
	"github.com/xh3b4sd/workflow/cmd/create/redis"
	"github.com/xh3b4sd/workflow/cmd/create/releasego"
	"github.com/xh3b4sd/workflow/cmd/create/typescript"
)

const (
	use = "create"
	sho = "Create github workflows and config files."
	lon = "Create github workflows and config files."
)

func New() *cobra.Command {
	var cmd *cobra.Command
	{
		cmd = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		cmd.AddCommand(dependabot.New())
		cmd.AddCommand(dockergo.New())
		cmd.AddCommand(dockerts.New())
		cmd.AddCommand(golang.New())
		cmd.AddCommand(npm.New())
		cmd.AddCommand(pbfgo.New())
		cmd.AddCommand(pbflint.New())
		cmd.AddCommand(pbfts.New())
		cmd.AddCommand(redigo.New())
		cmd.AddCommand(redis.New())
		cmd.AddCommand(releasego.New())
		cmd.AddCommand(typescript.New())
	}

	return cmd
}
