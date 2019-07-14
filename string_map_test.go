package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringMap(t *testing.T) {

	tests := []struct{
		Value    mdl.String
		Func     func(string)string
		Expected mdl.String
	}{
		{
			Value: mdl.NoString(),
			Func:  func(string)string{
				return "Hello world!"
			},
			Expected: mdl.NoString(),
		},
		{
			Value: mdl.SomeString("Hi!"),
			Func:  func(string)string{
				return "Hello world!"
			},
			Expected: mdl.SomeString("Hello world!"),
		},



		{
			Value: mdl.NoString(),
			Func:  func(s string)string{
				return "«" + s + "»"
			},
			Expected: mdl.NoString(),
		},
		{
			Value: mdl.SomeString("What?"),
			Func:  func(s string)string{
				return "Joe said, «" + s + "»"
			},
			Expected: mdl.SomeString("Joe said, «What?»"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Map(test.Func)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
