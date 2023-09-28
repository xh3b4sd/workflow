package releases3

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "s3-release"

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
          GOOS={{ $k }} GOARCH={{ $v }} go build -o {{ $.Repository.Name }}-{{ $k }}-{{ $v }} -ldflags="-X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/{{ $.Repository.Name }}/{{ $.Repository.Path }}.gitSHA=${{ "{{" }} github.sha {{ "}}" }}' -X 'github.com/${{ "{{" }} github.repository_owner {{ "}}" }}/{{ $.Repository.Name }}/{{ $.Repository.Path }}.version=${{ "{{" }} github.ref_name {{ "}}" }}'"
{{- end }}

      - name: "Configure AWS Credentials"
        uses: "aws-actions/configure-aws-credentials@v1"
        with:
          aws-access-key-id: "${{ "{{" }} secrets.AWS_ACCESS_KEY {{ "}}" }}"
          aws-secret-access-key: "${{ "{{" }} secrets.AWS_SECRET_KEY {{ "}}" }}"
          aws-region: "{{ .AWS.Region }}"

      - name: "Ensure S3 Bucket"
        run: |
          aws s3api create-bucket --bucket {{ .AWS.Bucket }} --create-bucket-configuration '{"LocationConstraint": "{{ .AWS.Region }}"}' || true

      - name: "Upload To S3"
        run: |
{{- range $k, $v := .Release.Assets }}
          aws s3 cp {{ $.Repository.Name }}-{{ $k }}-{{ $v }} s3://{{ $.AWS.Bucket }}/{{ $.Repository.Name }}/${{ "{{" }} github.ref_name {{ "}}" }}/{{ $.Repository.Name }}-{{ $k }}-{{ $v }}
{{- end }}
`
