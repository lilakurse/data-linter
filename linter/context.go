package linter

import (
	"github.com/GabbyyLS/data-linter/mocks"
	"golang.org/x/net/context"
)

type Context string

const (
	ContextChecker = Context("CHECKER") // TODO : rethink (we have 3 DIFFERENT MockCheckers!)
	ContextIteratorMockIterator = Context("MOCK_ITERATOR")
	ContextIteratorNumberIterator = Context("NUMBER_ITERATOR")
)

var (
	ctx = context.WithValue(context.Background(), ContextIteratorMockIterator, mocks.NewMockIterator())
	// We can add ContextIteratorNumberIterator to ctx when we deal with our number iterator
)

func FromContext(ctx context.Context) Iterator {
	return ctx.Value(ContextIteratorMockIterator).(Iterator)
}