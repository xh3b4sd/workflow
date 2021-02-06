# workflow

Command line tool for generating github workflows. Certain optinionated design
decisions have been made which limits the workflow generation and other tooling
to the github ecosystem. All commands generating workflows must be executed
within the root directory of the repository the desired workflows should be
generated for. To make most of the workflows work the `red` command line tool is
necessary. For more information see https://github.com/xh3b4sd/red.



### Create Workflows

```
$ workflow create -h
Create github workflows and config files.

Usage:
  workflow create [flags]
  workflow create [command]

Available Commands:
  cfmtest     Create a conformance workflow for e.g. running tests.
  dependabot  Create a dependabot workflow for e.g. golang and docker.
  dockergo    Create a docker workflow for building and pushing docker images of golang apps.
  dockerts    Create a docker workflow for building and pushing docker images of typescript apps.
  dsmupdate   Create a mutating workflow for e.g. app versions.
  dsmverify   Create a validation workflow for e.g. checking consistency.
  golang      Create a golang workflow for e.g. running tests and checking formatting.
  grpcgo      Create a grpc workflow for golang code generation.
  grpcts      Create a grpc workflow for typescript code generation.
  npm         Create a npm workflow for e.g. building and publishing npm packages.
  redigo      Create a golang workflow for e.g. running redis conformance tests.
  typescript  Create a typescript workflow for e.g. building and formatting typescript code.

Flags:
  -h, --help   help for create

Use "workflow create [command] --help" for more information about a command.
```



```
$ workflow create dependabot -h
Create a dependabot workflow for e.g. golang and docker.

Usage:
  workflow create dependabot [flags]

Flags:
  -b, --branch string           Dependabort target branch to merge pull requests into. (default "main")
  -h, --help                    help for dependabot
  -r, --reviewers strings       Reviewers assigned to dependabot PRs, e.g. xh3b4sd. Works with github usernames and teams.
  -g, --version-golang string   Golang version to use in, e.g. workflow files. (default "1.15.2")
```



```
$ workflow create grpcgo -h
Create a grpc workflow for golang code generation. The workflow generated
here works in a setup of two Github repositories. Call them apischema and
gocode. The workflow generated with the following command is added to the
apischema repository.

    workflow create grpcgo -o xh3b4sd -r gocode

In order to make the workflow function correctly a deploy key is generated
and distributed as follows. The public and private key files are added to the
apischema repository. The public key is added as deploy key with write access
to the gocode repository. Deploy keys and GPG password are generated with the
following command.

    red generate keys -d .github/asset/xh3b4sd/gocode

Generating the deploy keys also generates a GPG password which is used to
decrypt the encrypted private key within the build container of the workflow.
The GPG password needs to be added to the apischema secrets using the
following name.

    RED_GPG_PASS_XH3B4SD_GOCODE

Usage:
  workflow create grpcgo [flags]

Flags:
  -o, --github-organization string   Github organization to generate code for.
  -r, --github-repository string     Github repository to generate code for.
  -h, --help                         help for grpcgo
  -s, --silent                       Silence the command output to not give feedback.
  -g, --version-golang string        Golang version to use in, e.g. workflow files. (default "1.15.2")
  -p, --version-protoc string        Protoc version to use in, e.g. workflow files. (default "3.13.0")
```



```
$ workflow create grpcts -h
Create a grpc workflow for typescript code generation. The workflow generated
here works in a setup of two Github repositories. Call them apischema and
tscode. The workflow generated with the following command is added to the
apischema repository.

    workflow create grpcts -o xh3b4sd -r tscode

In order to make the workflow function correctly a deploy key is generated
and distributed as follows. The public key and the encrypted private key
files are added to the apischema repository. The public key is added as
deploy key with write access to the tscode repository. The GPG password and
the deploy keys are generated with the red command line tool. For more
information see https://github.com/xh3b4sd/red.

    red generate keys -d .github/asset/xh3b4sd/tscode

Generating the deploy keys also generates a GPG password which is used to
decrypt the encrypted private key within the build container of the workflow.
Considering the example described above, the GPG password needs to be added
to the apischema Github repository secrets using the following secret name.

    RED_GPG_PASS_XH3B4SD_TSCODE

Usage:
  workflow create grpcts [flags]

Flags:
  -o, --github-organization string   Github organization to generate code for.
  -r, --github-repository string     Github repository to generate code for.
  -h, --help                         help for grpcts
  -s, --silent                       Silence the command output to not give feedback.
  -g, --version-golang string        Golang version to use in, e.g. workflow files. (default "1.15.2")
  -w, --version-grpc-web string      Grpc Web version to use in, e.g. workflow files. (default "1.2.1")
  -p, --version-protoc string        Protoc version to use in, e.g. workflow files. (default "3.13.0")
```



### Update Workflows


```
$ workflow update all -h
Update all github workflows to the latest version. When creating a new
workflow file the original command instruction in form of os.Args is written
to the header of the workflow file. A typical workflow file header looks like
the following.

    #
    # Do not edit. This file was generated via the "workflow" command line tool.
    # More information about the tool can be found at github.com/xh3b4sd/workflow.
    #
    #     workflow create dependabot -r xh3b4sd
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
