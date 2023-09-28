package typescript

const templateTypescript = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#

name: "typescript"

on: "push"

jobs:
  typescript:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Setup Git Project"
        uses: "actions/checkout@v{{ .Version.Checkout }}"

      - name: "Setup Typescript Env"
        uses: "actions/setup-node@v{{ .Version.SetupNode }}"
        with:
          node-version: "{{ .Version.Node }}"

      - name: "Install Typescript Dependencies"
        run: |
          npm install
          npm install prettier --global

      - name: "Check Git Project"
        run: |
          if [[ $(git status --porcelain) ]]; then
            echo "found dirty files"
            exit 1
          fi

      - name: "Check Typescript Formatting"
        run: |
          prettier -c $(find ./src -name "*.ts" -o -name "*.tsx")

      - name: "Check Typescript Tests"
        run: |
          npm run test --if-present

      - name: "Check Typescript Coverage"
        run: |
          npm run cover --if-present

      - name: "Build Typescript Project"
        run: |
          npm run build --if-present
`
