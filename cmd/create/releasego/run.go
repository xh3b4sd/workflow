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

func (r *run) data() any {
	type Git struct {
		Sha string
		Tag string
	}
	type Linker struct {
		Path string
		Git  Git
	}
	type Release struct {
		Assets map[string]string
	}

	type Data struct {
		Command string
		Linker  Linker
		Release Release
		Version version.Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Linker: Linker{
			Path: r.flag.Linker.Path,
			Git: Git{
				Sha: r.flag.Linker.Git.Sha,
				Tag: r.flag.Linker.Git.Tag,
			},
		},
		Release: Release{
			Assets: assets(r.flag.Release.Assets),
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

	for s := range strings.SplitSeq(str, ",") {
		spl := strings.Split(s, "/")
		ass[spl[0]] = spl[1]
	}

	return ass
}
