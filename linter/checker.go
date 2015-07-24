package linter

// TODO: document
// The Checker interface expects to check files for validity(validation).
type Checker interface {
	Check() ([]*Problem, error)
}
