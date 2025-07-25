package dependabot

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/pkg/file"
)

type run struct {
	flag *flag
}

func (r *run) run(_ *cobra.Command, _ []string) error {
	{
		err := r.flag.Validate()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/CODEOWNERS"

		if !file.Exists(p) {
			var b bytes.Buffer
			{
				b.WriteString("*")
				for _, r := range r.flag.Reviewers {
					b.WriteString(fmt.Sprintf(" %s", r))
				}
				b.WriteString("\n")
			}

			err := os.WriteFile(p, b.Bytes(), 0600)
			if err != nil {
				return tracer.Mask(err)
			}
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
		p := ".github/dependabot.yaml"

		t, err := template.New(p).Parse(templateWorkflow)
		if err != nil {
			return tracer.Mask(err)
		}

		var b bytes.Buffer
		err = t.ExecuteTemplate(&b, p, r.dependabotData())
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

func (r *run) dependabotData() any {
	type Ecosystem struct {
		Branch    string
		Name      string
		Reviewers []string
	}

	type Data struct {
		Command    string
		Ecosystems []Ecosystem
	}

	var ecosystems []Ecosystem
	{
		if file.Exists("Dockerfile") {
			ecosystems = append(ecosystems, Ecosystem{
				Branch:    r.flag.Branch,
				Name:      "docker",
				Reviewers: r.flag.Reviewers,
			})
		}

		{
			ecosystems = append(ecosystems, Ecosystem{
				Branch:    r.flag.Branch,
				Name:      "github-actions",
				Reviewers: r.flag.Reviewers,
			})
		}

		if file.Exists("go.mod") && file.Exists("go.sum") {
			ecosystems = append(ecosystems, Ecosystem{
				Branch:    r.flag.Branch,
				Name:      "gomod",
				Reviewers: r.flag.Reviewers,
			})
		}
	}

	return Data{
		Command:    strings.Join(os.Args, " "),
		Ecosystems: ecosystems,
	}
}
