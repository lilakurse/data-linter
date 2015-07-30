package linter

import (
	"golang.org/x/net/context"
)
// Iterator contains documents and provides them in fixed-sized batches.
type Iterator interface {
	Name() string
	Count() int64
	Next(step int) []*Checker
}

func Name(ctx context.Context) string {
	return FromContext(ctx).Name()
}

func Count(ctx context.Context) int64{
	return FromContext(ctx).Count()
}

func Next(ctx context.Context, step int) []*Checker {
	return FromContext(ctx).Next(step)
}
