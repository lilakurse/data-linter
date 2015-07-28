package linter

import (
	"github.com/GabbyyLS/data-linter/mocks"
	"golang.org/x/net/context"
)

type Context string

const (
	ContextChecker = Context("CHECKER")
)

// Context creation.
var (
	ctx = context.WithValue(context.Background(), ContextChecker, mocks.NewMockChecker())
)

func FromContext(ctx context.Context) Checker {
	return ctx.Value(ContextChecker).(Checker)
}
