package dockergo

import (
	"github.com/spf13/cobra"
)

const (
	use = "dockergo"
	sho = "Create a docker workflow for building and pushing docker images of golang apps."
	lon = "Create a docker workflow for building and pushing docker images of golang apps."
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
