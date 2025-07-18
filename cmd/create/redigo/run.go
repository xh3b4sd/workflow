package redigo

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
		p := ".github/workflows/go-redis.yaml"

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
	type Data struct {
		Command string
		Version version.Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Version: version.Version{
			Checkout: version.Checkout,
			Golang:   desiredVersion(r.flag.Version.Golang),
		},
	}
}

func desiredVersion(v string) string {
	s := strings.Split(v, ".")
	if len(s) != 3 {
		// The given version must be something like 1.15.2 which implies 3 parts
		// when splitting at the period.
		panic("must have 3 parts")
	}

	return s[0] + "." + s[1]
}
