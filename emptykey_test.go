package mdl

import (
	"testing"
)

func TestInternalEmptyKeyAsError(t *testing.T) {
	var err error = internalEmptyKey{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == err {
		t.Errorf("This should never happen.")
		return
	}
}

func TestInternalEmptyKeyAsEmptyKey(t *testing.T) {
	var complainer EmptyKey = internalEmptyKey{} // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == complainer {
		t.Errorf("This should never happen.")
		return
	}
}
