package iterators

import (
	//"fmt"
	"github.com/GabbyyLS/data-linter/linter"
	"math/rand"
	"time"
)

// This is the realization of the Iterator interface that through Next() returns batches of random Numbers.
// Each Number is a Checker, i.e., it has the Check() method implemented.
type Number int
type Generator struct {
	Name  string
	Count int64
}

func (Number) Check() ([]*linter.Problem, error) {
	problems := []*linter.Problem{}
	value := Number()
	if value/2 != 0 {
		problem := "not even"
		problems = append(problems, problem)
		return problems, error
	}
	return value,nil
}

func (g *Generator) Name() string {
	return g.Name
}

func (g *Generator) Count() int64 {
	return g.Count
}

func (g *Generator) Next(Step int) []linter.Checker {
	checkers := []linter.Checker{}
	for i := 0; i < Step; i++ {
		rand.Seed(time.Now().UnixNano())
		k := rand.Int()
		checker := Number(k)
		checkers = append(checkers, checker)

	}
	//	if Step>Count()
	//		return nil
	return checkers
	// TODO: add some checking against Count
}
