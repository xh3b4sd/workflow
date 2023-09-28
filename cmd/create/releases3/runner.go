package releases3

import (
	"bytes"
	"context"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
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
	type AWS struct {
		Bucket string
		Region string
	}

	type Release struct {
		Assets map[string]string
	}

	type Repository struct {
		Name string
		Path string
	}

	type Data struct {
		AWS        AWS
		Command    string
		Release    Release
		Repository Repository
		Version    version.Version
	}

	return Data{
		AWS: AWS{
			Bucket: r.flag.AWS.Bucket,
			Region: r.flag.AWS.Region,
		},
		Command: strings.Join(os.Args, " "),
		Release: Release{
			Assets: assets(r.flag.Release.Assets),
		},
		Repository: Repository{
			Name: r.flag.Repository.Name,
			Path: r.flag.Repository.Path,
		},
		Version: version.Version{
			Checkout: version.Checkout,
			Golang:   r.flag.Version.Golang,
			SetupGo:  version.SetupGo,
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
		p := ".github/workflows/s3-release.yaml"

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

func assets(str string) map[string]string {
	ass := map[string]string{}

	for _, s := range strings.Split(str, ",") {
		spl := strings.Split(s, "/")
		ass[spl[0]] = spl[1]
	}

	return ass
}
