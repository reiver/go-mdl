package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestValueElseUnwrap(t *testing.T) {

	tests := []struct{
		Value    mdl.String
		Else     string
		Expected string
	}{
		{
			Value:    mdl.NoString(),
			Else:                    "defaulted",
			Expected:                "defaulted",
		},
		{
			Value:    mdl.SomeString("Hi!"),
			Else:                    "defaulted",
			Expected:                "Hi!",
		},



		{
			Value:    mdl.NoString(),
			Else:                    "Hello world!",
			Expected:                "Hello world!",
		},
		{
			Value:    mdl.SomeString("Hi!"),
			Else:                    "Hello world!",
			Expected:                "Hi!",
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.ElseUnwrap(test.Else)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tVALUE: %#v", test.Value)
			t.Errorf("\tELSE:                   %q", test.Else)
			t.Errorf("\tEXPECTED:               %q", expected)
			t.Errorf("\tACTUAL:                 %q", actual)
			continue
		}
	}
}
