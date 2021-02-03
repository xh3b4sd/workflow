package grpcts

const templateUsage = `Generated grpc workflow for typescript code generation. Please make sure to
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

name: gprc-ts

on:
  push:
    branches:
      - "main"
      - "master"
    paths:
      - "**.proto"

jobs:
  grpc-ts:
    runs-on: ubuntu-latest
    steps:

      - name: Setup Git Project
        uses: actions/checkout@v2

      - name: "Setup Go Env"
        uses: actions/setup-go@v2
        with:
          go-version: "{{ .Version.Golang }}"

      - name: "Setup Typescript Env"
        uses: "actions/setup-node@v2"
        with:
          node-version: "{{ .Version.Node }}"

      - name: Install Protoc Binary
        working-directory: ${{ "{{" }} runner.temp {{ "}}" }}
        run: |
          curl -LOs https://github.com/protocolbuffers/protobuf/releases/download/v{{ .Version.Protoc }}/protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          unzip protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          echo "${{ "{{" }} runner.temp {{ "}}" }}/bin/" >> $GITHUB_PATH

      - name: Install gRPC Dependencies
        working-directory: ${{ "{{" }} runner.temp {{ "}}" }}
        run: |
          curl -LOs https://github.com/grpc/grpc-web/releases/download/{{ .Version.GrpcWeb }}/protoc-gen-grpc-web-{{ .Version.GrpcWeb }}-linux-x86_64
          chmod +x protoc-gen-grpc-web-{{ .Version.GrpcWeb }}-linux-x86_64
          mv protoc-gen-grpc-web-{{ .Version.GrpcWeb }}-linux-x86_64 ${{ "{{" }} runner.temp {{ "}}" }}/bin/protoc-gen-grpc-web
          echo "${{ "{{" }} runner.temp {{ "}}" }}/bin/" >> $GITHUB_PATH

      - name: "Install Typescript Dependencies"
        run: |
          npm install prettier --global

      - name: Decrypt Private Key
        run: |
          go get github.com/xh3b4sd/red
          red decrypt -i .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa.enc -o .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa -p '${{ "{{" }} secrets.RED_GPG_PASS_{{ .Github.Organization | ToUpper }}_{{ .Github.Repository | ToUpper }} {{ "}}" }}'

      - name: Setup SSH Agent
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          mkdir -p ~/.ssh
          ssh-keyscan github.com >> ~/.ssh/known_hosts
          ssh-agent -a ${SSH_AUTH_SOCK} > /dev/null
          chmod 0600 .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa
          ssh-add .github/asset/{{ .Github.Organization }}/{{ .Github.Repository }}/id_rsa

      - name: Clone Typescript Code
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: git clone git@github.com:{{ .Github.Organization }}/{{ .Github.Repository }}.git "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"

      - name: Setup Git Config
        run: |
          cd "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
          git config user.name "${GITHUB_ACTOR}"
          git config user.email "${GITHUB_ACTOR}@users.noreply.github.com"
          git remote set-url origin git@github.com:{{ .Github.Organization }}/{{ .Github.Repository }}.git

      - name: Generate Typescript Code
        run: |
          go get github.com/xh3b4sd/pag
          rm -rf ${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/src/
          pag generate typescript -d ${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/src/

      - name: "Format Typescript Code"
        run: |
          prettier -w $(find ${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/src/ -name "*.ts" -o -name "*.tsx")

      - name: Commit And Push
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          cd "${{ "{{" }} runner.temp {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
          if [[ $(git status --porcelain) ]]; then
            git add .
            git commit -m 'update generated code'
            git push
          fi

      - name: Cleanup Build Container
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          ssh-add -D
          rm -Rf *
`
