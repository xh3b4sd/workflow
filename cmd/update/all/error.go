package all

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var commandNotFoundError = &tracer.Error{
	Description: "The desired command was not found in the list of available commands.",
}

func IsCommandNotFound(err error) bool {
	return errors.Is(err, commandNotFoundError)
}
