package linter

import (
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"testing"
)

var (
	mockReportWriter            = ReportWriter(mocks.NewMockReportWriter())
	mockReportWriterWithProblem = ReportWriter(mocks.NewMockWithProblemRepotWriter())
	mockReportWriterWithError   = ReportWriter(mocks.NewMockErrorReportWriter())
)

func TestStart(t *testing.T) {
	// ReportWriter with empty list of problems
	start, err := Start(mockReportWriter)
	if err != nil {
		return t.Error("Report has problems, the list of problems shoud be empty")
	}
	if start.Statistics.Total != start.Statistics.Inspected || start.Statistics.Valid == 0 {
		return t.Errorf("Report has problems")
	}
	// ReportWriter with some problems
	start, err = Start(mockReportWriterWithProblem)
	if err != nil {
		return t.Errorf("Error")
	}
	if start.Statistics.Total != start.Statistics.Inspected || start.Statistics.Invalid == 0 { //not sure about second part
		return t.Errorf("Report has problems")
	}
	// ReportWriter with errors
	start, err = Start(mockReportWriterWithError)
	if err != nil {
		return t.Errorf("Error")
	}
	if start.Statistics.Total == start.Statistics.Inspected {
		return t.Errorf("Statistics should not be equal")
	}
}

func TestFinish(t *testing.T) {
	//
}
