package linter

const Step = 50

// TODO: document (step 3) and check logic (step 4)
// The Inspect function collects  problems in iterator for report.
func Inspect(iterator Iterator) ([]*Problem, error) {

	docsChecked := 0
	problemsToCommit := []*Problem{}

	for docsChecked <= iterator.Count() {
		for _, document := range iterator.Next(Step) {
			docProblems, err := document.Check()
			if err != nil {
				return nil, err
			}
			problemsToCommit = append(problemsToCommit, docProblems)
		}
		docsChecked += Step
	}

	return problemsToCommit, nil
}
