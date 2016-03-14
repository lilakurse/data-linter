package mocks

import (
	"errors"
	"github.com/lilakurse/data-linter/mock"
)

func NewMockValidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(mock.EmptyProblemList, nil) // TODO: nil, nil or leave as is?
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
