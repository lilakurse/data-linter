package mock

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

var (
	StatsForReportWithNoProblems   = models.Statistics{Total: 50, Inspected: 50, Valid: 50, Invalid: 0}
	StatsForReportWithSomeProblems = models.Statistics{Total: 50, Inspected: 50, Valid: 48, Invalid: 2}
	StatsForReportWithError        = models.Statistics{Total: 50, Inspected: 10, Valid: 8, Invalid: 2}
	EmptyProblemList               = []*models.Problem{}
	Problems                       = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
	ReportWithNoProblems           = NewReportWithNoProblems()
	ReportWithSomeProblems         = NewReportWithSomeProblems()
	ReportWithError                = NewReportWithError()
	ListOfReports                  = []*models.Report{ReportWithNoProblems, ReportWithSomeProblems, ReportWithError}
	ReportCount                    = &models.ReportsCount{Total: 3, Failed: 1, Successful: 2}
	Created                        time.Time
	Updated                        time.Time
	Finished                       time.Time
)

// The function creates a report as if we have already commited to it at least once and finished it
// (i.e., report.Finished and report.Updated are already set, report has Statistics, etc.)
// This is needed for mock-tests.
func NewReportWithNoProblems() *models.Report {
	report := new(models.Report)
	report.Name = "Report With No Problems"
	Created = time.Date(2015, 11, 01, 14, 55, 13, 0, time.Local)
	report.Created = &Created
	Updated = time.Date(2015, 11, 02, 14, 55, 13, 0, time.Local)
	report.Updated = &Updated
	Finished = time.Date(2015, 11, 02, 21, 00, 00, 0, time.Local)
	report.Finished = &Finished
	report.Error = ""
	report.Problems = EmptyProblemList
	report.Statistics = StatsForReportWithNoProblems
	return report
}

// The function creates a report as if we have already commited to it at least once and finished it
// (i.e., report.Finished and report.Updated are already set, report has Statistics, etc.)
// This is needed for mock-tests.
func NewReportWithSomeProblems() *models.Report {
	report := new(models.Report)
	report.Name = "Report With Some Problems"
	Created = time.Date(2016, time.January, 13, 12, 10, 00, 00, time.UTC)
	report.Created = &Created
	Updated = time.Date(2016, time.January, 12, 12, 10, 00, 00, time.UTC)
	report.Updated = &Updated
	Finished = time.Date(2016, time.January, 12, 18, 00, 00, 00, time.UTC)
	report.Finished = &Finished
	report.Problems = Problems
	report.Error = ""
	report.Statistics = StatsForReportWithSomeProblems
	return report

}

// The function creates a report as if we have already commited to it at least once and finished it
// (i.e., report.Finished and report.Updated are already set, report has Statistics, etc.)
// This is needed for mock-tests.
func NewReportWithError() *models.Report {
	report := new(models.Report)
	report.Name = "Report With Error"
	Created = time.Date(2015, time.October, 05, 10, 00, 45, 00, time.UTC)
	report.Created = &Created
	Updated = time.Date(2015, time.October, 06, 10, 00, 00, 00, time.UTC)
	report.Updated = &Updated
	report.Error = "An error occured while collecting problems -- report is partial!"
	report.Problems = Problems
	report.Statistics = StatsForReportWithError
	return report

}
