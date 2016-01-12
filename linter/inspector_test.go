package linter

import (
	"github.com/lilakurse/data-linter/iterators"
	"github.com/lilakurse/data-linter/reporters/mem"
	"testing"
)

var (
	numGenerator  = iterators.NewGenerator("test number generator",100)
)

func TestInspect(t *testing.T) {
	reporter := mem.New()
	inspector := NewInspector(reporter)
	err := inspector.Inspect(numGenerator)
	if err != nil {
		t.Error("Should be no error")
	}
	//numIterator.Count
	// Need to check count of inspected docs, count of collected problems
}
