package dependabot

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/file"
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

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	{
		p := ".github/workflows/"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/go-mod-tidy.yaml"

		err := ioutil.WriteFile(p, []byte(templateGoModTidy), 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/dependabot.yaml"

		t, err := template.New(p).Parse(templateDependabot)
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

func (r *runner) data() interface{} {
	type Ecosystem struct {
		Name      string
		Reviewers []string
	}

	var data []Ecosystem
	{
		if file.Exists("Dockerfile") {
			data = append(data, Ecosystem{
				Name:      "docker",
				Reviewers: r.flag.Reviewers,
			})
		}

		{
			data = append(data, Ecosystem{
				Name:      "github-actions",
				Reviewers: r.flag.Reviewers,
			})
		}

		if file.Exists("go.mod") && file.Exists("go.sum") {
			data = append(data, Ecosystem{
				Name:      "gomod",
				Reviewers: r.flag.Reviewers,
			})
		}
	}

	return data
}
