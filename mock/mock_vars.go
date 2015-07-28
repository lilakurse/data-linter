package mock

type Document struct {
	Name string
}

var (
	invalidDoc = &Document{Name: "This is an invalid document"}
	validDoc   = &Document{Name: "This is an valid document"}
)
