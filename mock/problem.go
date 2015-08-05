package mock

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

var (
	StatisticS             = models.Count{Total: 2, Inspected: 1, Valid: 1, Invalid: 0}
	BadStatisticS          = models.Count{Total: 1, Inspected: 1, Valid: 0, Invalid: 1}
	time                   = time.Now()
	EmptyProblemList       = []*models.Problem{}
	Problems               = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
	ReportWithNoProblems   = *models.Report{Name: "Customers", Created: time, Updated: time + 2, Finished: time + 3, Error: "No error", Problems: EmptyProblemList, Statistics: StatisticS}
	UnfinishedReport       = *models.Report{Name: "Orders", Created: time + 3, Updated: time + 4, Error: "Error occured while collecting problems", Problems: Problems, Statistics: BadStatisticS}
	ReportWithSomeProblems = *models.Report{Name: "Freelancers", Created: time + 1, Updated: time + 2, Finished: time + 3, Error: "Some problems found", Statistics: StatisticS}
)
