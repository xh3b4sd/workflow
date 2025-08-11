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
    ignore:
      - dependency-name: "*"
        update-types: ["version-update:semver-patch"]
    open-pull-requests-limit: 10
    schedule:
      interval: "daily"
      time: "04:00"
{{ end }}`
