package mock

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

var (
	StatisticS             = models.Count{Total: 50, Inspected: 50, Valid: 25, Invalid: 25}
	UnfinishedStatistics   = models.Count{Total: 50, Inspected: 10, Valid: 5, Invalid: 5}
	ReportTime             = time.Now()
	EmptyProblemList       = []*models.Problem{}
	Problems               = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
	ReportWithNoProblems   = NewReportWithNoProblems()
	UnfinishedReport       = NewReportWithError()
	ReportWithSomeProblems = NewReportWithProblems()
)

func NewReportWithNoProblems()*models.Report{
	report:=new(models.Report)
	report.Name="Customers"
	created:=time.Now().UTC()
	report.Created=&created
	updated:=time.Now().UTC() // need to fix time
	report.Updated=&updated
	finished:=time.Now().UTC() // need to fix time
	report.Finished=&finished
	report.Error="No error"
	report.Problems= EmptyProblemList
	report.Statistics= StatisticS
	return report

}

func NewReportWithProblems()*models.Report{
	report:=new(models.Report)
	report.Name="Freelancers"
	created:=time.Now().UTC()
	report.Created=&created
	updated:=time.Now().UTC() // need to fix time
	report.Updated=&updated
	finished:=time.Now().UTC() // need to fix time
	report.Finished=&finished
	report.Error="Some problems found"
	report.Statistics= StatisticS
	return report

}

func NewReportWithError()*models.Report{
	report:=new(models.Report)
	report.Name="Orders"
	created:=time.Now().UTC()
	report.Created=&created
	updated:=time.Now().UTC() // need to fix time
	report.Updated=&updated
	finished:=time.Now().UTC() // need to fix time
	report.Finished=&finished
	report.Error="Error occured while collecting problems"
	report.Problems= Problems
	report.Statistics= UnfinishedStatistics
	return report

}