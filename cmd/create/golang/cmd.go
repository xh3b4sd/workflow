package golang

import (
	"github.com/spf13/cobra"
)

const (
	use = "golang"
	sho = "Create a golang workflow for e.g. running tests and checking formatting."
	lon = "Create a golang workflow for e.g. running tests and checking formatting."
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
