package dockerts

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "docker-ts"

on: "push"

jobs:
  docker-ts:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Checkout Git Project"
        uses: "actions/checkout@v2"

      - name: "Setup Ts Env"
        uses: "actions/setup-node@v2"
        with:
          node-version: "{{ .Version.Node }}"

      - name: "Install Ts Dependencies"
        run: |
          npm install

      - name: "Build Ts Project"
        run: |
          npm run build

      - name: "Setup Docker Buildx"
        uses: "docker/setup-buildx-action@v1"

      - name: "Login Container Registry"
        uses: "docker/login-action@v1"
        with:
          registry: "ghcr.io"
          username: "${{ "{{" }} github.repository_owner {{ "}}" }}"
          password: "${{ "{{" }} secrets.CONTAINER_REGISTRY_TOKEN {{ "}}" }}"

      - name: "Build Docker Image"
        uses: "docker/build-push-action@v2"
        with:
          context: "."
          push: true
          tags: "ghcr.io/${{ "{{" }} github.repository {{ "}}" }}:${{ "{{" }} github.sha {{ "}}" }}"
`
