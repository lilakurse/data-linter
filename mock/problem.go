package mock

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

var (
	StatsForReportWithNoProblems   = models.Count{Total: 50, Inspected: 50, Valid: 50, Invalid: 0}
	StatsForReportWithSomeProblems = models.Count{Total: 50, Inspected: 50, Valid: 48, Invalid: 2}
	StatsForReportWithError        = models.Count{Total: 50, Inspected: 10, Valid: 8, Invalid: 2}
	EmptyProblemList               = []*models.Problem{}
	Problems                       = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
	ReportWithNoProblems           = NewReportWithNoProblems()
	ReportWithSomeProblems         = NewReportWithSomeProblems()
	ReportWithError                = NewReportWithError()
	ListOfReports                  = []*models.Report{} // не знаю как сделать слайс репортов
	//ListOfReports          = append(ListOfReports, ReportWithNoProblems, UnfinishedReport, ReportWithSomeProblems)
	ReportCount = models.ReportsCount{Total: 3, Failed: 1, Successful: 2}
)

func NewReportWithNoProblems() *models.Report {
	report := new(models.Report)
	report.Name = "Report With No Problems"
	created := time.Now().UTC()
	report.Created = &created
	updated := time.Now().UTC().AddDate(0, 0, 1)
	report.Updated = &updated
	finished := time.Now().UTC().AddDate(0, 0, 2)
	report.Finished = &finished
	report.Error = ""
	report.Problems = EmptyProblemList
	report.Statistics = StatsForReportWithNoProblems
	return report

}

func NewReportWithSomeProblems() *models.Report {
	report := new(models.Report)
	report.Name = "Report With Some Problems"
	created := time.Now().UTC()
	report.Created = &created
	updated := time.Now().UTC().AddDate(0, 0, 1)
	report.Updated = &updated
	finished := time.Now().UTC().AddDate(0, 0, 2)
	report.Finished = &finished
	report.Problems = Problems
	report.Error = ""
	report.Statistics = StatsForReportWithSomeProblems
	return report

}

func NewReportWithError() *models.Report {
	report := new(models.Report)
	report.Name = "Report With Error"
	created := time.Now().UTC()
	report.Created = &created
	updated := time.Now().UTC().AddDate(0, 0, 1)
	report.Updated = &updated
	report.Error = "An error occured while collecting problems -- report is partial!"
	report.Problems = Problems
	report.Statistics = StatsForReportWithError
	return report

}
