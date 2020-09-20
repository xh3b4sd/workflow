package completion

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name             = "completion"
	shortDescription = "generate shell completions"
	longDescription  = `Generating zsh completion for Oh My Zsh can be done by writing the
generated completion to the custom plugin folder.

    mkdir -p ~/.oh-my-zsh/custom/plugins/workflow && workflow completion zsh > ~/.oh-my-zsh/custom/plugins/workflow/_workflow

	`
)

type Config struct {
	Logger logger.Interface
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	r := &runner{
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:                   name,
		Short:                 shortDescription,
		Long:                  longDescription,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE:                  r.Run,
	}

	return c, nil
}
