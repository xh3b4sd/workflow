package dockerts

import (
	"github.com/spf13/cobra"
)

const (
	use = "dockerts"
	sho = "Create a docker workflow for building and pushing docker images of typescript apps."
	lon = "Create a docker workflow for building and pushing docker images of typescript apps."
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
