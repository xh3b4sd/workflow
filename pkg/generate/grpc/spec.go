package grpc

type Interface interface {
	Generate() ([]byte, error)
}
