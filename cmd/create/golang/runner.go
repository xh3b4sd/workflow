package golang

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/env"
	"github.com/xh3b4sd/workflow/pkg/version"
)

var (
	versionExpression = regexp.MustCompile(`go (\d+\.\d+(?:\.\d+)?)`)
)

type runner struct {
	flag   *flag
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	err := r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run()
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run() error {
	{
		p := ".github/workflows/"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := ".github/workflows/go-build.yaml"

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

	{
		p := "go.mod"

		b, err := os.ReadFile(p)
		if os.IsNotExist(err) {
			// Just fall through since it can happen that repositories do not
			// always have a go.mod file.
		} else if err != nil {
			return tracer.Mask(err)
		} else {
			c := strings.Replace(string(b), currentVersion(b), desiredVersion(r.flag.Version.Golang), -1)

			err = os.WriteFile(p, []byte(c), 0600)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (r *runner) data() any {
	type Data struct {
		Command string
		Env     map[string]string
		Private string
		User    string
		Version version.Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
		Env:     env.Env(),
		Private: r.flag.Private,
		User:    r.flag.User,
		Version: version.Version{
			Checkout:     version.Checkout,
			Golang:       r.flag.Version.Golang,
			GolangCiLint: version.GolangCiLint,
			SetupGo:      version.SetupGo,
		},
	}
}

func currentVersion(b []byte) string {
	r := versionExpression.FindSubmatch(b)
	if len(r) != 2 {
		// FindSubmatch returns the full match and the capturing group. As such
		// we expect 2 results, of which the first is the current version string
		// we are interested in.
		//
		//     ["go 1.15" "1.15"]
		//
		panic("must find two results")
	}

	return string(r[0])
}

func desiredVersion(v string) string {
	s := strings.Split(v, ".")
	if len(s) != 3 {
		// The given version must be something like 1.15.2 which implies 3 parts
		// when splitting at the period.
		panic("must have 3 parts")
	}

	return fmt.Sprintf("go %s.%s", s[0], s[1])
}
