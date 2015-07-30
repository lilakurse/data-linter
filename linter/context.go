package linter

import (
	"github.com/GabbyyLS/data-linter/mocks"
	"golang.org/x/net/context"
)

type Context string

const (
	ContextCheckerMocValidDoc     = Context("MOCK_CHECKER_VALIDDOC")
	ContextCheckerMocInvalidDoc   = Context("MOCK_CHECKER_INVALIDDOC")
	ContextCheckerMocBadDoc       = Context("MOCK_CHECKER_BADDOC")
	ContextIteratorMockIterator   = Context("MOCK_ITERATOR")
	ContextIteratorNumberIterator = Context("NUMBER_ITERATOR")
)

var (
	ctx = context.WithValue(context.Background(), ContextIteratorMockIterator, mocks.NewMockIterator())
	// TODO: add 3 mock checkers
	// We can add ContextIteratorNumberIterator to ctx when we deal with our number iterator
)

func IteratorFromContext(ctx context.Context) Iterator {
	// TODO: + add an n-case switch so we don't need n separate methods
	return ctx.Value(ContextIteratorMockIterator).(Iterator)
}

func CheckerFromContext (ctx context.Context)Checker{
	// TODO: add a 3-case switch so we don't need 3 separate methods
	return ctx.Value().(Checker)
}