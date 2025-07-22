package releasego

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "go-release"

on:
  push:
    tags:
      - "v*.*.*"

jobs:
  release:
    permissions:
      contents: write

    runs-on: "ubuntu-latest"
    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Cross Compile Binaries"
        run: |
{{- range $k, $v := .Release.Assets }}
          GOOS={{ $k }} GOARCH={{ $v }} go build -o ${{ "{{" }} github.event.repository.name {{ "}}" }}-{{ $k }}-{{ $v }} -ldflags="-X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/${{ "{{" }} github.event.repository.name {{ "}}" }}/{{ $.Linker.Path }}.{{ $.Linker.Git.Sha }}=${{ "{{" }} github.sha {{ "}}" }}' -X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/${{ "{{" }} github.event.repository.name {{ "}}" }}/{{ $.Linker.Path }}.{{ $.Linker.Git.Tag }}=${{ "{{" }} github.ref_name {{ "}}" }}'"
{{- end }}

      - name: "Upload To Github"
        uses: "softprops/action-gh-release@v2"
        with:
          files: |
{{- range $k, $v := .Release.Assets }}
            ${{ "{{" }} github.event.repository.name {{ "}}" }}-{{ $k }}-{{ $v }}
{{- end }}
`
