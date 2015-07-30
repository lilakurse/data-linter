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
	ContextIteratorMockValid      = Context("MOCK_ITERATOR_VALID")
	ContextIteratorMockInvalid    = Context("MOCK_ITERATOR_INVALID")
	ContextIteratorMockBad        = Context("MOCK_ITERATOR_BAD")
)

var (
	ctx = context.WithValue(context.Background(), ContextIteratorMockIterator, mocks.NewMockIterator())
	ctx = context.WithValue(context.Background(), ContextIteratorMockValid, mocks.NewMockFilledIterator())
	ctx = context.WithValue(context.Background(), ContextIteratorMockInvalid, mocks.NewMockEmptyIterator())
	ctx = context.WithValue(context.Background(), ContextIteratorMockIterator, mocks.NewMockWithErrorIterator())
	// TODO: add 3 mock checkers
	// We can add ContextIteratorNumberIterator to ctx when we deal with our number iterator
	ctx = context.WithValue(context.Background(), ContextCheckerMocValidDoc, mocks.NewMockValidDoc())
	ctx = context.WithValue(context.Background(), ContextCheckerMocInvalidDoc, mocks.NewMockInvalidDoc())
	ctx = context.WithValue(context.Background(), ContextCheckerMocBadDoc, mocks.NewMockBadDoc())
	// Should we create context for checkers here with name ctx, or we should writ different names?
)

func IteratorFromContext(ctx context.Context) Iterator {
	// TODO: + add an n-case switch so we don't need n separate methods
	switch ctx{
		case 1:
			if context==ctx.Value(ContextIteratorMockValid){		// Or just ctx.Value()
				return ctx.Value(ContextIteratorMockValid).(Iterator)
			}
		case 2:
			if context==ctx.Value(ContextIteratorMockInvalid){
				return ctx.Value(ContextIteratorMockInvalid).(Iterator)
			}
		case 3:
			if context ==ctx.Value(ContextIteratorMockBad){
				return ctx.Value(ContextIteratorMockBad).(Iterator)
			}
	}
	return ctx.Value(ContextIteratorMockIterator).(Iterator)
}

func CheckerFromContext (ctx context.Context)Checker{
	// TODO: add a 3-case switch so we don't need 3 separate methods
	switch ctx{
	case 1:
		if ctx.Value(ContextCheckerMocValidDoc) {    // How to check that it is the correct ctx
			return ctx.Value(ContextCheckerMocValidDoc).(Checker)
		}
	case 2:
		if ctx.Value(ContextCheckerMocInvalidDoc) {
			return ctx.Value(ContextCheckerMocInvalidDoc).(Checker)
		}
	case 3:
		if ctx.Value(ContextCheckerMocBadDoc) {
			return ctx.Value(ContextCheckerMocBadDoc).(Checker)
		}
	}
	return // should we return anything
}
