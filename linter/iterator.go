package linter

type Iterator interface {
	Name() string // TODO: do we need this one?
	Count() int64
	Next(Step int) []Checker // TODO: []*Checker?
}
