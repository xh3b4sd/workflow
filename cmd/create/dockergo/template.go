package dockergo

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "docker-go"

on: "push"

jobs:
  docker-go:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Build Go Binary"
        env:
          CGO_ENABLED: "0"
        run: |
          go build .

      - name: "Setup Docker Buildx"
        uses: "docker/setup-buildx-action@v1.3.0"

      - name: "Login Container Registry"
        uses: "docker/login-action@v1.9.0"
        with:
          registry: "ghcr.io"
          username: "${{ "{{" }} github.repository_owner {{ "}}" }}"
          password: "${{ "{{" }} secrets.CONTAINER_REGISTRY_TOKEN {{ "}}" }}"

      - name: "Build Docker Image"
        uses: "docker/build-push-action@v2.5.0"
        with:
          context: "."
          push: true
          tags: "ghcr.io/${{ "{{" }} github.repository {{ "}}" }}:${{ "{{" }} github.sha {{ "}}" }}"
`
