package pbfgo

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "Pbf Go"

on:
  push:
    branches:
      - "main"
      - "master"
    paths:
      - "**.proto"
  workflow_dispatch:

jobs:
  pbf-go:
    runs-on: "ubuntu-latest"
    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Install Protoc Binary"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}"
        run: |
          curl -LOs https://github.com/protocolbuffers/protobuf/releases/download/v{{ .Version.Protoc }}/protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          unzip protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          echo "${{ "{{" }} runner.temp {{ "}}" }}/bin" >> $GITHUB_PATH

      - name: "Install Go Dependencies"
        run: |
          go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.10
          go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.3

      - name: "Clone Go Code"
        run: |
          git clone https://github.com/{{ .Github.Organization }}/{{ .Github.Repository }}.git "${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"

      - name: "Generate Go Code"
        run: |
          inp="./pbf"
          out=${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/pkg

          rm -rf $out

          for x in $(ls $inp); do
            if [ -n "$(ls $inp/$x)" ]; then
              mkdir -p $out/$x

              lin=()
              for y in $(ls -F $inp/$x); do
                lin+=($inp/$x/$y)
              done

              protoc --go_out=$out/$x --twirp_out=$out/$x ${lin[@]}
            fi
          done

      - name: "Go Mod Tidy"
        working-directory: "${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
        run: |
          if [[ -e go.sum ]]; then
            rm -f go.sum
            go mod tidy
          fi

      - name: "Commit And Push"
        uses: "cpina/github-action-push-to-another-repository@v1.7.3"
        env:
          SSH_DEPLOY_KEY: "${{ "{{" }} secrets.SSH_DEPLOY_KEY_{{ .Github.Repository | ToUpper }} {{ "}}" }}"
        with:
          source-directory: "${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
          destination-github-username: "{{ .Github.Organization }}"
          destination-repository-name: "{{ .Github.Repository }}"
          commit-message: "update generated code"
          target-branch: "main"
`
