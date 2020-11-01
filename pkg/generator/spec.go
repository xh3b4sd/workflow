package generator

type Interface interface {
	Usage() ([]byte, error)
	Workflow() ([]byte, error)
}
