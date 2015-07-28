package mocks

import "github.com/GabbyyLS/data-linter/linter"
import "github.com/stretchr/testify/mock"

import "time"

type ReportReader struct {
	mock.Mock
}

func (m *ReportReader) GetAllReports() []*linter.Report {
	ret := m.Called()

	var r0 []*linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*linter.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByName(Name string) *linter.Report {
	ret := m.Called(Name)

	var r0 *linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByCreationTime(time *time.Time) *linter.Report {
	ret := m.Called(time)

	var r0 *linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByUpdateTime(time *time.Time) *linter.Report {
	ret := m.Called(time)

	var r0 *linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.Report)
	}

	return r0
}
func (m *ReportReader) GetReportByCommitTime(time *time.Time) *linter.Report {
	ret := m.Called(time)

	var r0 *linter.Report
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.Report)
	}

	return r0
}
func (m *ReportReader) TotalReportsCount() *linter.ReportsCount {
	ret := m.Called()

	var r0 *linter.ReportsCount
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*linter.ReportsCount)
	}

	return r0
}
