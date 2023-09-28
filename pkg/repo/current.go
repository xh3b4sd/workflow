package repo

import (
	"path/filepath"
	"strings"
)

// Current tries to lookup the location of the current Github repository. If
// Current fails to determine the absolute file system path of the program's
// location it panics. Consider the current absolut path.
//
//	/Users/xh3b4sd/go/src/github.com/xh3b4sd/workflow/
//	/Users/xh3b4sd/projects/xh3b4sd/apischema
//
// The resulting Github repository returned would be the following respectively.
//
//	github.com/xh3b4sd/workflow
//	github.com/xh3b4sd/apischema
func Current() string {
	a, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	return current(a)
}

func current(a string) string {
	l := strings.Split(a, "/")

	{
		var c int
		var p []string
		var t bool

		for _, s := range l {
			if t {
				p = append(p, s)
				c++
			}

			if s == "github.com" {
				p = append(p, s)
				t = true
			}

			if c >= 2 {
				return filepath.Join(p...)
			}
		}
	}

	{
		var p []string

		p = append(p, "github.com")
		p = append(p, l[len(l)-2:]...)

		return filepath.Join(p...)
	}
}
