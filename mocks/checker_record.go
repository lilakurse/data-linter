package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/linter"
	"github.com/GabbyyLS/data-linter/mock"
)

func NewMockChecker() *Checker {
	mockObj := new(Checker)

	// Tests for Checker().
	mockObj.On("Check", mock.ValidDoc).Return([]*linter.Problem{}, nil)
	mockObj.On("Check", mock.InvalidDoc).Return([]*linter.Problem{mock.Problem}, nil)
	mockObj.On("Chect", mock.BadDoc).Return(nil, errors.New("Error: bad document"))

	return mockObj
}
