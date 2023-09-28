package generator

type Interface interface {
	Workflow() ([]byte, error)
}
