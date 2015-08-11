package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
	"github.com/GabbyyLS/data-linter/models"
)

// ReportWriter with empty list of problems.
func NewMockReportWriterWithNoProblems() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithNoProblems, nil)
	mockObj.On("Finish", mock.ReportWithNoProblems).Return(nil)
	mockObj.On("Commit", mock.ReportWithNoProblems, []*models.Problem{}).Return(nil)
	return mockObj
}

// ReportWriter with some problems.
func NewMockRepotWriterWithProblems() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithSomeProblems, nil)
	mockObj.On("Finish", mock.ReportWithSomeProblems).Return(nil)
	mockObj.On("Commit", mock.ReportWithSomeProblems, mock.Problems).Return(errors.New("Report has some problems"))
	return mockObj
}

// ReportWriter with an error and unfinished report.
func NewMockReportWriterWithError() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.UnfinishedReport, nil)
	mockObj.On("Finish", mock.UnfinishedReport).Return(nil)
	mockObj.On("Commit", mock.UnfinishedReport, mock.Problems).Return(errors.New("Report wasn't created due to the error/problems"))
	return mockObj
}
