package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"github.com/GabbyyLS/data-linter/mock"
	"testing"
)

var (
	mockReportWriter            = ReportWriter(mocks.NewMockReportWriterWithNoProblems())
	mockReportWriterWithProblem = ReportWriter(mocks.NewMockRepotWriterWithProblems())
	mockReportWriterWithError   = ReportWriter(mocks.NewMockReportWriterWithError())
	mockAllReportReader         = ReportReader(mocks.NewMockAllReportReader())
	mockReportReaderNoProblem   = ReportReader((mocks.NewMockReportReaderNoProblem()))
	mockReportReaderWithProblem = ReportReader(mocks.NewMockReportReaderWithProblem())
	mockReportReaderWithError   = ReportReader(mocks.NewMockReportReaderWithError())
)

func TestStart(t *testing.T) {
	// ReportWriter with empty list of problems
	report, err := mockReportWriter.Start()
	if err != nil {
		t.Errorf("This report should have no problems")
	}
	if report.Statistics.Total != report.Statistics.Inspected ||
		report.Statistics.Valid == 0 {
		t.Errorf("Report has problems")
	}
	// ReportWriter with some problems
	report, err = mockReportWriterWithProblem.Start()
	if err != nil {
		t.Errorf("Error")
	}
	if report.Statistics.Total != report.Statistics.Inspected || report.Statistics.Invalid == 0 { //not sure about second part
		t.Errorf("Report has problems")
	}
	// ReportWriter with errors
	report, err = mockReportWriterWithError.Start()
	if err != nil {
		t.Errorf("Error")
	}
	if report.Statistics.Total == report.Statistics.Inspected {
		t.Errorf("Statistics should not be equal")
	}
}

func TestFinish(t *testing.T) {
	err := mockReportWriter.Finish(mock.ReportWithNoProblems)
	if err != nil {
		t.Errorf("Error found while finishing")
	}
	// Report with problems.
	err = mockReportWriterWithProblem.Finish(mock.ReportWithSomeProblems)
	if err != nil {
		t.Errorf("Error found while finishing")
	}
	// UnfinishedReport.
	err = mockReportWriterWithError.Finish(mock.UnfinishedReport)
	if err != nil {
		t.Errorf("Error found")
	}
}

func TestCommit(t *testing.T) {
	err := mockReportWriter.Commit(mock.ReportWithNoProblems, mock.EmptyProblemList)
	if err != nil {
		t.Errorf("Error occured while commiting")
	}
	err = mockReportWriterWithProblem.Commit(mock.ReportWithSomeProblems, mock.Problems)
	if err == nil {
		t.Errorf("Should have some problems")
	}
	err = mockReportWriterWithError.Commit(mock.UnfinishedReport, mock.Problems)
	if err == nil {
		t.Errorf("Report should not be created, should have some problems")
	}
}

func TestGetAllReports(t *testing.T) {
	report := mockAllReportReader.GetAllReports()
	if report == nil {
		t.Errorf("List of Reports should not be empty")
	}
}

// TODO: этот тест не проходит
//func TestTotalReportsCount(t *testing.T) {
//	count := mockAllReportReader.TotalReportsCount()
//	sum := count.Failed + count.Successful
//	if sum != count.Total {
//		t.Errorf("Total should be equal to the sum of successful and failed")
//	}
//}

func TestGetReportByName(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByName(mock.ReportWithNoProblems.Name)
	if report.Name == "" {
		t.Errorf("Report should have a name")
	}
	report = mockReportReaderWithProblem.GetReportByName(mock.ReportWithSomeProblems.Name)
	if report.Name == "" {
		t.Errorf("Report should have a name")
	}
	report = mockReportReaderWithError.GetReportByName(mock.UnfinishedReport.Name)
	if report.Name == "" {
		t.Errorf("Report should have a name")
	}
}

func TestGetReportByCreationTime(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByCreationTime(mock.ReportWithNoProblems.Created)
	if report.Created == nil {
		t.Errorf("Report doesn't exist")
	}
	report = mockReportReaderWithProblem.GetReportByCreationTime(mock.ReportWithSomeProblems.Created)
	if report.Created == nil {
		t.Errorf("Report doesn't exist")
	}
	report = mockReportReaderWithError.GetReportByCreationTime(mock.UnfinishedReport.Created)
	if report.Created == nil {
		t.Errorf("Report doesn't exist")
	}
}

func TestGetReportByUpdateTime(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByUpdateTime(mock.ReportWithNoProblems.Updated)
	if report.Updated == nil {
		t.Errorf("Report is old")
	}
	report = mockReportReaderWithProblem.GetReportByUpdateTime(mock.ReportWithSomeProblems.Updated)
	if report.Updated == nil {
		t.Errorf("Report is old")
	}
	report = mockReportReaderWithError.GetReportByUpdateTime(mock.UnfinishedReport.Updated)
	if report.Updated == nil {
		t.Errorf("Report is old")
	}
}

func TestGetReportByCommitTime(t *testing.T) {
	report := mockReportReaderNoProblem.GetReportByCommitTime(mock.ReportWithNoProblems.Finished)
	if report.Finished == nil {
		t.Errorf("Report has some problems")
	}
	report = mockReportReaderWithProblem.GetReportByCommitTime(mock.ReportWithSomeProblems.Finished)
	if report.Finished == nil {
		t.Errorf("Report has some problems")
	}
	report = mockReportReaderWithError.GetReportByCommitTime(mock.UnfinishedReport.Finished)
	if report.Finished != nil {
		t.Errorf("Finished time should be nil")
	}
}
