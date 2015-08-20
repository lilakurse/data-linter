package linter

import (
	"github.com/GabbyyLS/data-linter/linter/checker"
)

// Iterator contains documents and provides them in fixed-sized batches.
type Iterator interface {
	Name() string
	Count() int
	Next(step int) []*checker.Checker
}

func Name(iterator Iterator) string {
	return iterator.Name()
}

func Count(iterator Iterator) int{
	return iterator.Count()
}

func Next(iterator Iterator, step int) []*checker.Checker {
	return iterator.Next(step)
}