package update

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/workflow/cmd/update/all"
)

const (
	use = "update"
	sho = "Update github workflows and config files."
	lon = "Update github workflows and config files."
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
		cmd.AddCommand(all.New())
	}

	return cmd
}
