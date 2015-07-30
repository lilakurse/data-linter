package linter

import (
	"github.com/GabbyyLS/data-linter/mocks"
	"golang.org/x/net/context"
)

type Context string

const (
	ContextCheckerMocValidDoc     = Context("CHECKER_VALIDDOC") // TODO : rethink (we have 3 DIFFERENT MockCheckers!)
	ContextCheckerMocInvalidDoc   = Context("CHECKER_INVALIDDOC")
	ContextCheckerMocBadDoc       = Context("CHECKER_BADDOC")
	ContextIteratorMockIterator   = Context("MOCK_ITERATOR")
	ContextIteratorNumberIterator = Context("NUMBER_ITERATOR")
)

var (
	ctx = context.WithValue(context.Background(), ContextIteratorMockIterator, mocks.NewMockIterator())
	// We can add ContextIteratorNumberIterator to ctx when we deal with our number iterator
)

func IteratorFromContext(ctx context.Context) Iterator {
	return ctx.Value(ContextIteratorMockIterator).(Iterator)
}

func CheckerFromContext (ctx context.Context)Checker{
	return ctx.Value().(Checker)
}