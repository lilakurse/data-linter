package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"reflect"
	"testing"
	"github.com/GabbyyLS/data-linter/mock"
)

var (
	mockValidDoc   = Checker(mocks.NewMockValidDoc())
	mockInvalidDoc = Checker(mocks.NewMockInvalidDoc())
	mockBadDoc     = Checker(mocks.NewMockBadDoc())
)

func TestCheck(t *testing.T) {
	// Valid document
	problems, err := Check(mockValidDoc) // fix error Multiple-value Check()in single value context
	if err != nil {
		t.Error("The doc is valid, there should be no errors")
	}
	if !reflect.DeepEqual(problems, mock.EmptyProblemList) {
		t.Error("The problem list should be empty")
	}

	// Invalid document
	problems, err = Check(mockInvalidDoc)
	if err != nil {
		t.Error("The doc is invalid, but there should be no errors")
	}
	if !reflect.DeepEqual(problems, mock.Problems) {
		t.Error("The problem list should be empty")
	}

	// Bad document
	problems, err = Check(mockBadDoc)
	if err == nil {
		t.Error("There should be an error")
	}
	if problems != nil {
		t.Error("Problem list should be nil -- the doc contains an error")
	}
}
