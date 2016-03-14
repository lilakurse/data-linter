package elastic

import (
	"testing"
	"github.com/lilakurse/data-linter/linter"
	"github.com/lilakurse/data-linter/iterators"
)

func TestReportElastic(t *testing.T) {
	elWReporter,err := New()
	if err != nil {
		t.Error("create new elastic reporter",err)
	}
	inspector := linter.NewInspector(elWReporter)
	err = inspector.Inspect(iterators.NewGenerator("test number generator",10000))
	if err != nil {
		t.Error("Should be no error")
	}
}
