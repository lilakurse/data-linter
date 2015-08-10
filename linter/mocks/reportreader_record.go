package mocks

import "github.com/GabbyyLS/data-linter/mock"

func NewMockAllReportReader() *ReportReader {
	mockObj := new(ReportReader)
	mockObj.On("GetAllReports").Return(mock.ListOfReports)
	mockObj.On("TotalReportsCount").Return(mock.ReportCount)
	return mockObj
}

func NewMockReportReaderNoProblem() *ReportReader {
	mockObj := new(ReportReader)
	mockObj.On("GetReportByName", mock.ReportWithNoProblems.Name).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByCreationTime", mock.ReportWithNoProblems.Created).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByUpdateTime", mock.ReportWithNoProblems.Updated).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByCommitTime", mock.ReportWithNoProblems.Finished).Return(mock.ReportWithNoProblems)
	return mockObj
}

func NewMockReportReaderWithProblem() *ReportReader {
	mockObj := new(ReportReader)
	mockObj.On("GetReportByName", mock.ReportWithSomeProblems.Name).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByCreationTime", mock.ReportWithSomeProblems.Created).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByUpdateTime", mock.ReportWithSomeProblems.Updated).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByCommitTime", mock.ReportWithSomeProblems.Finished).Return(mock.ReportWithSomeProblems)
	return mockObj
}
func NewMockReportReaderWithError() *ReportReader {
	mockObj := new(ReportReader)
	mockObj.On("GetReportByName", mock.UnfinishedReport.Name).Return(mock.UnfinishedReport)
	mockObj.On("GetReportByCreationTime", mock.UnfinishedReport.Created).Return(mock.UnfinishedReport)
	mockObj.On("GetReportByUpdateTime", mock.UnfinishedReport.Updated).Return(mock.UnfinishedReport)
	mockObj.On("GetReportByCommitTime", mock.UnfinishedReport.Finished).Return(mock.UnfinishedReport)
	return mockObj
}
