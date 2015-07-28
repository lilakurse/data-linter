package mocks

import (
	"github.com/GabbyyLS/data-linter/linter"
	"github.com/stretchr/testify/mock"
)

type Iterator struct {
	mock.Mock
}

func (m *Iterator) Name() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *Iterator) Count() int64 {
	ret := m.Called()

	r0 := ret.Get(0).(int64)

	return r0
}
func (m *Iterator) Next(Step int) []*linter.Checker {
	ret := m.Called(Step)

	var r0 []*linter.Checker
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*linter.Checker)
	}

	return r0
}
