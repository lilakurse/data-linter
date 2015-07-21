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
	problem :=[]linter.Problem{}
	for _,value:=range checkers{
		if value/2 !=0 {
			return append(problem, "not even"), error
		}
	}
	return nil,nil
}

func (g *Generator) Name() string {
	testGenerator := Generator{Name:"Test generator", Count:1000}
	//return "This is a test random number generator"
	return testGenerator.Name
}

func (g *Generator) Count() int64{
	testGenerator := Generator{Name:"Test generator", Count:1000}
	return testGenerator.Count
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
