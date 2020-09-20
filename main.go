package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/workflow/cmd"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		mErr, ok := tracer.Cause(err).(*tracer.Error)
		if ok && mErr.Desc != "" {
			fmt.Println(strings.Title(err.Error()))
			fmt.Println()
			fmt.Println("    " + mErr.Desc)
			fmt.Println()
			os.Exit(1)
		} else {
			panic(tracer.JSON(err))
		}
	}
}

func mainE(ctx context.Context) error {
	var err error

	var l logger.Interface
	{
		c := logger.Config{}

		l, err = logger.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var r *cobra.Command
	{
		c := cmd.Config{
			Logger: l,
		}

		r, err = cmd.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	err = r.Execute()
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}
