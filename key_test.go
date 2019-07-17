package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestNoKey(t *testing.T) {

	var key mdl.Key

	if expected, actual := mdl.NoKey(), key; expected != actual {
		t.Errorf("Expected an uninitialized mdl.Key to have a value of mdl.NoKey(), but actually didn't.")
		return
	}
}
