package runtime

import "runtime"

var (
	src = "https://github.com/xh3b4sd/workflow"
	sha = "n/a"
	tag = "n/a"
)

func Arc() string {
	return runtime.GOARCH
}

func Gos() string {
	return runtime.GOOS
}

func Sha() string {
	return sha
}

func Src() string {
	return src
}

func Tag() string {
	return tag
}

func Ver() string {
	return runtime.Version()
}
