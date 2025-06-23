package dsmupdate

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "dsm-update"

on:
  push:
    branches:
      - "main"
      - "master"

jobs:
  dsm-update:
    runs-on: "ubuntu-latest"
    steps:
      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Go Env"
        uses: "actions/setup-go@v{{ .Version.SetupGo }}"
        with:
          cache: true
          go-version: "{{ .Version.Golang }}"

      - name: "Decrypt Private Key"
        run: |
          go get github.com/xh3b4sd/red
          red decrypt -i .github/asset/venturemark/flux/id_rsa.enc -o .github/asset/venturemark/flux/id_rsa -p '${{ "{{" }} secrets.RED_GPG_PASS_VENTUREMARK_FLUX {{ "}}" }}'

      - name: "Setup SSH Agent"
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          mkdir -p ~/.ssh
          ssh-keyscan github.com >> ~/.ssh/known_hosts
          ssh-agent -a ${SSH_AUTH_SOCK} > /dev/null
          chmod 0600 .github/asset/venturemark/flux/id_rsa
          ssh-add .github/asset/venturemark/flux/id_rsa

      - name: "Clone Go Code"
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: git clone git@github.com:venturemark/flux.git "${{ "{{" }} runner.temp {{ "}}" }}/venturemark/flux/"

      - name: "Setup Git Config"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/venturemark/flux/"
        run: |
          git config user.name "${GITHUB_ACTOR}"
          git config user.email "${GITHUB_ACTOR}@users.noreply.github.com"
          git remote set-url origin git@github.com:venturemark/flux.git

      - name: "Update Project Version"
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/venturemark/flux/"
        run: |
          go get github.com/xh3b4sd/dsm
          dsm update -r HelmRelease -n {{ .Repository.Name }} -k spec.values.image.tag -v ${{ "{{" }} github.sha {{ "}}" }}

      - name: "Commit And Push"
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        working-directory: "${{ "{{" }} runner.temp {{ "}}" }}/venturemark/flux/"
        run: |
          if [[ $(git status --porcelain) ]]; then
            git add .
            git commit -m 'update {{ .Repository.Name }} version'
            git push
          fi

      - name: "Cleanup Build Container"
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          ssh-add -D
          rm -Rf *
`
