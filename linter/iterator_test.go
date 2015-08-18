package linter

import (
	"testing"
	"github.com/GabbyyLS/data-linter/linter/mocks"
	"github.com/GabbyyLS/data-linter/mock"
)

var (
	mockIteratorWithoutCheckers = Iterator(mocks.NewMockIteratorWithoutCheckers())
	mockIteratorWithCheckers = Iterator(mocks.NewMockIteratorWithCheckers())
)

func TestName(t *testing.T) {
	// Empty iterator
	name := Name(mockIteratorWithoutCheckers)
	if len(name) == 0 {
		t.Errorf("Iterator should have a name")
	}

	// Non-empty iterator
	name = Name(mockIteratorWithCheckers)
	if len(name) == 0 {
		t.Errorf("Iterator should have a name")
	}

}

func TestCount(t *testing.T) {
	// Empty iterator
	count := Count(mockIteratorWithoutCheckers)
	if count != mock.EmptyCount {
		t.Errorf("Iterator should be empty")
	}

	// Non-empty iterator
	count = Count(mockIteratorWithCheckers)
	if count != mock.NonEmptyCount {
		t.Errorf("Iterator should not be empty")
	}

}

func TestNext(t *testing.T) {
	// Empty iterator
	checkers := Next(mockIteratorWithoutCheckers, mock.Step)
	if checkers != nil {
		t.Errorf("The list of checkers should be empty")
	}

	// Non-empty iterator
	checkers = Next(mockIteratorWithCheckers, mock.Step)
	if checkers == nil && len(checkers) != mock.Step {
		t.Errorf("The list of checkers should not be empty")
	}

}