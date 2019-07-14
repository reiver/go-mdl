package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringUnwrap(t *testing.T) {

	tests := []struct{
		Value          mdl.String
		ExpectedLoaded bool
		Expected       string
	}{
		{
			Value: mdl.NoString(),
			ExpectedLoaded: false,
			Expected:             "",
		},



		{
			Value: mdl.SomeString("Hi!"),
			ExpectedLoaded: true,
			Expected:             "Hi!",
		},



		{
			Value: mdl.SomeString("Hello world!"),
			ExpectedLoaded: true,
			Expected:             "Hello world!",
		},



		{
			Value: mdl.SomeString("Apple banana CHERRY 🙂 “👾”."),
			ExpectedLoaded: true,
			Expected:             "Apple banana CHERRY 🙂 “👾”.",
		},
	}

	for testNumber, test := range tests {

		actual, actualLoaded := test.Value.Unwrap()

		if expected, expectedLoaded := test.Expected, test.ExpectedLoaded; expectedLoaded != actualLoaded {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", expectedLoaded)
			t.Errorf("\tACTUAL   loaded: %t", actualLoaded)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected, expectedLoaded := test.Expected, test.ExpectedLoaded; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", expectedLoaded)
			t.Errorf("\tACTUAL   loaded: %t", actualLoaded)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
