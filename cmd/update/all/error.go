package all

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var commandNotFoundError = &tracer.Error{
	Kind: "commandNotFoundError",
}

func IsCommandNotFound(err error) bool {
	return errors.Is(err, commandNotFoundError)
}

var invalidConfigError = &tracer.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}
