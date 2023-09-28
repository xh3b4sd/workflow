package releases3

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/pkg/version"
)

type flag struct {
	AWS struct {
		Bucket string
		Region string
	}
	Release struct {
		Assets string
	}
	Repository struct {
		Name string
		Path string
	}
	Version struct {
		Golang string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.AWS.Bucket, "aws-bucket", "b", "", "AWS S3 bucket to upload to.")
	cmd.Flags().StringVarP(&f.AWS.Region, "aws-region", "r", "eu-central-1", "AWS S3 region to upload to.")
	cmd.Flags().StringVarP(&f.Release.Assets, "release-assets", "a", "darwin/amd64,linux/amd64", "Binary architectures to compile.")
	cmd.Flags().StringVarP(&f.Repository.Name, "repository-name", "n", "", "Repository name to generate the workflow for, e.g. apiworker or webclient.")
	cmd.Flags().StringVarP(&f.Repository.Path, "repository-path", "p", "pkg/project", "Repository path to compile linker flags into.")
	cmd.Flags().StringVarP(&f.Version.Golang, "version-golang", "g", version.Golang, "Golang version to use in, e.g. workflow files.")
}

func (f *flag) Validate() error {
	{
		if f.AWS.Bucket == "" {
			return tracer.Maskf(invalidFlagError, "-b/--aws-bucket must not be empty")
		}
		if f.AWS.Region == "" {
			return tracer.Maskf(invalidFlagError, "-r/--aws-region must not be empty")
		}
	}

	{
		if f.Release.Assets == "" {
			return tracer.Maskf(invalidFlagError, "-a/--release-assets must not be empty")
		}
	}

	{
		if f.Repository.Name == "" {
			return tracer.Maskf(invalidFlagError, "-n/--repository-name must not be empty")
		}
		if f.Repository.Path == "" {
			return tracer.Maskf(invalidFlagError, "-p/--repository-path must not be empty")
		}
	}

	{
		if f.Version.Golang == "" {
			return tracer.Maskf(invalidFlagError, "-g/--version-golang must not be empty")
		}

		s := strings.Split(f.Version.Golang, ".")
		if len(s) != 3 {
			return tracer.Maskf(invalidFlagError, "-g/--version-golang must have 3 parts like 1.15.2")
		}
	}

	return nil
}
