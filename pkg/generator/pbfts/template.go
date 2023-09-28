package pbfts

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "pbf-ts"

on:
  push:
    branches:
      - "main"
      - "master"
    paths:
      - "**.proto"
  workflow_dispatch:

jobs:
  pbf-ts:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Typescript Env"
        uses: "actions/setup-node@v{{ .Version.SetupNode }}"
        with:
          node-version: "{{ .Version.Node }}"

      - name: "Install Protoc Binary"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}"
        run: |
          curl -LOs https://github.com/protocolbuffers/protobuf/releases/download/v{{ .Version.Protoc }}/protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          unzip protoc-{{ .Version.Protoc }}-linux-x86_64.zip
          echo "${{ "{{" }} runner.temp {{ "}}" }}/bin" >> $GITHUB_PATH

      - name: "Install Typescript Dependencies"
        run: |
          npm install prettier --global
          npm install @protobuf-ts/plugin --global

      - name: "Clone Typescript Code"
        run: |
          git clone https://github.com/{{ .Github.Organization }}/{{ .Github.Repository }}.git "${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"

      - name: "Generate Typescript Code"
        run: |
          inp="./pbf"
          out=${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/src
          tmp="./tmp"

          rm -rf $out

          for x in $(ls $inp); do
            if [ -n "$(ls $inp/$x)" ]; then
              mkdir -p $out/$x
              mkdir -p $tmp/$x

              lin=()
              for y in $(ls -F $inp/$x); do
                lin+=($inp/$x/$y)
              done

              npx protoc --proto_path=. --ts_out=$tmp/$x --ts_opt=output_javascript ${lin[@]}
              mv $tmp/$x/$inp/$x/* $out/$x
              rm -rf $tmp/$x
            fi
          done

          rm -rf $tmp

      - name: "Format Typescript Code"
        run: |
          prettier -w $(find ${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/src/ -name "*.ts" -o -name "*.tsx")

      - name: "Commit And Push"
        uses: "cpina/github-action-push-to-another-repository@v1.7.2"
        env:
          SSH_DEPLOY_KEY: "${{ "{{" }} secrets.SSH_DEPLOY_KEY_{{ .Github.Repository | ToUpper }} {{ "}}" }}"
        with:
          source-directory: "${{ "{{" }} github.sha {{ "}}" }}/{{ .Github.Organization }}/{{ .Github.Repository }}/"
          destination-github-username: "{{ .Github.Organization }}"
          destination-repository-name: "{{ .Github.Repository }}"
          commit-message: "update generated code"
          target-branch: "main"
`
