package iterators

import (
	"fmt"
	"github.com/GabbyyLS/data-linter/linter"
	"math/rand"
	"time"
)

// This is the realization of the Iterator interface that through Next() returns batches of random Numbers.
// Each Number is a Checker, i.e., it has the Check() method implemented.
type Number int
type Generator struct {
	Name  *linter.Report
	Count *linter.Count
	Next  *linter.Checker
}

func (Number) Check() ([]*linter.Problem, error) {
	// TODO
}

func (g *Generator) Name() string {
	return fmt.Sprintf("This is a test random number generator")
}

func (g *Generator) Next(Step int) []linter.Checker {
	checkers := []linter.Checker{}
	for i := 0; i < Step; i++ {
		rand.Seed(time.Now().UnixNano())
		k := rand.Int()
		checker := Number(k)
		checkers = append(checkers, checker)
	}
	return checkers
	// TODO: add some checking against Count
}
