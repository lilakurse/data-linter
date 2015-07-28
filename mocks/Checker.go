package mocks

import (
	"github.com/GabbyyLS/data-linter/linter"
	"github.com/stretchr/testify/mock"
)

type Checker struct {
	mock.Mock
}

func (m *Checker) Check() ([]*linter.Problem, error) {
	ret := m.Called()

	var r0 []*linter.Problem
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*linter.Problem)
	}
	r1 := ret.Error(1)

	return r0, r1
}
