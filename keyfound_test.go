package mdl

import (
	"testing"
)

func TestInternalKeyFoundAsError(t *testing.T) {
	var err error = internalKeyFound{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

func TestInternalKeyFoundAsKeyFound(t *testing.T) {
	var complainer KeyFound = internalKeyFound{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == complainer {
		t.Errorf("This should never happen.")
		return
	}
}
