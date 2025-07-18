package dependabot

import (
	"github.com/spf13/cobra"
)

const (
	use = "dependabot"
	sho = "Create a dependabot workflow for e.g. golang and docker."
	lon = "Create a dependabot workflow for e.g. golang and docker."
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
