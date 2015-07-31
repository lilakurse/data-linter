package mocks

import "github.com/stretchr/testify/mock"

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
func (m *Iterator) Next(step int) []*Checker {
	ret := m.Called(step)

	var r0 []*Checker
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*Checker)
	}

	return r0
}
