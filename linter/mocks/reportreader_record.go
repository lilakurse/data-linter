package mocks

import "github.com/GabbyyLS/data-linter/mock"

func NewMockAllReportReader() *ReportReader {
	mockObj := new(ReportReader)
	// For all reports
	mockObj.On("GetAllReports").Return(mock.ListOfReports)
	mockObj.On("TotalReportsCount").Return(mock.ReportCount)

	//  The ReportReader with no problem.
	mockObj.On("GetReportByName", mock.ReportWithNoProblems.Name).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByCreationTime", mock.ReportWithNoProblems.Created).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByUpdateTime", mock.ReportWithNoProblems.Updated).Return(mock.ReportWithNoProblems)
	mockObj.On("GetReportByCommitTime", mock.ReportWithNoProblems.Finished).Return(mock.ReportWithNoProblems)

	// The ReportReader with some problem.
	mockObj.On("GetReportByName", mock.ReportWithSomeProblems.Name).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByCreationTime", mock.ReportWithSomeProblems.Created).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByUpdateTime", mock.ReportWithSomeProblems.Updated).Return(mock.ReportWithSomeProblems)
	mockObj.On("GetReportByCommitTime", mock.ReportWithSomeProblems.Finished).Return(mock.ReportWithSomeProblems)

	// The ReportReader with error.
	mockObj.On("GetReportByName", mock.ReportWithError.Name).Return(mock.ReportWithError)
	mockObj.On("GetReportByCreationTime", mock.ReportWithError.Created).Return(mock.ReportWithError)
	mockObj.On("GetReportByUpdateTime", mock.ReportWithError.Updated).Return(mock.ReportWithError)
	mockObj.On("GetReportByCommitTime", mock.ReportWithError.Finished).Return(mock.ReportWithError)
	return mockObj
}

//func NewMockReportReaderNoProblem() *ReportReader {
//	mockObj := new(ReportReader)
//	mockObj.On("GetReportByName", mock.ReportWithNoProblems.Name).Return(mock.ReportWithNoProblems)
//	mockObj.On("GetReportByCreationTime", mock.ReportWithNoProblems.Created).Return(mock.ReportWithNoProblems)
//	mockObj.On("GetReportByUpdateTime", mock.ReportWithNoProblems.Updated).Return(mock.ReportWithNoProblems)
//	mockObj.On("GetReportByCommitTime", mock.ReportWithNoProblems.Finished).Return(mock.ReportWithNoProblems)
//	return mockObj
//}
//
//func NewMockReportReaderWithProblem() *ReportReader {
//	mockObj := new(ReportReader)
//	mockObj.On("GetReportByName", mock.ReportWithSomeProblems.Name).Return(mock.ReportWithSomeProblems)
//	mockObj.On("GetReportByCreationTime", mock.ReportWithSomeProblems.Created).Return(mock.ReportWithSomeProblems)
//	mockObj.On("GetReportByUpdateTime", mock.ReportWithSomeProblems.Updated).Return(mock.ReportWithSomeProblems)
//	mockObj.On("GetReportByCommitTime", mock.ReportWithSomeProblems.Finished).Return(mock.ReportWithSomeProblems)
//	return mockObj
//}
//func NewMockReportReaderWithError() *ReportReader {
//	mockObj := new(ReportReader)
//	mockObj.On("GetReportByName", mock.ReportWithError.Name).Return(mock.ReportWithError)
//	mockObj.On("GetReportByCreationTime", mock.ReportWithError.Created).Return(mock.ReportWithError)
//	mockObj.On("GetReportByUpdateTime", mock.ReportWithError.Updated).Return(mock.ReportWithError)
//	mockObj.On("GetReportByCommitTime", mock.ReportWithError.Finished).Return(mock.ReportWithError)
//	return mockObj
//}
