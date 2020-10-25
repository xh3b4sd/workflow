package grpc

import (
	"flag"
	"strconv"
)

var update = flag.Bool("update", false, "update .golden files")

func fileName(i int) string {
	return "case-" + strconv.Itoa(i) + ".golden"
}
