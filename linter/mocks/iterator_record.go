package mocks

import (
	"github.com/GabbyyLS/data-linter/mock"
)

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
	mockObj.On("Next", mock.Step).Return(mock.CheckerList)
	return mockObj
}


