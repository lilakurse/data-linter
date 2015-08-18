package mocks

import "github.com/stretchr/testify/mock"

import "github.com/GabbyyLS/data-linter/models"

type Checker struct {
	mock.Mock
}

func (m *Checker) Check() ([]*models.Problem, error) {
	ret := m.Called()

	var r0 []*models.Problem
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*models.Problem)
	}
	r1 := ret.Error(1)

	return r0, r1
}
