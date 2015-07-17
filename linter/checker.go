package linter

type Checker interface {
	Check() ([]*Problem, error)
}
