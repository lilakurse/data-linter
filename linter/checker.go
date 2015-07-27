package linter

// The Checker runs various checks for validity on itself and returns a list of problems, if any.
type Checker interface {
	Check() ([]*Problem, error)
}
