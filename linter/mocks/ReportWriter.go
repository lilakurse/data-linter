package mocks

import "github.com/stretchr/testify/mock"

import "github.com/lilakurse/data-linter/models"

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
func (m *ReportWriter) Finish(r *models.Report) error {
	ret := m.Called(r)

	r0 := ret.Error(0)

	return r0
}
func (m *ReportWriter) Commit(r *models.Report, p []*models.Problem) error {
	ret := m.Called(r, p)

	r0 := ret.Error(0)

	return r0
}
