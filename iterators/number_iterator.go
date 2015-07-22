package iterators

import (
	//"fmt"
	"github.com/GabbyyLS/data-linter/linter"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"src-tms.als.local/go/marketplace/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

// This is the realization of the Iterator interface that through Next() returns batches of random Numbers.
// Each Number is a Checker, i.e., it has the Check() method implemented.
type Number int
type Generator struct {
	Name  string
	Count int64
}

func (n Number) Check() ([]*linter.Problem, error) {
	err := nil
	problems := []*linter.Problem{}
	details := []*linter.ProblemDetails{}
	if Number/2 != 0 {
		problem := new(linter.Problem)
		detail := new(linter.ProblemDetails)
		problem.Id = bson.NewObjectId().Hex()
		detail.Id = "1"
		problem.Original = strconv.Itoa(Number())
		detail.Fragment = strconv.Itoa(Number())
		detail.Description = "not even"
		details := append(details, detail)
		problem.Details = details
		problems = append(problems, problem)
		return problems, nil
	}
	return nil, err
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
