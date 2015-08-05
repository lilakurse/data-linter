package mock

import (
	"github.com/GabbyyLS/data-linter/models"
	"time"
)

var (
	Statistic              = models.Count{Total: 1, Inspected: 1, Valid: 1, Invalid: 0}
	BadStatistic           = models.Count{Total: 1, Inspected: 1, Valid: 0, Invalid: 1}
	t                      = time.Now()
	EmptyProblemList       = []*models.Problem{}
	Problems               = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
	ReportWithNoProblems   = models.Report{Name: "Customer", Created: t, Updated: t + 2, Finished: t + 3, Error: "No error", Problems: Problems, Statistics: Statistic}
	UnfinishedReport       = models.Report{Name: "Orders", Created: t + 3, Updated: t + 4, Finished: t + 5, Error: "Error occured while finishing", Problems:Problems ,Statistics: BadStatistic}
	ReportWithSomeProblems = models.Report{Name:"Freelancer", Created: t + 1, Updated: t + 2, Finished: t + 3, Error: "Error some problems occured", Statistics: Statistic}
)
