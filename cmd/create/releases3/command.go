package releases3

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "releases3"
	short = "Create a golang workflow for e.g. uploading cross compiled release assets."
	long  = `Create a golang workflow for e.g. uploading cross compiled release assets.
Uploading release assets to Github is straight forward and works out of the box.
Uploading release assets to AWS S3 requires some configurations in the
respective Github repository and AWS account.

The respective Github repository MUST have AWS typical credentials configured
via repository secrets. The required secret names must be set as follows.

	AWS_ACCESS_KEY
	AWS_SECRET_KEY

The respective AWS account SHOULD use a restricted IAM policy based on the least
priviledged principle. This IAM policy would be attached to the AWS credentials
used in this workflow. A working example of the minimal required IAM policy
might look like shown below.

	{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Sid": "VisualEditor0",
				"Effect": "Allow",
				"Action": [
					"s3:PutObject",
					"s3:PutBucketOwnershipControls",
					"s3:CreateBucket"
				],
				"Resource": [
					"arn:aws:s3:::<aws-bucket>",
					"arn:aws:s3:::<aws-bucket>/*"
				]
			}
		]
	}
`
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var c *cobra.Command
	{
		f := &flag{}

		r := &runner{
			flag:   f,
			logger: config.Logger,
		}

		c = &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}

		f.Init(c)
	}

	return c, nil
}
