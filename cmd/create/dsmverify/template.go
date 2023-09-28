package dsmverify

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "dsm-verify"

on: "push"

jobs:
  dsm-verify:
    runs-on: "ubuntu-latest"

    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          go get github.com/xh3b4sd/dsm

      - name: "Check ApiServer Version"
        run: |
          dsm verify -r HelmRelease -n apiserver -k spec.values.image.tag

      - name: "Check ApiWorker Version"
        run: |
          dsm verify -r HelmRelease -n apiworker -k spec.values.image.tag

      - name: "Check WebClient Version"
        run: |
          dsm verify -r HelmRelease -n webclient -k spec.values.image.tag
`
