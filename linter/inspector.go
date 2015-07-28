package linter

const Step = 50

// Inspect runs through the iterator and collects problems for report.
func Inspect(iterator Iterator) ([]*Problem, error) {
	docsChecked := 0
	problemsToCommit := []*Problem{}
	for docsChecked <= iterator.Count() {
		for _, document := range iterator.Next(Step) {
			docProblems, err := document.(Checker).Check()
			if err != nil {
				return nil, err
			}
			problemsToCommit = append(problemsToCommit, docProblems)
		}
		docsChecked += Step
	}

	return problemsToCommit, nil
}
