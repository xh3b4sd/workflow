package typescript

import (
	"github.com/spf13/cobra"
)

const (
	use = "typescript"
	sho = "Create a typescript workflow for e.g. building and formatting typescript code."
	lon = "Create a typescript workflow for e.g. building and formatting typescript code."
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
