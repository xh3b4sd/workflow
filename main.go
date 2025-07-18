package main

import (
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/cmd"
)

func main() {
	err := cmd.New().Execute()
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}
}
