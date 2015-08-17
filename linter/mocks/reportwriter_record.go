package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
	"github.com/GabbyyLS/data-linter/models"
)

// Signatures:
//Start() (*Report, error)
//Finish(*Report) error
//Commit(*Report, []*Problem) error

// ReportWriter with an empty list of problems.
func NewMockReportWriterWithNoProblems() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithNoProblems, nil)
	mockObj.On("Commit", mock.ReportWithNoProblems, []*models.Problem{}).Return(nil)
	mockObj.On("Finish", mock.ReportWithNoProblems).Return(nil)
	return mockObj
}

// ReportWriter with some problems.
func NewMockReportWithSomeProblems() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithSomeProblems, nil)
	mockObj.On("Commit", mock.ReportWithSomeProblems, mock.Problems).Return(errors.New("Report has some problems"))
	mockObj.On("Finish", mock.ReportWithSomeProblems).Return(nil)
	return mockObj
}

// ReportWriter with an error and unfinished report.
func NewMockReportWriterWithError() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithError, nil)
	mockObj.On("Commit", mock.ReportWithError, mock.Problems).Return(errors.New("Report wasn't created due to the error/problems"))
	mockObj.On("Finish", mock.ReportWithError).Return(nil)
	return mockObj
}
