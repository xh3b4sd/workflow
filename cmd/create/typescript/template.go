package typescript

const templateTypescript = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

# to be up to date. Further

name: "typescript"

on: "push"

jobs:
  typescript:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Checkout Git Project"
        uses: "actions/checkout@v2"

      - name: "Setup Ts Env"
        uses: "actions/setup-node@v1"
        with:
          node-version: "{{ .Version.Node }}"

      - name: "Install Ts Dependencies"
        run: |
          npm install
          npm install prettier --global

      - name: "Check Ts Formatting"
        run: |
          prettier -c $(find ./src/ -name "*.ts" -o -name "*.tsx")

      - name: "Build Ts Project"
        run: |
          npm run build
`
