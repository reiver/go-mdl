package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringElse(t *testing.T) {

	tests := []struct{
		Value    mdl.String
		Else     string
		Expected mdl.String
	}{
		{
			Value:    mdl.NoString(),
			Else:                    "defaulted",
			Expected: mdl.SomeString("defaulted"),
		},
		{
			Value:    mdl.SomeString("Hi!"),
			Else:                    "defaulted",
			Expected: mdl.SomeString("Hi!"),
		},



		{
			Value:    mdl.NoString(),
			Else:                    "Hello world!",
			Expected: mdl.SomeString("Hello world!"),
		},
		{
			Value:    mdl.SomeString("Hi!"),
			Else:                    "Hello world!",
			Expected: mdl.SomeString("Hi!"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Else(test.Else)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tVALUE:    %#v", test.Value)
			t.Errorf("\tELSE:                      %q", test.Else)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
