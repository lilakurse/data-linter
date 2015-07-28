package linter

import (
	"github.com/GabbyyLS/data-linter/mock"
	"testing"
)

func TestCheck(t *testing.T) {
	// Valid document.
	check, err := Check(ctx, mock.ValidDoc) // fix error Multiple-value Check()in single value context
	if err == nil || check != nil {
		t.Error("Valid document")
	}
	// Invalid document.
	check, err = Check(ctx, mock.InvalidDoc)
	if err == nil || check != nil {
		t.Error("Invalid document")
	}
	check, err = Check(ctx, mock.BadDoc)
	if err != nil || check == nil {
		t.Error("Error")
	}
}
