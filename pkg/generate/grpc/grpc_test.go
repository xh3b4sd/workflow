package grpc

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_Generate_GRPC tests workflow generation for the gRPC code generation
// workflow. The workflow template is quite complex and not easily readable.
// Considering input parameter like Github organization and Golang version we
// need a way to reliable verify the integrity of the YAML file rendering.
//
//     go test ./... -run Test_Generate_GRPC -update
//
func Test_Generate_GRPC(t *testing.T) {
	testCases := []struct {
		organization string
		repository   string
		golang       string
		protoc       string
	}{
		// Case 0 ensures that a workflow file can be generated according to its
		// configuration.
		{
			organization: "xh3b4sd",
			repository:   "gocode",
			golang:       "1.15.2",
			protoc:       "3.13.0",
		},
		// Case 1 ensures that a workflow file can be generated according to its
		// configuration.
		{
			organization: "some-org",
			repository:   "some-repo",
			golang:       "1.14.0",
			protoc:       "3.5.1",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var g Interface
			{
				c := Config{
					FilePath:           "workflow.yaml",
					GithubOrganization: tc.organization,
					GithubRepository:   tc.repository,
					VersionGolang:      tc.golang,
					VersionProtoc:      tc.protoc,
				}

				g, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var actual []byte
			{
				actual, err = g.Generate()
				if err != nil {
					t.Fatal(err)
				}
			}

			p := filepath.Join("testdata/grpc", fileName(i))
			if *update {
				err := ioutil.WriteFile(p, []byte(actual), 0644) // nolint:gosec
				if err != nil {
					t.Fatal(err)
				}
			}

			expected, err := ioutil.ReadFile(p)
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
