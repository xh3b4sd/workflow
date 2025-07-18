package runtime

import (
	"github.com/xh3b4sd/tracer"
)

var InvalidFlagError = &tracer.Error{
	Description: "At least one command line flag was missing or misconfigured.",
}
