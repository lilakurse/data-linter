package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
	"github.com/GabbyyLS/data-linter/models"
)

// Writer with empty list of problems.
func NewMockReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithNoProblems, nil)
	mockObj.On("Finish", mock.ReportWithNoProblems).Return(nil)
	mockObj.On("Commit", mock.ReportWithNoProblems, []*models.Problem{}).Return(nil)
	return mockObj
}

// Writer with some problems.
func NewMockWithProblemRepotWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithSomeProblems, nil)
	mockObj.On("Finish", mock.ReportWithSomeProblems).Return(nil)
	mockObj.On("Commit", mock.ReportWithSomeProblems, mock.Problems).Return(errors.New("Report has some problems"))
	return mockObj
}

// Writer with error and no report.
func NewMockErrorReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.UnfinishedReport, nil)
	mockObj.On("Finish", mock.UnfinishedReport).Return(nil)
	mockObj.On("Commit", mock.UnfinishedReport, mock.Problems).Return(errors.New("Report wasn't created due to the error/problems"))
	return mockObj
}
