package golang

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

var (
	versionExpression = regexp.MustCompile(`go ([0-9]+\.[0-9]+)`)
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
	type Version struct {
		Golang string
	}

	type Data struct {
		Command string
		Version Version
	}

	return Data{
		Command: strings.Join(os.Args, " "),
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

		err = ioutil.WriteFile(p, b.Bytes(), 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := "go.mod"

		b, err := ioutil.ReadFile(p)
		if os.IsNotExist(err) {
			// Just fall through since it can happen that repositories do not
			// always have a go.mod file.
		} else if err != nil {
			return tracer.Mask(err)
		} else {
			c := strings.Replace(string(b), currentVersion(b), desiredVersion(r.flag.Version.Golang), -1)

			err = ioutil.WriteFile(p, []byte(c), 0600)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func currentVersion(b []byte) string {
	r := versionExpression.FindSubmatch(b)
	if len(r) != 2 {
		// FindSubmatch returns the full match and the capturing group. As such
		// we expect 2 results, of which the second is the current version we
		// are interested in.
		//
		//     ["go 1.15" "1.15"]
		//
		panic("must find two results")
	}

	return string(r[1])
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
