# workflow

Command line tool for generating github workflows. Certain optinionated design
decisions have been made which limits the workflow generation and other tooling
to the github ecosystem. All commands generating workflows must be executed
within the root directory of the repository the desired workflows should be
generated for. To make most of the workflows work the `red` command line tool is
necessary. For more information see https://github.com/xh3b4sd/red.



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



```
$ workflow generate dependabot -h
Generate a dependabot workflow for e.g. golang and docker. The dependabot
workflow for golang includes a separate action of executing "go mod tidy" due
to some dependabot limitations. This limitation requires automated fixing and
therefore dedicated push access to the configured repository. The explicitly
authorized push access is necessary in order to trigger another execution of
the "go build" workflow after we fixed the go.mod and go.sum files. Due to
these implementation details we setup deploy keys in each repository via the
"red" command line tool. More information about the tool can be found at
github.com/xh3b4sd/red. For each repository using the dependabot workflow the
following command must be used in order to generate deploy keys and the
associated GPG password.

    red generate keys -d .github/asset

The GPG encrypted private key must be put into the ".github/asset" directory.
During each build the GPG encrypted private key is decrypted within the build
container and used to setup the local SSH agent for the explicit push access.

    .github/asset/id_rsa.enc

The plain text public key must be added as deploy key with write access to
the configured repository. During builds of dependabot pull requests the
configured deploy key verifies that the configured private key is allowed to
push changes during builds.

    .github/asset/id_rsa.pub

During builds, a password is required for the decryption of the GPG encrypted
private key. This password gets generated together with the RSA public and
private key as shown above. The password must be set as secret to the
configured repository. The secret name must be as follows.

    RED_GPG_PASS

Usage:
  workflow generate dependabot [flags]

Flags:
  -h, --help                    help for dependabot
  -r, --reviewers strings       Reviewers assigned to dependabot PRs, e.g. xh3b4sd. Works with github usernames and teams.
  -g, --version-golang string   Golang version to use in, e.g. workflow files. (default "1.15.2")
```



```
$ workflow generate grpcgo -h
Generate a grpc workflow for golang code generation. The workflow generated
here works in a setup of two Github repositories. Call them apischema and
gocode. The workflow generated with the following command is added to the
apischema repository.

    workflow generate grpcgo -o xh3b4sd -r gocode

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
  workflow generate grpcgo [flags]

Flags:
  -o, --github-organization string   Github organization to generate code for.
  -r, --github-repository string     Github repository to generate code for.
  -h, --help                         help for grpcgo
  -g, --version-golang string        Golang version to use in, e.g. workflow files. (default "1.15.2")
  -p, --version-protoc string        Protoc version to use in, e.g. workflow files. (default "3.13.0")
```



```
$ workflow generate grpcts -h
Generate a grpc workflow for typescript code generation. The workflow
generated here works in a setup of two Github repositories. Call them
apischema and tscode. The workflow generated with the following command is
added to the apischema repository.

    workflow generate grpcts -o xh3b4sd -r tscode

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
  workflow generate grpcts [flags]

Flags:
  -o, --github-organization string   Github organization to generate code for.
  -r, --github-repository string     Github repository to generate code for.
  -h, --help                         help for grpcts
  -g, --version-golang string        Golang version to use in, e.g. workflow files. (default "1.15.2")
  -w, --version-grpc-web string      Grpc Web version to use in, e.g. workflow files. (default "1.2.1")
  -p, --version-protoc string        Protoc version to use in, e.g. workflow files. (default "3.13.0")
```
