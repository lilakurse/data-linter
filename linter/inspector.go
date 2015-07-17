package linter

const Step = 50

func Inspect(iterator Iterator) error {

	docsChecked := 0

	for docsChecked <= iterator.Count() {
		for _, element := range iterator.Next(Step) {
			err := element.Check()
			if err != nil {
				return err
			}
		}
		docsChecked += Step
	}

	return nil
}
