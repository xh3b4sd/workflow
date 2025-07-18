package npm

import (
	"github.com/spf13/cobra"
)

const (
	use = "npm"
	sho = "Create a npm workflow for e.g. building and publishing npm packages."
	lon = "Create a npm workflow for e.g. building and publishing npm packages."
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
