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

func TestSomeKey(t *testing.T) {

	var this mdl.Key = mdl.SomeKey("apple", "banana", "cherry")
	var that mdl.Key = mdl.SomeKey("apple", "banana", "cherry")

	if this != that {
		t.Errorf("Expected two mdl.Key assigned the same value with mdl.SomeKey() to be equal, but actually aren't.")
		return
	}
}

func TestSomeKeyNoKey(t *testing.T) {

	var key mdl.Key = mdl.SomeKey()

	if expected, actual := mdl.NoKey(), key; expected != actual {
		t.Errorf("Expected mdl.SomeKey() to equal mdl.NoKey(), but actually wasn't.")
		return
	}
}

