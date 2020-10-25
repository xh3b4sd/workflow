package generate

type Interface interface {
	Generate() ([]byte, error)
}
