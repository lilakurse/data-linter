package mocks

import (
	"errors"
	"github.com/GabbyyLS/data-linter/mock"
	"github.com/GabbyyLS/data-linter/models"
)

// Writer with empty list of problems.
func NewMockReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.Report, nil)
	mockObj.On("Finish", mock.Report).Return(nil)
	mockObj.On("Commit", mock.Report, []*models.Problem{}).Return(nil)
	return mockObj
}

// Writer with some problems.
func NewMockWithProblemRepotWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return(mock.Report, nil)
	mockObj.On("Finish", mock.Report).Return(nil)
	mockObj.On("Commit", mock.Report, mock.Problems).Return(errors.New("Report has some problems"))
	return mockObj
}

// Writer with error and no report.
func NewMockErrorReportWriter() *ReportWriter {
	mockObj := new(ReportWriter)
	mockObj.On("Start").Return([]*models.Report{}, nil)
	mockObj.On("Finish", mock.Report).Return(nil)
	mockObj.On("Commit", []*models.Report{}, mock.Problems).Return(errors.New("Report wasn't created due to the error/problems"))
	return mockObj
}
