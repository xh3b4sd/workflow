package releasego

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

func (r *run) run(_ *cobra.Command, _ []string) error {
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
		p := ".github/workflows/go-release.yaml"

		t, err := template.New(p).Parse(templateWorkflow)
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

func (r *run) data() interface{} {
	type Release struct {
		Assets map[string]string
	}

	type Repository struct {
		Name string
		Path string
	}

	type Variable struct {
		GitSha string
		GitTag string
	}

	type Data struct {
		Command    string
		Release    Release
		Repository Repository
		Variable   Variable
		Version    version.Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Release: Release{
			Assets: assets(r.flag.Release.Assets),
		},
		Repository: Repository{
			Name: r.flag.Repository.Name,
			Path: r.flag.Repository.Path,
		},
		Variable: Variable{
			GitSha: r.flag.Variable.GitSha,
			GitTag: r.flag.Variable.GitTag,
		},
		Version: version.Version{
			Checkout: version.Checkout,
			Golang:   r.flag.Version.Golang,
			SetupGo:  version.SetupGo,
		},
	}
}

func assets(str string) map[string]string {
	ass := map[string]string{}

	for _, s := range strings.Split(str, ",") {
		spl := strings.Split(s, "/")
		ass[spl[0]] = spl[1]
	}

	return ass
}
