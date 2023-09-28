package redigo

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "go-redis"

on: "push"

jobs:
  go-redis:
    runs-on: "ubuntu-latest"
    container: "golang:{{ .Version.Golang }}-alpine"

    services:
      redis:
        image: "redis"
        options: >-
          --health-cmd "redis-cli ping"
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

      redis-sentinel:
        image: "bitnami/redis-sentinel:6.0"
        ports:
            - "26379:26379"
        env:
          REDIS_SENTINEL_QUORUM: "1"

    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Install Race Dependency"
        run: |
          apk add gcc
          apk add musl-dev

      - name: "Check Go Tests"
        env:
          CGO_ENABLED: "1"
          REDIS_HOST: "redis"
          REDIS_PORT: "6379"
          REDIS_SENTINEL_HOST: "redis-sentinel"
          REDIS_SENTINEL_PORT: "26379"
        run: |
          go test ./... -race -tags single,sentinel
`
