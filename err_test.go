package mye_test

import (
	"errors"
	"testing"

	"github.com/edgarSucre/mye"
)

func TestSome(t *testing.T) {
	ve := mye.Validation.New("validation error")
	chained := mye.Wrap(ve, "nonsense")
	chained2 := mye.Wrap(chained, "more nonsense")

	if !errors.As(chained2, new(mye.Err)) {
		t.Fail()
	}
}
