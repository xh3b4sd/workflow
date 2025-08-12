package dependabot

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

version: 2

updates:
{{- range $e := .Ecosystems }}
  - package-ecosystem: "{{ $e.Name }}"
    directory: "/"
    open-pull-requests-limit: 10
    schedule:
      interval: "weekly"
      time: "04:00"
{{- if $e.Group.Aws }}
    groups:
      aws-sdk-go-v2:
        patterns:
          - "github.com/aws/aws-sdk-go-v2"
          - "github.com/aws/aws-sdk-go-v2/*"
{{- end }}
{{- end }}
`
