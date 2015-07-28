package mocks

import "github.com/GabbyyLS/data-linter/linter"
import "github.com/stretchr/testify/mock"

type ReportWriter struct {
	mock.Mock
}

func (m *ReportWriter) Start() (*linter.Report, error) {
	ret := m.Called()

	var r0 *linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.Report)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ReportWriter) Finish(_a0 *linter.Report) error {
	ret := m.Called(_a0)

	r0 := ret.Error(0)

	return r0
}
func (m *ReportWriter) Commit(_a0 *linter.Report, _a1 []*linter.Problem) error {
	ret := m.Called(_a0, _a1)

	r0 := ret.Error(0)

	return r0
}
