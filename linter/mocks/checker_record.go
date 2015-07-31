package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
)

func NewMockValidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(mock.EmptyProblemList, nil)
	return mockObj
}

func NewMockInvalidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(mock.Problems, nil)
	return mockObj
}

func NewMockBadDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(nil, errors.New("Error: bad document"))
	return mockObj
}
