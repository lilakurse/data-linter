package checker

import (
	"github.com/GabbyyLS/data-linter/linter/checker/mocks"
	//"github.com/GabbyyLS/data-linter/mock"
	//"reflect"
	"testing"
)

var (
	mockValidDoc   = Checker(mocks.NewMockValidDoc())
	mockInvalidDoc = Checker(mocks.NewMockInvalidDoc())
	mockBadDoc     = Checker(mocks.NewMockBadDoc())
)

func TestCheck(t *testing.T) {
	// Valid document
	problems, err := mockValidDoc.Check()
	if err != nil {
		t.Error("The doc is valid, there should be no errors")
	}
	if len(problems) != 0 {
		t.Error("The problem list should be empty")
	}

	// Invalid document
	problems, err = mockInvalidDoc.Check()
	if err != nil {
		t.Error("The doc is invalid, but there should be no errors")
	}
	if len(problems) == 0 {
		t.Error("The problem list should be empty")
	}

	// Bad document
	problems, err = mockBadDoc.Check()
	if err == nil {
		t.Error("There should be an error")
	}
	if problems != nil {
		t.Error("Problem list should be nil -- the doc contains an error")
	}

}
