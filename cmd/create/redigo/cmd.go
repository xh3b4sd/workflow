package redigo

import (
	"github.com/spf13/cobra"
)

const (
	use = "redigo"
	sho = "Create a redis workflow for e.g. running conformance tests."
	lon = "Create a redis workflow for e.g. running conformance tests."
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
