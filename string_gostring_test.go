package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringGoString(t *testing.T) {

	tests := []struct{
		Value    mdl.String
		Expected string
	}{
		{
			Value:     mdl.NoString(),
			Expected: "mdl.NoString()",
		},



		{
			Value:     mdl.SomeString("Hi!"),
			Expected: `mdl.SomeString("Hi!")`,
		},



		{
			Value:     mdl.SomeString("Hello world!"),
			Expected: `mdl.SomeString("Hello world!")`,
		},



		{
			Value:     mdl.SomeString("Apple banana CHERRY 🙂 “👾”."),
			Expected: `mdl.SomeString("Apple banana CHERRY 🙂 “👾”.")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.GoString()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
