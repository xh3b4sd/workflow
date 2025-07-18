package pbfgo

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/xh3b4sd/workflow/pkg/generator"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_PbfGo_Workflow tests workflow generation for the protocol buffer Golang
// code generation workflow. The workflow template is quite complex and not
// easily readable. Considering input parameter like Github organization and
// Golang version we need a way to reliable verify the integrity of the YAML
// file rendering.
//
//	go test ./pkg/generator/pbfgo -update
func Test_PbfGo_Workflow(t *testing.T) {
	testCases := []struct {
		command      string
		organization string
		repository   string
		golang       string
		protoc       string
	}{
		// Case 0 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create pbfgo",
			organization: "xh3b4sd",
			repository:   "gocode",
			golang:       "1.15.2",
			protoc:       "3.13.0",
		},
		// Case 1 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create pbfgo --some argument",
			organization: "some-org",
			repository:   "some-repo",
			golang:       "1.14.0",
			protoc:       "3.5.1",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var g generator.Interface
			{
				g = New(Config{
					Command:            tc.command,
					FilePath:           "workflow.yaml",
					GithubOrganization: tc.organization,
					GithubRepository:   tc.repository,
					VersionGolang:      tc.golang,
					VersionProtoc:      tc.protoc,
				})
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
