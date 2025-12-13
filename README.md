# workflow

Command line tool for generating github workflows. Certain optinionated design
decisions have been made which limits the workflow generation and other tooling
to the github ecosystem. All commands generating workflows must be executed
within the root directory of the repository the desired workflows should be
generated for.

### Create Workflows

```
workflow create -h
```

```
Create github workflows and config files.

Usage:
  workflow create [flags]
  workflow create [command]

Available Commands:
  dependabot  Create a dependabot workflow for e.g. golang and docker.
  dockergo    Create a docker workflow for building and pushing docker images of golang apps.
  dockerts    Create a docker workflow for building and pushing docker images of typescript apps.
  golang      Create a golang workflow for e.g. running tests and checking formatting.
  npm         Create a npm workflow for e.g. building and publishing npm packages.
  pbfgo       Create a protocol buffer workflow for golang code generation.
  pbflint     Create a protocol buffer workflow for schema validation.
  pbfts       Create a protocol buffer workflow for typescript code generation.
  redigo      Create a redis workflow for e.g. running conformance tests.
  releasego   Create a golang workflow for e.g. uploading cross compiled release assets.
  typescript  Create a typescript workflow for e.g. building and formatting typescript code.
  valkey      Create a valkey workflow for e.g. running integration tests.

Flags:
  -h, --help   help for create

Use "workflow create [command] --help" for more information about a command.
```

```
workflow create dependabot -h
```

```
Create a dependabot workflow for e.g. golang and docker.

Usage:
  workflow create dependabot [flags]

Flags:
  -b, --branch string           Dependabort target branch to merge pull requests into. (default "main")
  -h, --help                    help for dependabot
  -r, --reviewers strings       Reviewers assigned to dependabot PRs, e.g. @xh3b4sd.
  -g, --version-golang string   Golang version to use in, e.g. workflow files. (default "1.24.0")
```

### Update Workflows

```
workflow update all -h
```

```
Update all github workflows to the latest version. When creating a new
workflow file the original command instruction in form of os.Args is written
to the header of the workflow file. A typical workflow file header looks like
the following.

    #
    # Do not edit. This file was generated via the "workflow" command line tool.
    # More information about the tool can be found at github.com/xh3b4sd/workflow.
    #
    #     workflow create dependabot -r @xh3b4sd
    #

This information of the executable command is used to make workflow updates
reproducible. All workflow files within the github specific workflow
directory are inspected when collecting command instructions. Once all
commands are known they are executed dynamically while new behaviour is
applied.

    .github/workflows/

Usage:
  workflow update all [flags]

Flags:
  -h, --help   help for all
```
