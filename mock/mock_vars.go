package mock

type Document struct {
	Name string
}

var (
	InvalidDoc = &Document{Name: "This is an invalid document"}
	ValidDoc   = &Document{Name: "This is an valid document"}
)
