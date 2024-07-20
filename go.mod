module github.com/xh3b4sd/workflow

go 1.22

require (
	github.com/google/go-cmp v0.6.0
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.0
	github.com/xh3b4sd/logger v0.8.1
	github.com/xh3b4sd/tracer v0.11.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.14.0 // indirect
)

retract [v0.0.0, v0.14.0]
