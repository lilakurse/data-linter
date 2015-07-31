package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
//	"github.com/GabbyyLS/data-linter/mock"
	"testing"
	"reflect"
)

var mockValidDoc = Checker(mocks.NewMockValidDoc())

func TestCheck(t *testing.T) {
	// Valid document
	problems, err := Check(mockValidDoc) // fix error Multiple-value Check()in single value context
	if err != nil {
		t.Error("The doc is valid, there should be no error")
	}
	if !reflect.DeepEqual(problems, mocks.EmptyProblemList) {
		t.Error("The problem list should be empty")
	}



	// Invalid document
//	check, err = Check(ctx, mock.InvalidDoc)
//	if err == nil || check != nil {
//		t.Error("Invalid document")
//	}

	// Bad document
//	check, err = Check(ctx, mock.BadDoc)
//	if err != nil || check == nil {
//		t.Error("Error")
//	}
}
