module github.com/xh3b4sd/workflow

go 1.21

require (
	github.com/google/go-cmp v0.5.9
	github.com/spf13/afero v1.10.0
	github.com/spf13/cobra v1.7.0
	github.com/xh3b4sd/logger v0.7.1
	github.com/xh3b4sd/tracer v0.10.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.4.0 // indirect
)

retract [v0.0.0, v0.14.0]
