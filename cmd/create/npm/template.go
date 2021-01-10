package npm

const templateWorkflow = `#
# Do not edit. This file was generated via the "workflow" command line tool.
# More information about the tool can be found at github.com/xh3b4sd/workflow.
#
#     {{ .Command }}
#
# Note that this workflow has several requirements in order to function
# correctly throughout the development process. See the desired sequence of
# steps below.
#
#     1. A developer bumps the semver version within the package.json.
#     2. A developer runs "npm install" to update package-lock.json.
#     3. A developer creates and merges a pull request for the version changes.
#     4. A developer creates a new github release according to step 1.
#     5. A Github action builds and publishes the new npm package into the Github package registry.
#

name: "npm-publish"

on:
  release:
    types:
      - "created"

jobs:
  npm-publish:
    runs-on: "ubuntu-latest"
    steps:

      - name: "Checkout Git Project"
        uses: "actions/checkout@v2"

      - name: "Setup Ts Env"
        uses: "actions/setup-node@v2"
        with:
          node-version: "{{ .Version.Node }}"
          registry-url: "https://npm.pkg.github.com"

      - name: "Install Ts Dependencies"
        run: |
          npm install

      - name: "Build Ts Project"
        run: |
          npm run build

      - name: "Publish NPM Package"
        env:
          NODE_AUTH_TOKEN: "${{ "{{" }} secrets.GITHUB_TOKEN {{ "}}" }}"
        run: |
          npm publish
`
