package grpc

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

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
	type Github struct {
		Organization string
		Repository   string
	}

	type Version struct {
		Golang string
		Protoc string
	}

	type Data struct {
		Github  Github
		Version Version
	}

	return Data{
		Github: Github{
			Organization: r.flag.Github.Organization,
			Repository:   r.flag.Github.Repository,
		},
		Version: Version{
			Golang: r.flag.Version.Golang,
			Protoc: r.flag.Version.Protoc,
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
		p := ".github/workflows/grpc-go.yaml"

		f := template.FuncMap{
			"ToUpper": strings.ToUpper,
		}

		t, err := template.New(p).Funcs(f).Parse(templateGolang)
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
