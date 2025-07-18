package completion

import (
	"github.com/spf13/cobra"
)

const (
	use = "completion"
	sho = "Generate shell completions."
	lon = `Supported positional arguments and respective shell completions are
as follows.

    bash
    fish
    powershell
    zsh

Generating zsh completion for Oh My Zsh can be done by writing the
generated completion to the custom plugin folder.

    mkdir -p ~/.oh-my-zsh/custom/plugins/workflow && workflow completion zsh > ~/.oh-my-zsh/custom/plugins/workflow/_workflow

	`
)

func New() *cobra.Command {
	var cmd *cobra.Command
	{
		cmd = &cobra.Command{
			Use:                   use,
			Short:                 sho,
			Long:                  lon,
			DisableFlagsInUseLine: true,
			ValidArgs:             []string{"bash", "fish", "powershell", "zsh"},
			Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
			RunE:                  (&run{}).run,
		}
	}

	return cmd
}
