package checker

import "github.com/lilakurse/data-linter/models"

//import (
//	"golang.org/x/net/context"
//)

// The Checker runs various checks for validity on itself and returns a list of problems, if any.
type Checker interface {
	Check() ([]*models.Problem, error)
}

//func Check(ctx context.Context) ([]*Problem, error){
//	return CheckerFromContext(ctx).Check()
//}



