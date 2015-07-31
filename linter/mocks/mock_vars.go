package mocks

import "github.com/GabbyyLS/data-linter/models"

//import "github.com/GabbyyLS/data-linter/linter"

//type Document struct {
//	Name string
//}

var (
// TODO: let's switch to these as soon as we're done with writing record files
//	InvalidDoc = &Document{Name: "This is an invalid document"}
//	ValidDoc   = &Document{Name: "This is a valid document"}
//	BadDoc     = &Document{Name: "666"}
//	Problem    = linter.Problem{Original: InvalidDoc.Name}

//	EmptyProblemList = []*struct{
//		Original string
//	}{
//
//	}

	EmptyProblemList = []*models.Problem{}

//	Problems = []*struct{
//		Original string
//	}{
//		{Original: "A file with a problem"},
//		{Original: "Another file with a problem"},
//	}

)
