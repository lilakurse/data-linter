package linter

type Checker interface {
	Check() error
}
