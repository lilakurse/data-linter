package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/linter"
)

func NewValidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return([]*linter.Problem{}, nil)
	return mockObj
}

// TODO: implement InvalidDoc as a MockChecker
func NewInvalidDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(mock.Problem, nil)
	return  mockObj
}

func NewBadDoc() *Checker {
	mockObj := new(Checker)
	mockObj.On("Check").Return(nil, errors.New("Error: bad document"))
	return mockObj
}
