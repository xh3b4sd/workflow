package cfmtest

const templateApiServer = `#
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
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/apiserver"
          repository: "venturemark/apiserver"

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/cfm"
          repository: "venturemark/cfm"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          sudo apt update
          sudo apt install gcc -y

      - name: "Install Test Dependency"
        run: |
          cd ./venturemark/apiserver && go install .

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiworker@$(git ls-remote git://github.com/venturemark/apiworker.git HEAD | awk '{print $1;}')

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/cfm@$(git ls-remote git://github.com/venturemark/cfm.git HEAD | awk '{print $1;}')

      - name: "Check Conformance Tests"
        env:
          CGO_ENABLED: "1"
          APIWORKER_POSTMARK_TOKEN_ACCOUNT: "foo"
          APIWORKER_POSTMARK_TOKEN_SERVER: "foo"
        run: |
          apiserver daemon --metrics-port 8081 &
          apiworker daemon --metrics-port 8082 &
          cd ./venturemark/cfm && go test ./... -race -tags conformance
`

const templateApiWorker = `#
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
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/apiworker"
          repository: "venturemark/apiworker"

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/cfm"
          repository: "venturemark/cfm"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          sudo apt update
          sudo apt install gcc -y

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiserver@$(git ls-remote git://github.com/venturemark/apiserver.git HEAD | awk '{print $1;}')

      - name: "Install Test Dependency"
        run: |
          cd ./venturemark/apiworker && go install .

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/cfm@$(git ls-remote git://github.com/venturemark/cfm.git HEAD | awk '{print $1;}')

      - name: "Check Conformance Tests"
        env:
          CGO_ENABLED: "1"
          APIWORKER_POSTMARK_TOKEN_ACCOUNT: "foo"
          APIWORKER_POSTMARK_TOKEN_SERVER: "foo"
        run: |
          apiserver daemon --metrics-port 8081 &
          apiworker daemon --metrics-port 8082 &
          cd ./venturemark/cfm && go test ./... -race -tags conformance
`

const templateCfm = `#
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
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/cfm"
          repository: "venturemark/cfm"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          sudo apt update
          sudo apt install gcc -y

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiserver@$(git ls-remote git://github.com/venturemark/apiserver.git HEAD | awk '{print $1;}')

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiworker@$(git ls-remote git://github.com/venturemark/apiworker.git HEAD | awk '{print $1;}')

      - name: "Install Test Dependency"
        env:
          GO111MODULE: "on"
        run: |
          # TODO this needs to be "go install ." again once cfm can be compiled
          # into an executable like apiserver and apiworker.
          cd ./venturemark/cfm && go test ./...

      - name: "Check Conformance Tests"
        env:
          CGO_ENABLED: "1"
          APIWORKER_POSTMARK_TOKEN_ACCOUNT: "foo"
          APIWORKER_POSTMARK_TOKEN_SERVER: "foo"
        run: |
          apiserver daemon --metrics-port 8081 &
          apiworker daemon --metrics-port 8082 &
          cd ./venturemark/cfm && go test ./... -race -tags conformance
`

const templateFlux = `#
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
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/cfm"
          repository: "venturemark/cfm"

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"
        with:
          path: "venturemark/flux"
          repository: "venturemark/flux"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Test Dependency"
        run: |
          sudo apt update
          sudo apt install gcc -y

      - name: "Install Test Dependency"
        env:
          GO111MODULE: "on"
        run: |
          go get -u github.com/xh3b4sd/dsm

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiserver@$(dsm search -r HelmRelease -n apiserver -k spec.values.image.tag | head -n 1)

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/apiworker@$(dsm search -r HelmRelease -n apiworker -k spec.values.image.tag | head -n 1)

      - name: "Install Test Dependency"
        run: |
          go get -u github.com/venturemark/cfm@$(git ls-remote git://github.com/venturemark/cfm.git HEAD | awk '{print $1;}')

      - name: "Check Conformance Tests"
        env:
          CGO_ENABLED: "1"
          APIWORKER_POSTMARK_TOKEN_ACCOUNT: "foo"
          APIWORKER_POSTMARK_TOKEN_SERVER: "foo"
        run: |
          apiserver daemon --metrics-port 8081 &
          apiworker daemon --metrics-port 8082 &
          cd ./venturemark/cfm && go test ./... -race -tags conformance
`
