package mocks

import "github.com/stretchr/testify/mock"

import "github.com/lilakurse/data-linter/linter/checker"

type Iterator struct {
	mock.Mock
}

func (m *Iterator) Name() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *Iterator) Count() int {
	ret := m.Called()

	r0 := ret.Get(0).(int)

	return r0
}
func (m *Iterator) Next(step int) []checker.Checker {
	ret := m.Called(step)

	var r0 []checker.Checker
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]checker.Checker)
	}

	return r0
}
