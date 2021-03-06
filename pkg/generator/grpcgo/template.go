package grpcgo

const templateUsage = `Generated grpc workflow for golang code generation. Please make sure to
generate the RSA deploy keys and the GPG password as follows. For more
information see https://github.com/xh3b4sd/red.

    red generate keys -d .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}

Upon RSA deploy key generation as described above, make sure to add the
generated public key as deploy key to the following github repository.

    github.com/{{ .Github.Organization }}/{{ .Github.Repository }}

Upon GPG password generation as described above, make sure to add a secret
environment variable to the following github repository. Further use the
following secret name.

    repository:     {{ .Github.Current }}

    secret:         RED_GPG_PASS_{{ .Github.Organization | ToUpper }}_{{ .Github.Repository | ToUpper }}

`

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "grpc-go"

on:
  push:
    branches:
      - "main"
      - "master"
    paths:
      - "**.proto"

jobs:
  grpc-go:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Setup Git Project"
        uses: "actions/checkout@v2.3.4"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v2"
        with:
          go-version: "{{ .Version.Golang }}"

      - name: "Install Protoc Binary"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}"
        run: |
          curl -LOs https://github.com/protocolbuffers/protobuf/releases/download/v{{ .Version.Protoc }}/protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          unzip protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          echo "${{ "{{" }} runner.temp {{ "}}" }}/bin" >> $GITHUB_PATH

      - name: "Install gRPC Dependencies"
        env:
          GO111MODULE: "on"
        run: |
          go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
          go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1

      - name: "Decrypt Private Key"
        run: |
          go get github.com/xh3b4sd/red
          red decrypt -i .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa.enc -o .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa -p '${{ "{{" }} secrets.RED_GPG_PASS_{{ .Github.Organization | ToUpper }}_{{ .Github.Repository | ToUpper }} {{ "}}" }}'

      - name: "Setup SSH Agent"
        env:
          SSH_AUTH_SOCK: "/tmp/ssh_agent.sock"
        run: |
          mkdir -p ~/.ssh
          ssh-keyscan github.com >> ~/.ssh/known_hosts
          ssh-agent -a ${SSH_AUTH_SOCK} > /dev/null
          chmod 0600 .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa
          ssh-add .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa

      - name: "Clone Go Code"
        env:
          SSH_AUTH_SOCK: "/tmp/ssh_agent.sock"
        run: git clone git@github.com:{{ .Github.Organization }}/{{ .Github.Repository }}.git "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"

      - name: "Setup Git Config"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
        run: |
          git config user.name "${GITHUB_ACTOR}"
          git config user.email "${GITHUB_ACTOR}@users.noreply.github.com"
          git remote set-url origin git@github.com:{{ .Github.Organization }}/{{ .Github.Repository }}.git

      - name: "Generate Go Code"
        run: |
          go get github.com/xh3b4sd/pag
          rm -rf ${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/pkg/
          pag generate golang -d ${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/pkg/

      - name: "Go Mod Tidy"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
        run: |
          if [[ -e go.sum ]]; then
            rm -f go.sum
            go mod tidy
          fi

      - name: "Commit And Push"
        env:
          SSH_AUTH_SOCK: "/tmp/ssh_agent.sock"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
        run: |
          if [[ $(git status --porcelain) ]]; then
            git add .
            git commit -m 'update generated code'
            git push
          fi

      - name: "Cleanup Build Container"
        env:
          SSH_AUTH_SOCK: "/tmp/ssh_agent.sock"
        run: |
          ssh-add -D
          rm -Rf *
`
