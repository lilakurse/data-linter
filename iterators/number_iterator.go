package iterators

import (
	"math/rand"
	"time"
	"fmt"
)

//Seed. This method initializes the random source.
//When we use a seed based on the current time,
//each program execution will start with a different random sequence.

func Generator() int {
	rand.Seed(time.Now().UnixNano())
	g :=rand.Int()
	fmt.Print(g)
	return g
}

