package command

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"github.com/xh3b4sd/tracer"
)

var (
	// header is the regular expression used to parse the common workflow header
	// that tracks the command used for the workflow creation.
	header = regexp.MustCompile(`(?m)^[\s]*#[\s]+(workflow create .*)$`)
)

type Config struct {
	FileSystem   afero.Fs
	WorkflowPath string
}

type Command struct {
	fileSystem   afero.Fs
	workflowPath string
}

func New(c Config) *Command {
	if c.FileSystem == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.FileSystem must not be empty", c)))
	}
	if c.WorkflowPath == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.WorkflowPath must not be empty", c)))
	}

	return &Command{
		fileSystem:   c.FileSystem,
		workflowPath: c.WorkflowPath,
	}
}

func (c *Command) Parse() ([][]string, error) {
	files, err := c.files(".yaml")
	if err != nil {
		return nil, tracer.Mask(err)
	}

	commands, err := c.commands(files...)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var args [][]string
	for _, c := range commands {
		args = append(args, strings.Split(c, " "))
	}

	return args, nil
}

func (c *Command) commands(files ...string) ([]string, error) {
	var commands []string

	for _, f := range files {
		b, err := afero.ReadFile(c.fileSystem, f)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		matches := header.FindSubmatch(b)

		if len(matches) == 0 {
			continue
		}

		commands = append(commands, string(matches[1]))
	}

	return commands, nil
}

func (c *Command) files(exts ...string) ([]string, error) {
	var files []string
	{
		walkFunc := func(p string, i os.FileInfo, err error) error {
			if err != nil {
				return tracer.Mask(err)
			}

			// We do not want to track files with the wrong extension. We are
			// interested in workflow files having the ".yaml" extension.
			for _, e := range exts {
				if filepath.Ext(i.Name()) != e {
					return nil
				}
			}

			files = append(files, p)

			return nil
		}

		err := afero.Walk(c.fileSystem, c.workflowPath, walkFunc)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return files, nil
}
