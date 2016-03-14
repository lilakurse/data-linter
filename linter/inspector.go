package linter

import "github.com/lilakurse/data-linter/models"

const Step = 50

//
type Inspector struct {
	reporter ReportWriter
}

func NewInspector(reporter ReportWriter) *Inspector {
	return &Inspector{reporter}
}

// Inspect runs through the iterator and collects problems for report.
// TODO: probably need to revise the signature
func (i *Inspector) Inspect(iterator Iterator) error {
	report,err := i.reporter.Start()
	if err != nil {
		return err
	}
	report.Statistics.Total = int64(iterator.Count())
	report.Statistics.Inspected = 0
	for report.Statistics.Inspected < report.Statistics.Total {
		problemsToCommit := []*models.Problem{}
		for _, document := range iterator.Next(Step) {
			problems,err := document.Check()
			if err != nil {
				return err
			}
			if len(problems) == 0 {
				report.Statistics.Valid++
			} else {
				report.Statistics.Invalid++
			}
			problemsToCommit = append(problemsToCommit, problems...)
		}
		err = i.reporter.Commit(report,problemsToCommit)
		if err != nil {
			return err
		}
		report.Statistics.Inspected += Step
	}
	err = i.reporter.Finish(report)
	if err != nil {
		return err
	}
	return nil
}
