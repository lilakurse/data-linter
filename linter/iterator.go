package linter

// The Iterator contains documents and provides them in fixed-sized batches.
type Iterator interface {
	Name() string
	Count() int64
	Next(Step int) []*Checker
}
