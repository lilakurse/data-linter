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
	ReportWithNoProblems   = *models.Report{Name: "Customers", Created: ReportTime, Updated: ReportTime, Finished: ReportTime, Error: "No error", Problems: EmptyProblemList, Statistics: StatisticS}
	UnfinishedReport       = *models.Report{Name: "Orders", Created: ReportTime, Updated: ReportTime, Error: "Error occured while collecting problems", Problems: Problems, Statistics: UnfinishedStatistics}
	ReportWithSomeProblems = *models.Report{Name: "Freelancers", Created: ReportTime, Updated: ReportTime, Finished: ReportTime, Error: "Some problems found", Statistics: StatisticS}
)
