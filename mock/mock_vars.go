package mock

import "github.com/GabbyyLS/data-linter/linter"

type Document struct {
	Name string
}

var (
	InvalidDoc = &Document{Name: "This is an invalid document"}
	ValidDoc   = &Document{Name: "This is a valid document"}
	BadDoc     = &Document{Name: "666"}
	Problem    = linter.Problem{Original: InvalidDoc.Name}

)
