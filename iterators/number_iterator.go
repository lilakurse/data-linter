package iterators

import (
	"fmt"
	"github.com/GabbyyLS/data-linter/linter"
	"math/rand"
	"time"
)

// This is the realization of the Iterator interface that through Next() returns batches of random Numbers.
// Each Number is a Checker, i.e., it has the Check() method implemented.
func Generator() int {
	// TODO: implement the Iterator interface
	rand.Seed(time.Now().UnixNano())
	g := rand.Int()
	fmt.Print(g) // TODO: eliminate from the final version
	return g
}

type Number int

func (Number) Check() ([]*linter.Problem, error) {
	// TODO
}
