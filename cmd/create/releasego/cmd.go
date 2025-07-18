package releasego

import (
	"github.com/spf13/cobra"
)

const (
	use = "releasego"
	sho = "Create a golang workflow for e.g. uploading cross compiled release assets."
	lon = "Create a golang workflow for e.g. uploading cross compiled release assets."
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
