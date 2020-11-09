package grpcts

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/xh3b4sd/workflow/pkg/generator"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_GrpcTs_Usage tests the generation of the additional command instructions
// when generating the workflow. The usage template is quite complex and not
// easily readable. Considering input parameter like Github organization and the
// current repository we need a way to reliable verify the integrity of the YAML
// file rendering.
//
//     go test ./pkg/generator/... -run Test_GrpcTs_Usage -update
//
func Test_GrpcTs_Usage(t *testing.T) {
	testCases := []struct {
		current      string
		organization string
		repository   string
	}{
		// Case 0 ensures that a command instruction can be generated according
		// to its configuration.
		{
			current:      "github.com/xh3b4sd/workflow",
			organization: "xh3b4sd",
			repository:   "tscode",
		},
		// Case 1 ensures that a command instruction can be generated according
		// to its configuration.
		{
			current:      "github.com/some-org/some-repo",
			organization: "some-org",
			repository:   "some-repo",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var g generator.Interface
			{
				c := Config{
					Command:            "workflow create grpcts",
					FilePath:           "workflow.yaml",
					GithubCurrent:      tc.current,
					GithubOrganization: tc.organization,
					GithubRepository:   tc.repository,
					VersionGolang:      "1.15.2",
					VersionGrpcWeb:     "1.2.1",
					VersionNode:        "15.x.x",
					VersionProtoc:      "3.13.0",
				}

				g, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var actual []byte
			{
				actual, err = g.Usage()
				if err != nil {
					t.Fatal(err)
				}
			}

			p := filepath.Join("testdata/usage", fileName(i))
			if *update {
				err := ioutil.WriteFile(p, []byte(actual), 0600)
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

// Test_GrpcTs_Workflow tests workflow generation for the gRPC Typescript code
// generation workflow. The workflow template is quite complex and not easily
// readable. Considering input parameter like Github organization and Typescript
// version we need a way to reliable verify the integrity of the YAML file
// rendering.
//
//     go test ./... -run Test_GrpcTs_Workflow -update
//
func Test_GrpcTs_Workflow(t *testing.T) {
	testCases := []struct {
		command      string
		organization string
		repository   string
		golang       string
		grpcWeb      string
		node         string
		protoc       string
	}{
		// Case 0 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create grpcts",
			organization: "xh3b4sd",
			repository:   "tscode",
			golang:       "1.15.2",
			grpcWeb:      "1.2.1",
			node:         "15.x.x",
			protoc:       "3.13.0",
		},
		// Case 1 ensures that a workflow file can be generated according to its
		// configuration.
		{
			command:      "workflow create grpcts --some argument",
			organization: "some-org",
			repository:   "some-repo",
			golang:       "1.14.0",
			grpcWeb:      "2.1.5",
			node:         "15.x.x",
			protoc:       "3.5.1",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var g generator.Interface
			{
				c := Config{
					Command:            tc.command,
					FilePath:           "workflow.yaml",
					GithubCurrent:      "github.com/xh3b4sd/workflow",
					GithubOrganization: tc.organization,
					GithubRepository:   tc.repository,
					VersionGolang:      tc.golang,
					VersionGrpcWeb:     tc.grpcWeb,
					VersionNode:        tc.node,
					VersionProtoc:      tc.protoc,
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
				err := ioutil.WriteFile(p, []byte(actual), 0600)
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
