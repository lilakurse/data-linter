package linter

const Step = 50

// TODO: check logic (step 4)
// Inspect runs through the iterator and collects problems for report.
func Inspect(iterator Iterator) ([]*Problem, error) {
	//check := Checker()
	docsChecked := 0
	problemsToCommit := []*Problem{}
	for docsChecked <= iterator.Count() {
		for _, document := range iterator.Next(Step) {
			document = Checker().Check()  // We need to do document checker and then call Check() but i m not sure about that.// TODO: fix
			docProblems, err := document
			if err != nil {
				return nil, err
			}
			problemsToCommit = append(problemsToCommit, docProblems)
		}
		docsChecked += Step
	}

	return problemsToCommit, nil
}
