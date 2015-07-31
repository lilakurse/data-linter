package mocks

import "github.com/stretchr/testify/mock"

import "github.com/GabbyyLS/data-linter/models"

type ReportWriter struct {
	mock.Mock
}

func (m *ReportWriter) Start() (*models.Report, error) {
	ret := m.Called()

	var r0 *models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Report)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ReportWriter) Finish(_a0 *models.Report) error {
	ret := m.Called(_a0)

	r0 := ret.Error(0)

	return r0
}
func (m *ReportWriter) Commit(_a0 *models.Report, _a1 []*models.Problem) error {
	ret := m.Called(_a0, _a1)

	r0 := ret.Error(0)

	return r0
}
