package mocks

import "github.com/stretchr/testify/mock"

import "time"
import "github.com/lilakurse/data-linter/models"

type ReportReader struct {
	mock.Mock
}

func (m *ReportReader) GetAllReports() []*models.Report {
	ret := m.Called()

	var r0 []*models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*models.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByName(Name string) *models.Report {
	ret := m.Called(Name)

	var r0 *models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByCreationTime(time *time.Time) *models.Report {
	ret := m.Called(time)

	var r0 *models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByUpdateTime(time *time.Time) *models.Report {
	ret := m.Called(time)

	var r0 *models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByCommitTime(time *time.Time) *models.Report {
	ret := m.Called(time)

	var r0 *models.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Report)
	}

	return r0
}
func (m *ReportReader) TotalReportsCount() *models.ReportsCount {
	ret := m.Called()

	var r0 *models.ReportsCount
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.ReportsCount)
	}

	return r0
}
