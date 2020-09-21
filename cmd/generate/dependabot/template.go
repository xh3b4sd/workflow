package dependabot

const templateDependabot = `#
# Do not edit. This file was generated via the workflow command line tool.
#
#     workflow generate dependabot
#

version: 2
updates:
{{ range $e := . }}
  - package-ecosystem: "{{ $e.Name }}"
    directory: "/"
    open-pull-requests-limit: 10
    reviewers:
{{- range $r := $e.Reviewers }}
      - "{{ $r }}"
{{- end }}
    schedule:
      interval: "daily"
      time: "04:00"
    target-branch: "master"
{{ end }}`

const templateGoModTidy = `#
# Do not edit. This file was generated via the workflow command line tool.
#
#     workflow generate dependabot
#

name: go-mod-tidy

on:
  push:
    branches:
      - 'dependabot/**'

jobs:
  go-mod-tidy:
    runs-on: ubuntu-latest
    steps:

      - name: Checkout Git Project
        uses: actions/checkout@v2

      - name: Fix Detached HEAD
        run: git checkout ${GITHUB_REF#refs/heads/}

      - name: Go Mod Tidy
        run: |
          rm -f go.sum
          go mod tidy

      - name: Setup Git Config
        run: |
          git config user.name "${GITHUB_ACTOR}"
          git config user.email "${GITHUB_ACTOR}@users.noreply.github.com"
          git remote set-url origin https://x-access-token:${{ secrets.GITHUB_TOKEN }}@github.com/${GITHUB_REPOSITORY}.git

      - name: Commit And Push
        run: |
          git add .
          if output=$(git status --porcelain) && [ ! -z "$output" ]; then
            git commit -m 'go mod tidy'
            git push
          fi
`
