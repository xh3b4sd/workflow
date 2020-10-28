package repo

import (
	"path/filepath"
	"strings"
)

// Current tries to lookup the name of the current Github repository. If Current
// fails to determine the absolute file system path of the program's location it
// panics. Consider the current absolut path.
//
//     /Users/xh3b4sd/go/src/github.com/xh3b4sd/workflow/
//
// The resulting Github repository returned would be the following.
//
//     github.com/xh3b4sd/workflow
//
func Current() string {
	a, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	var c int
	var p []string
	var t bool
	for _, s := range strings.Split(a, "/") {
		if t {
			p = append(p, s)
			c++
		}

		if s == "github.com" {
			p = append(p, s)
			t = true
		}

		if c >= 2 {
			break
		}
	}

	return filepath.Join(p...)
}
