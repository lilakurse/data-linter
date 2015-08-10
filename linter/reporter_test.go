package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"github.com/GabbyyLS/data-linter/mock"
	"testing"
)

var (
	mockReportWriter            = ReportWriter(mocks.NewMockReportWriter())
	mockReportWriterWithProblem = ReportWriter(mocks.NewMockWithProblemRepotWriter())
	mockReportWriterWithError   = ReportWriter(mocks.NewMockErrorReportWriter())
	mockAllReportReader         = ReportReader(mocks.NewMockAllReportReader())
)

func TestStart(t *testing.T) {
	// ReportWriter with empty list of problems
	start, err := mockReportWriter.Start()
	if err != nil {
		t.Errorf("Report has problems")
	}
	if start.Statistics.Total != start.Statistics.Inspected || start.Statistics.Valid == 0 {
		t.Errorf("Report has problems")
	}
	// ReportWriter with some problems
	start, err = mockReportWriterWithProblem.Start()
	if err != nil {
		t.Errorf("Error")
	}
	if start.Statistics.Total != start.Statistics.Inspected || start.Statistics.Invalid == 0 { //not sure about second part
		t.Errorf("Report has problems")
	}
	// ReportWriter with errors
	start, err = mockReportWriterWithError.Start()
	if err != nil {
		t.Errorf("Error")
	}
	if start.Statistics.Total == start.Statistics.Inspected {
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

// этот тест не проходит
func TestTotalReportsCount(t *testing.T) {
	count := mockAllReportReader.TotalReportsCount()
	sum := count.Failed + count.Successful
	if sum != count.Total {
		t.Errorf("Total should be equal to the sum of successful and failed")
	}
}
