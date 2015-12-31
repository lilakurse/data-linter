package linter

import (
	"github.com/lilakurse/data-linter/iterators"
	"github.com/lilakurse/data-linter/reporters/mem"
	"testing"
	"github.com/lilakurse/data-linter/reporters/multi"
)

var (
	numGenerator  = iterators.NewGenerator("test number generator",100)
)

func TestInspect(t *testing.T) {
	reporterOne := mem.New()
	reporterSec := mem.New()
	reporter := multi.New(reporterOne,reporterSec)
	inspector := NewInspector(reporter)
	err := inspector.Inspect(numGenerator)
	if err != nil {
		t.Error("Should be no error")
	}
	//numIterator.Count
	// Need to check count of inspected docs, count of collected problems
}
