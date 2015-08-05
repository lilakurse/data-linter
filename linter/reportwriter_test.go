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

}
