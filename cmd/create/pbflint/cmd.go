package pbflint

import (
	"github.com/spf13/cobra"
)

const (
	use = "pbflint"
	sho = "Create a protocol buffer workflow for schema validation."
	lon = "Create a protocol buffer workflow for schema validation."
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
