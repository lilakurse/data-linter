package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
)

func NewMockChecker() *Checker {
	mockObj := new(Checker)

	// Tests for Checker().
	mockObj.On("Check", mock.ValidDoc).Return(nil)
	mockObj.On("Check", mock.InvalidDoc).Return(errors.New("invalid document"))

	return mockObj
}
