package pbfts

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

// Test_PbfTs_Workflow tests workflow generation for the protocol buffer
// Typescript code generation workflow. The workflow template is quite complex
// and not easily readable. Considering input parameter like Github organization
// and Node version we need a way to reliable verify the integrity of the YAML
// file rendering.
//
//	go test ./pkg/generator/pbfts -update
func Test_GrpcTs_Workflow(t *testing.T) {
	testCases := []struct {
		command      string
		organization string
		repository   string
		node         string
		protoc       string
	}{
		// Case 0 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create pbfts",
			organization: "xh3b4sd",
			repository:   "tscode",
			node:         "20.x.x",
			protoc:       "3.13.0",
		},
		// Case 1 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create pbfts --some argument",
			organization: "some-org",
			repository:   "some-repo",
			node:         "20.x.x",
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
					VersionNode:        tc.node,
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
