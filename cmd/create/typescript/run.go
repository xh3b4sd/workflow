package typescript

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type run struct {
	flag *flag
}

func (r *run) run(cmd *cobra.Command, args []string) error {
	{
		err := r.flag.Validate()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/typescript.yaml"

		t, err := template.New(p).Parse(templateTypescript)
		if err != nil {
			return tracer.Mask(err)
		}

		var b bytes.Buffer
		err = t.ExecuteTemplate(&b, p, r.data())
		if err != nil {
			return tracer.Mask(err)
		}

		err = os.WriteFile(p, b.Bytes(), 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) data() any {
	type Data struct {
		Command string
		Version version.Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Version: version.Version{
			Checkout:  version.Checkout,
			Node:      r.flag.Version.Node,
			SetupNode: version.SetupNode,
		},
	}
}
