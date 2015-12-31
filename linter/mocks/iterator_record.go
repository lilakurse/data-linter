package mocks

import (
	"github.com/lilakurse/data-linter/mock"
)
import "github.com/lilakurse/data-linter/linter/checker"


func NewMockIteratorWithoutCheckers() *Iterator {
	mockObj := new(Iterator)
	mockObj.On("Name").Return("Empty mock iterator")
	mockObj.On("Count").Return(mock.EmptyCount)
	mockObj.On("Next", mock.Step).Return(nil)
	return mockObj
}

func NewMockIteratorWithCheckers() *Iterator {
	mockObj := new(Iterator)
	mockObj.On("Name").Return("Non-empty mock iterator")
	mockObj.On("Count").Return(mock.NonEmptyCount)
	mockObj.On("Next", mock.Step).Return(convertMockCheckerList(mock.CheckerList))
	return mockObj
}

// This conversion is needed to avoid import cycle:
// import "github.com/lilakurse/data-linter/linter/checker" in
// /mock/checker.go
func convertMockCheckerList(mockCheckerList []*mock.MockChecker) []*checker.Checker {
	finalList := []*checker.Checker{}

	for _, mockChecker := range mockCheckerList {
		checker := checker.Checker(mockChecker)
		finalList = append(finalList, &checker)
	}

	return finalList
}
