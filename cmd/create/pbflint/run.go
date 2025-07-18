package pbflint

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/pkg/generator"
	"github.com/xh3b4sd/workflow/pkg/generator/pbflint"
)

const (
	path = ".github/workflows/pbf-lint.yaml"
)

type run struct {
	flag *flag
}

func (r *run) run(_ *cobra.Command, _ []string) error {
	var err error

	{
		err := r.flag.Validate()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var g generator.Interface
	{
		g = pbflint.New(pbflint.Config{
			Command:       strings.Join(os.Args, " "),
			FilePath:      path,
			VersionGolang: r.flag.Version.Golang,
			VersionProtoc: r.flag.Version.Protoc,
		})
	}

	var b []byte
	{
		b, err = g.Workflow()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}

		err = os.WriteFile(path, b, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
