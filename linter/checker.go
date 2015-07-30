package linter

import (
	"golang.org/x/net/context"
)

// The Checker runs various checks for validity on itself and returns a list of problems, if any.
type Checker interface {
	Check() ([]*Problem, error)
}

func Check(ctx context.Context) ([]*Problem, error){
	return CheckerFromContext(ctx).Check() // TODO: see context.go
}

