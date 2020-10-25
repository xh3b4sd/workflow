# workflow

Command line tool for generating github workflows.



### Workflow Generation

```
$ workflow generate -h
Generate github workflows and config files.

Usage:
  workflow generate [flags]
  workflow generate [command]

Available Commands:
  dependabot  Generate a dependabot workflow for e.g. golang and docker.
  docker      Generate a docker workflow for e.g. building and pushing docker images.
  golang      Generate a golang workflow for e.g. running tests and checking formatting.
  grpcgo      Generate a grpc workflow for golang code generation.
  grpcts      Generate a grpc workflow for typescript code generation.

Flags:
  -h, --help   help for generate

Use "workflow generate [command] --help" for more information about a command.
```
