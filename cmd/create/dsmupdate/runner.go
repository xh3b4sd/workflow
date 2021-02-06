package dsmupdate

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type runner struct {
	flag   *flag
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) data() interface{} {
	type Repository struct {
		Name string
	}

	type Version struct {
		Golang string
	}

	type Data struct {
		Command    string
		Repository Repository
		Version    Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Repository: Repository{
			Name: r.flag.Repository.Name,
		},
		Version: Version{
			Golang: r.flag.Version.Golang,
		},
	}
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	{
		p := ".github/workflows/"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/dsm-update.yaml"

		t, err := template.New(p).Parse(templateWorkflow)
		if err != nil {
			return tracer.Mask(err)
		}

		var b bytes.Buffer
		err = t.ExecuteTemplate(&b, p, r.data())
		if err != nil {
			return tracer.Mask(err)
		}

		err = ioutil.WriteFile(p, b.Bytes(), 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
