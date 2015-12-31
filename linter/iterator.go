package linter

import (
	"github.com/lilakurse/data-linter/linter/checker"
)

// Iterator contains documents and provides them in fixed-sized batches.
type Iterator interface {
	Name() string
	Count() int
	Next(step int) []checker.Checker
}