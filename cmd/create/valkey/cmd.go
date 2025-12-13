package valkey

import (
	"github.com/spf13/cobra"
)

const (
	use = "valkey"
	sho = "Create a valkey workflow for e.g. running integration tests."
	lon = "Create a valkey workflow for e.g. running integration tests."
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
