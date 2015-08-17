package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"github.com/GabbyyLS/data-linter/mock"
	"testing"
)

var (
	mockReportWriterWithNoProblems = ReportWriter(mocks.NewMockReportWriterWithNoProblems())
	mockReportWithSomeProblems     = ReportWriter(mocks.NewMockReportWithSomeProblems())
	mockReportWriterWithError      = ReportWriter(mocks.NewMockReportWriterWithError())
	mockReportReader               = ReportReader(mocks.NewMockAllReportReader())

	mockReportReaderNoProblem      = ReportReader((mocks.NewMockReportReaderNoProblem()))
//	mockReportReaderWithProblem    = ReportReader(mocks.NewMockReportReaderWithProblem())
//	mockReportReaderWithError      = ReportReader(mocks.NewMockReportReaderWithError())
)

func TestStart(t *testing.T) {

	// Report with no problems
	_, err := mockReportWriterWithNoProblems.Start()
	if err != nil {
		t.Errorf("Report should have started with no problems")
	}

	// Report with some problems
	_, err = mockReportWithSomeProblems.Start()
	if err != nil {
		t.Errorf("Report should have started with no problems")
	}

	// Report with an error
	_, err = mockReportWriterWithError.Start()
	if err != nil {
		t.Errorf("Report should have started with no problems")
	}

}

func TestCommit(t *testing.T) {

	err := mockReportWriterWithNoProblems.Commit(mock.ReportWithNoProblems, mock.EmptyProblemList)
	if err != nil {
		t.Errorf("There should be no error, report is OK")
	}
	err = mockReportWithSomeProblems.Commit(mock.ReportWithSomeProblems, mock.Problems)
	if err != nil {
		t.Errorf("There should be no error, report is OK")
	}
	err = mockReportWriterWithError.Commit(mock.ReportWithError, mock.Problems)
	if err == nil {
		t.Errorf("Report should have some problems")
	}
}

func TestFinish(t *testing.T) {

	err := mockReportWriterWithNoProblems.Finish(mock.ReportWithNoProblems)
	if err != nil {
		t.Errorf("There should be no error, report is OK")
	}

	if mock.NewReportWithNoProblems().Statistics.Total != mock.NewReportWithNoProblems().Statistics.Inspected {
		t.Errorf("Number of inspected docs should be the same as the total number of docs")

	}

	if mock.NewReportWithNoProblems().Statistics.Valid != mock.NewReportWithNoProblems().Statistics.Total ||
		mock.NewReportWithNoProblems().Statistics.Invalid != 0 {
		t.Errorf("All docs should be valid")
	}

	if len(mock.NewReportWithNoProblems().Problems) != 0 {
		t.Errorf("Report should have no problems")
	}

	if mock.NewReportWithNoProblems().Error != "" {
		t.Errorf("Report should have no errors")
	}

	err = mockReportWithSomeProblems.Finish(mock.ReportWithSomeProblems)
	if err != nil {
		t.Errorf("There should be no error, report is OK")
	}

	if mock.NewReportWithSomeProblems().Statistics.Total != mock.NewReportWithSomeProblems().Statistics.Inspected {
		t.Errorf("Number of inspected docs should be the same as the total number of docs")
	}

	if mock.NewReportWithSomeProblems().Error != "" {
		t.Errorf("Report should have no errors")
	}

	if mock.NewReportWithSomeProblems().Statistics.Valid == mock.NewReportWithSomeProblems().Statistics.Total ||
		mock.NewReportWithSomeProblems().Statistics.Invalid == 0 {
		t.Errorf("Some docs should be invalid")
	}

	err = mockReportWriterWithError.Finish(mock.ReportWithError)
	if err == nil {
		t.Errorf("There should be an error")
	}

	if mock.NewReportWithError().Error == "" {
		t.Errorf("Report should contain an error")
	}

	if mock.NewReportWithError().Statistics.Valid == mock.NewReportWithError().Statistics.Total ||
		mock.NewReportWithError().Statistics.Invalid == 0 {
		t.Errorf("Some docs should be invalid")
	}

	if mock.NewReportWithError().Statistics.Total == mock.NewReportWithError().Statistics.Inspected {
		t.Errorf("Number of inspected docs should not be the same as the total number of docs")

	}

}

func TestGetAllReports(t *testing.T) {
	reportList := mockReportReader.GetAllReports()
	if len(reportList) == 0 {
		t.Errorf("List of reports should not be empty")
	}
}

func TestTotalReportsCount(t *testing.T) {
	reportCount := mockReportReader.TotalReportsCount()

	if reportCount.Failed + reportCount.Successful != reportCount.Total {
		t.Errorf("Total should be equal to the sum of successful and failed")
	}
}

func TestGetReportByName(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByName(mock.ReportWithNoProblems.Name)
	if report.Name == "" {
		t.Errorf("Report should have a name")
	}
//	report = mockReportReaderWithProblem.GetReportByName(mock.ReportWithSomeProblems.Name)
//	if report.Name == "" {
//		t.Errorf("Report should have a name")
//	}
//	report = mockReportReaderWithError.GetReportByName(mock.ReportWithError.Name)
//	if report.Name == "" {
//		t.Errorf("Report should have a name")
//	}
}

func TestGetReportByCreationTime(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByCreationTime(mock.ReportWithNoProblems.Created)
	if report.Created == nil {
		t.Errorf("Report doesn't exist")
	}
//	report = mockReportReaderWithProblem.GetReportByCreationTime(mock.ReportWithSomeProblems.Created)
//	if report.Created == nil {
//		t.Errorf("Report doesn't exist")
//	}
//	report = mockReportReaderWithError.GetReportByCreationTime(mock.ReportWithError.Created)
//	if report.Created == nil {
//		t.Errorf("Report doesn't exist")
//	}
}
//
//func TestGetReportByUpdateTime(t *testing.T) {
//	report := mockReportReaderNoProblem.GetReportByUpdateTime(mock.ReportWithNoProblems.Updated)
//	if report.Updated == nil {
//		t.Errorf("Report is old")
//	}
//	report = mockReportReaderWithProblem.GetReportByUpdateTime(mock.ReportWithSomeProblems.Updated)
//	if report.Updated == nil {
//		t.Errorf("Report is old")
//	}
//	report = mockReportReaderWithError.GetReportByUpdateTime(mock.ReportWithError.Updated)
//	if report.Updated == nil {
//		t.Errorf("Report is old")
//	}
//}
//
//func TestGetReportByCommitTime(t *testing.T) {
//	report := mockReportReaderNoProblem.GetReportByCommitTime(mock.ReportWithNoProblems.Finished)
//	if report.Finished == nil {
//		t.Errorf("Report has some problems")
//	}
//	report = mockReportReaderWithProblem.GetReportByCommitTime(mock.ReportWithSomeProblems.Finished)
//	if report.Finished == nil {
//		t.Errorf("Report has some problems")
//	}
//	report = mockReportReaderWithError.GetReportByCommitTime(mock.ReportWithError.Finished)
//	if report.Finished != nil {
//		t.Errorf("Finished time should be nil")
//	}
//}
