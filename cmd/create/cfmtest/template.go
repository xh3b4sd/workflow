package cfmtest

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "cfm-test"

on: "push"

jobs:
  cfm-test:
    runs-on: "ubuntu-latest"

    services:
      redis:
        image: "redis"
        options: >-
          --health-cmd "redis-cli ping"
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 6379:6379

    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v2"
        with:
          path: "venturemark/flux"
          repository: "venturemark/flux"

      - name: "Setup Git Project"
        uses: "actions/checkout@v2"
        with:
          path: "venturemark/fmz"
          repository: "venturemark/fmz"

      - name: "Setup Go Env"
        uses: actions/setup-go@v2
        with:
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          sudo apt update
          sudo apt install gcc -y

      - name: "Install Test Dependency"
        run: |
          go get github.com/xh3b4sd/dsm

      - name: "Install Test Dependency"
        env:
          GO111MODULE: "on"
        run: |
          go get github.com/venturemark/apiserver@$(dsm search -r HelmRelease -n apiserver -k spec.values.image.tag | head -n 1)

      - name: "Install Test Dependency"
        env:
          GO111MODULE: "on"
        run: |
          go get github.com/venturemark/apiworker@$(dsm search -r HelmRelease -n apiworker -k spec.values.image.tag | head -n 1)

      - name: "Install Test Dependency"
        env:
          GO111MODULE: "on"
        run: |
          go get github.com/venturemark/fmz

      - name: "Check Conformance Tests"
        env:
          CGO_ENABLED: "1"
        run: |
          apiserver daemon &
          apiworker daemon &
          cd ./venturemark/fmz && go test ./... -race -tags conformance
`
