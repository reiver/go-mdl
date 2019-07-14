package mdl

import (
	"testing"
)

func TestInternalUnsupportedSourceAsError(t *testing.T) {
	var err error = internalUnsupportedSource{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

func TestInternalUnsupportedSourceAUnsupportedSource(t *testing.T) {
	var complainer UnsupportedSource = internalUnsupportedSource{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == complainer {
		t.Errorf("This should never happen.")
		return
	}
}
