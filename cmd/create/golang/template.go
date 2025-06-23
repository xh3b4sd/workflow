package golang

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "go-build"

on: "push"

jobs:
  go-build:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"
{{ if .Private }}
      - name: "Setup Private Dependencies"
        env:
          PAT: "${{ "{{" }} secrets.GOPRIVATE_PAT {{ "}}" }}"
        run: |
          git config --global url."https://{{ .User }}:${PAT}@github.com".insteadOf "https://github.com"
{{ end }}
      - name: "Check Go Dependencies"
{{- if .Private }}
        env:
          GOPRIVATE: "{{ .Private }}"
          PAT: "${{ "{{" }} secrets.GOPRIVATE_PAT {{ "}}" }}"
{{- end }}
        run: |
          go mod tidy
          git diff --exit-code

      - name: "Build Go Binary"
        env:
          CGO_ENABLED: "0"
{{- range $k, $v := .Env }}
          {{ $k }}: "{{ $v }}"
{{- end }}
        run: |
          go build .

      - name: "Check Go Tests"
{{- if .Env }}
        env:
{{- end }}
{{- range $k, $v := .Env }}
          {{ $k }}: "{{ $v }}"
{{- end }}
        run: |
          go test ./... -race

      - name: "Check Go Formatting"
        run: |
          test -z $(gofmt -l -s .)

      - name: "Check Go Linters"
        run: |
          curl -LOs https://github.com/golangci/golangci-lint/releases/download/v{{ .Version.GolangCiLint }}/golangci-lint-{{ .Version.GolangCiLint }}-linux-amd64.tar.gz
          tar -xzf golangci-lint-{{ .Version.GolangCiLint }}-linux-amd64.tar.gz
          ./golangci-lint-{{ .Version.GolangCiLint }}-linux-amd64/golangci-lint run
`
