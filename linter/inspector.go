package linter

import "github.com/GabbyyLS/data-linter/models"

const Step = 50

// Inspect runs through the iterator and collects problems for report.
func Inspect(iterator Iterator) ([]*models.Problem, error) {
	docsChecked := 0
	problemsToCommit := []*models.Problem{}
	for docsChecked <= iterator.Count() {
		for _, document := range iterator.Next(Step) {
			docToCheck := *document
			docProblems, err := docToCheck.Check()
			if err != nil {
				return nil, err
			}
			problemsToCommit = append(problemsToCommit, docProblems...)
		}
		// TODO: We need to commit here!
		docsChecked += Step
	}
	return problemsToCommit, nil
}
