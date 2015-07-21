package linter

// TODO: document
type Checker interface {
	Check() ([]*Problem, error)
}
