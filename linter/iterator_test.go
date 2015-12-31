package linter

import (
	"testing"
	"github.com/lilakurse/data-linter/linter/mocks"
	"github.com/lilakurse/data-linter/mock"
)

var (
	mockIteratorWithoutCheckers = Iterator(mocks.NewMockIteratorWithoutCheckers())
	mockIteratorWithCheckers = Iterator(mocks.NewMockIteratorWithCheckers())
)

func TestName(t *testing.T) {
	// Empty iterator
	name := mockIteratorWithoutCheckers.Name()
	if len(name) == 0 {
		t.Errorf("Iterator should have a name")
	}

	// Non-empty iterator
	name = mockIteratorWithCheckers.Name()
	if len(name) == 0 {
		t.Errorf("Iterator should have a name")
	}

}

func TestCount(t *testing.T) {
	// Empty iterator
	count := mockIteratorWithoutCheckers.Count()
	if count != mock.EmptyCount {
		t.Errorf("Iterator should be empty")
	}

	// Non-empty iterator
	count = mockIteratorWithCheckers.Count()
	if count != mock.NonEmptyCount {
		t.Errorf("Iterator should not be empty")
	}

}

func TestNext(t *testing.T) {
	// Empty iterator
	checkers := mockIteratorWithoutCheckers.Next(mock.Step)
	if checkers != nil {
		t.Errorf("The list of checkers should be empty")
	}

	// Non-empty iterator
	checkers = mockIteratorWithCheckers.Next(mock.Step)
	if checkers == nil && len(checkers) != mock.Step {
		t.Errorf("The list of checkers should not be empty")
	}

}