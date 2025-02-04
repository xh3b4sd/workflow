package all

import (
	"context"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/parser"
	"github.com/xh3b4sd/workflow/pkg/parser/command"
)

type runner struct {
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var p parser.Interface
	{
		c := command.Config{
			FileSystem: afero.NewOsFs(),

			WorkflowPath: ".github/",
		}

		p, err = command.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var l [][]string
	{
		l, err = p.Parse()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// We are going to mess with the command line args of the program's process
	// during the update command execution. In order to be "good" citizens we
	// set the os.Args back to the value they had before once we are done with
	// the execution of other commands.
	{
		args := os.Args
		defer func() {
			os.Args = args
		}()
	}

	for _, a := range l {
		c := cmd.Root()

		c, err = commandFor("create", c.Commands())
		if err != nil {
			return tracer.Mask(err)
		}

		// At this point we have the dynamic command. The implementation of the
		// parser interface guarantees that the third element in the argument
		// list is the workflow command we can execute dynamically.
		//
		//     workflow create dependabot
		//     workflow create golang
		//
		c, err = commandFor(a[2], c.Commands())
		if err != nil {
			return tracer.Mask(err)
		}

		// We need to overwrite the os.Args because the commands we are going to
		// execute dynamically use os.Args to write the proper command of the
		// workflow creation to the header of the workflow file. Only due to
		// this mechanism we enable reproducible workflow updates.
		os.Args = a

		// The command we execute dynamically needs to parse its own flags in
		// order to function correctly.
		err = c.Flags().Parse(a[1:])
		if err != nil {
			return tracer.Mask(err)
		}

		err = c.RunE(c, nil)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func commandFor(action string, cmds []*cobra.Command) (*cobra.Command, error) {
	for _, c := range cmds {
		if c.Name() == action {
			return c, nil
		}
	}

	return nil, tracer.Maskf(commandNotFoundError, "%s", action)
}
