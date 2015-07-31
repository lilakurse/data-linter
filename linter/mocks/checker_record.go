package mocks

import (
	"errors"
)

func NewMockValidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(EmptyProblemList, nil)
	return mockObj
}

func NewMockInvalidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(Problems, nil)
	return mockObj
}

func NewMockBadDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(nil, errors.New("Error: bad document"))
	return mockObj
}
