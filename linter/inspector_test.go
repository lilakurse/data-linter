package linter

import (
	"github.com/GabbyyLS/data-linter/iterators"
	"testing"
)

var (
	numIterator  Iterator
	numGenerator iterators.Generator
)

func TestInspect(t *testing.T) {
	numIterator.Count()
	numIterator.Name()
	numIterator.Next(Step)
	problem, err := Inspect(numIterator)
	if err != nil {
		t.Error("Should be no error")
	}
	if len(problem) == 0 {
		t.Error("There should be a list of problems")
	}
	//numIterator.Count
	// Need to check count of inspected docs, count of collected problems
}
