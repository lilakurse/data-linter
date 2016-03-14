package iterators

import (
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	//"strconv"
	"time"
	"github.com/lilakurse/data-linter/models"
	"github.com/lilakurse/data-linter/linter/checker"
	"strconv"
)

// This is the realization of the Iterator interface that through Next() returns batches of random Numbers.
// Each Number is a Checker, i.e., it has the Check() method implemented.
type Number int

type Generator struct {
	name  string			// Changed name because of error occuried "Generator has same name for the method and field"
	count int
}

func NewGenerator(name string,count int) Generator {
	return Generator{name,count}
}

func (n Number) Check() ([]*models.Problem, error) {
	var err error
	err = nil
	problems := []*models.Problem{}
	details := []*models.ProblemDetails{}
	if (n%2) != 0 {
		problem := new(models.Problem)
		detail := new(models.ProblemDetails)
		problem.Id = bson.NewObjectId().Hex()
		detail.Id = "1"
		problem.Original = strconv.Itoa(int(n))
		detail.Fragment = strconv.Itoa(int(n)) + " % 2"
		detail.Description = detail.Fragment + " != 0 - it's bad value for our case"
		details := append(details, detail)
		problem.Details = details
		problems = append(problems, problem)
		return problems, nil
	}
	return nil, err
}

func (g Generator) Name() string {
	return g.name
}

func (g Generator) Count() int {
	return g.count
}

func (g Generator) Next(Step int) []checker.Checker {
	checkers := []checker.Checker{}
	for i := 0; i < Step; i++ {
		rand.Seed(time.Now().UnixNano())
		checker := checker.Checker(Number(rand.Int()))
		checkers = append(checkers, checker)
	}
	return checkers
}
