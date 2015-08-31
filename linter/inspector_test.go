package linter

import (
	"github.com/GabbyyLS/data-linter/iterators"
	"testing"
)

var (
	numGenerator iterators.Generator
)

func TestInspect(t *testing.T) {
	numGenerator.Count()
	numGenerator.Name()
	numGenerator.Next(Step)
	problem, err := Inspect(numGenerator)
	if err != nil {
		t.Error("Should be no error")
	}
	if len(problem) == 0 {
		t.Error("There should be a list of problems")
	}
	//numIterator.Count
	// Need to check count of inspected docs, count of collected problems
}
