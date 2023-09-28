package all

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "all"
	short = "Update all github workflows to the latest version."
	long  = `Update all github workflows to the latest version. When creating a new
workflow file the original command instruction in form of os.Args is written
to the header of the workflow file. A typical workflow file header looks like
the following.

    #
    # Do not edit. This file was generated via the "workflow" command line tool.
    # More information about the tool can be found at github.com/xh3b4sd/workflow.
    #
    #     workflow create dependabot -r xh3b4sd
    #

This information of the executable command is used to make workflow updates
reproducible. All workflow files within the github specific workflow
directory are inspected when collecting command instructions. Once all
commands are known they are executed dynamically while new behaviour is
applied.

    .github/workflows/

`
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var c *cobra.Command
	{
		r := &runner{
			logger: config.Logger,
		}

		c = &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}
	}

	return c, nil
}
