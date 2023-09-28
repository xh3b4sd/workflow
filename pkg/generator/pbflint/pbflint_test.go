package pbflint

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/xh3b4sd/workflow/pkg/generator"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_PbfLint_Workflow tests workflow generation for the protocol buffer
// validation workflow.
//
//	go test ./pkg/generator/pbflint -update
func Test_PbfLint_Workflow(t *testing.T) {
	testCases := []struct {
		command string
		golang  string
		protoc  string
	}{
		// Case 0 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command: "workflow create pbflint",
			golang:  "1.15.2",
			protoc:  "3.13.0",
		},
		// Case 1 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command: "workflow create pbflint --some argument",
			golang:  "1.14.0",
			protoc:  "3.5.1",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var g generator.Interface
			{
				c := Config{
					Command:       tc.command,
					FilePath:      "workflow.yaml",
					VersionGolang: tc.golang,
					VersionProtoc: tc.protoc,
				}

				g, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var actual []byte
			{
				actual, err = g.Workflow()
				if err != nil {
					t.Fatal(err)
				}
			}

			p := filepath.Join("testdata/workflow", fileName(i))
			if *update {
				err := os.WriteFile(p, []byte(actual), 0600)
				if err != nil {
					t.Fatal(err)
				}
			}

			expected, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(expected, []byte(actual)) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(actual), string(expected)))
			}
		})
	}
}

func fileName(i int) string {
	return "case-" + strconv.Itoa(i) + ".golden"
}
