package linter

import (
	"golang.org/x/net/context"
	"github.com/GabbyyLS/data-linter/mock"
)

// The Checker runs various checks for validity on itself and returns a list of problems, if any.
type Checker interface {
	Check() ([]*Problem, error)
}

func Check(ctx context.Context,doc *mock.Document) ([]*Problem, error){
	return FromContext(ctx).Check(doc)
}
