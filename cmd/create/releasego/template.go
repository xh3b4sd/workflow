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
          GOOS={{ $k }} GOARCH={{ $v }} go build -o {{ $.Repository.Name }}-{{ $k }}-{{ $v }} -ldflags="-X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/{{ $.Repository.Name }}/{{ $.Repository.Path }}.{{ $.Variable.GitSha }}=${{ "{{" }} github.sha {{ "}}" }}' -X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/{{ $.Repository.Name }}/{{ $.Repository.Path }}.{{ $.Variable.GitTag }}=${{ "{{" }} github.ref_name {{ "}}" }}'"
{{- end }}

      - name: "Upload To Github"
        uses: "softprops/action-gh-release@v1"
        with:
          files: |
{{- range $k, $v := .Release.Assets }}
            {{ $.Repository.Name }}-{{ $k }}-{{ $v }}
{{- end }}
`
