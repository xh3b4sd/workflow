package command

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
	"github.com/xh3b4sd/workflow/pkg/parser"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_Command_Parse tests the parsing of the command instructions used to
// generate workflows. Different workflows and their associated commands are
// quite complex while the parsing code is not easily comprehensible. Therefore
// we need a way to reliably verify the integrity of the command parsing.
//
//	go test ./pkg/parser/command -run Test_Command_Parse -update
func Test_Command_Parse(t *testing.T) {
	testCases := []struct {
		fileSystem afero.Fs
	}{
		// Case 0 ensures that a command instruction can be parsed according to
		// its workflow header.
		{
			fileSystem: func() afero.Fs {
				fs := afero.NewMemMapFs()

				mustCreateDir(fs, ".github/workflows/")

				mustCreateFile(fs, ".github/workflows/some.yaml", `
				#
				#     workflow create foo bar
				#
				`)

				return fs
			}(),
		},
		// Case 1 ensures that a command instruction can be parsed according to
		// its workflow header.
		{
			fileSystem: func() afero.Fs {
				fs := afero.NewMemMapFs()

				mustCreateDir(fs, ".github/workflows/")

				mustCreateFile(fs, ".github/workflows/some.yaml", `
				#
				#     workflow create bar foo
				#
				`)

				mustCreateFile(fs, ".github/workflows/dependabot.yaml", `
				#
				#     workflow create dependabot -r @xh3b4sd
				#
				`)

				mustCreateFile(fs, ".github/workflows/grpc-ts.yaml", `
				#
				#     workflow create grpcts -n 15.x.x
				#
				`)

				mustCreateFile(fs, ".github/dependabot.yaml", `
				#
				#     workflow create dependabot should not be taken into account due to path
				#
				`)

				return fs
			}(),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var p parser.Interface
			{
				p = New(Config{
					FileSystem: tc.fileSystem,

					WorkflowPath: ".github/workflows/",
				})
			}

			var actual string
			{
				l, err := p.Parse()
				if err != nil {
					t.Fatal(err)
				}

				var s []string
				for _, a := range l {
					s = append(s, strings.Join(a, " "))
				}

				actual = strings.Join(s, "\n") + "\n"
			}

			var expected []byte
			{
				p := filepath.Join("testdata/parse", fmt.Sprintf("case.%03d.golden", i))
				if *update {
					err := os.WriteFile(p, []byte(actual), 0600)
					if err != nil {
						t.Fatal(err)
					}
				}

				expected, err = os.ReadFile(p)
				if err != nil {
					t.Fatal(err)
				}
			}

			if !bytes.Equal(expected, []byte(actual)) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(actual), string(expected)))
			}
		})
	}
}

func mustCreateDir(fs afero.Fs, p string) {
	err := fs.MkdirAll(p, 0755)
	if err != nil {
		panic(err)
	}
}

func mustCreateFile(fs afero.Fs, p string, b string) {
	err := afero.WriteFile(fs, p, []byte(b), 0600)
	if err != nil {
		panic(err)
	}
}
