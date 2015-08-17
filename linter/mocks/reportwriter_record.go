package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
	"github.com/GabbyyLS/data-linter/models"
)

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
	mockObj.On("Commit", mock.ReportWithSomeProblems, mock.Problems).Return(nil)
	mockObj.On("Finish", mock.ReportWithSomeProblems).Return(nil)
	return mockObj
}

func NewMockReportWriterWithError() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.ReportWithError, nil)
	mockObj.On("Commit", mock.ReportWithError, mock.Problems).Return(errors.New("Report was not finshed due to error(s)"))
	mockObj.On("Finish", mock.ReportWithError).Return(errors.New("Report was not finshed due to error(s)"))
	return mockObj
}
