package parser

type Interface interface {
	// Parse tries to find all commands that got used to create github
	// workflows. The command used during workflow creation is written to the
	// header of the workflow file. Parse tries to find all workflow files and
	// returns the list of commands it found in the form of os.Args. That way
	// updating existing workflow files can be implemented by dynamically.
	Parse() ([][]string, error)
}
