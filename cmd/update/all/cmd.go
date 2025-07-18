package all

import (
	"github.com/spf13/cobra"
)

const (
	use = "all"
	sho = "Update all github workflows to the latest version."
	lon = `Update all github workflows to the latest version. When creating a new
workflow file the original command instruction in form of os.Args is written
to the header of the workflow file. A typical workflow file header looks like
the following.

    #
    # Do not edit. This file was generated via the "workflow" command line tool.
    # More information about the tool can be found at github.com/xh3b4sd/workflow.
    #
    #     workflow create dependabot -r @xh3b4sd
    #

This information of the executable command is used to make workflow updates
reproducible. All workflow files within the github specific workflow
directory are inspected when collecting command instructions. Once all
commands are known they are executed dynamically while new behaviour is
applied.

    .github/workflows/

`
)

func New() *cobra.Command {
	var cmd *cobra.Command
	{
		cmd = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{}).run,
		}
	}

	return cmd
}
