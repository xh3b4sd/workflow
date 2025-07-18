package repo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Repo_current(t *testing.T) {
	testCases := []struct {
		abs string
		cur string
	}{
		// Case 0 ensures the current repo based on the absolut path structured
		// according to legacy golang conventions.
		{
			abs: "/Users/xh3b4sd/go/src/github.com/xh3b4sd/workflow/",
			cur: "github.com/xh3b4sd/workflow",
		},
		// Case 1 ensures the current repo based on the absolut path structured
		// according to some custom location.
		{
			abs: "/Users/xh3b4sd/projects/xh3b4sd/apischema",
			cur: "github.com/xh3b4sd/apischema",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			cur := current(tc.abs)

			if tc.cur != cur {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.cur, cur))
			}
		})
	}
}
