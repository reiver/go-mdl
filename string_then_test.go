package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringThen(t *testing.T) {

	tests := []struct{
		Value    mdl.String
		Func     func(string)mdl.String
		Expected mdl.String
	}{
		{
			Value: mdl.NoString(),
			Func:  func(string)mdl.String{
				return mdl.NoString()
			},
			Expected: mdl.NoString(),
		},
		{
			Value: mdl.SomeString("Hi!"),
			Func:  func(string)mdl.String{
				return mdl.NoString()
			},
			Expected: mdl.NoString(),
		},



		{
			Value: mdl.NoString(),
			Func:  func(string)mdl.String{
				return mdl.SomeString("Hello world!")
			},
			Expected: mdl.NoString(),
		},
		{
			Value: mdl.SomeString("Hi!"),
			Func:  func(string)mdl.String{
				return mdl.SomeString("Hello world!")
			},
			Expected: mdl.SomeString("Hello world!"),
		},



		{
			Value: mdl.NoString(),
			Func:  func(s string)mdl.String{
				return mdl.SomeString("«" + s + "»")
			},
			Expected: mdl.NoString(),
		},
		{
			Value: mdl.SomeString("What?"),
			Func:  func(s string)mdl.String{
				return mdl.SomeString("Joe said, «" + s + "»")
			},
			Expected: mdl.SomeString("Joe said, «What?»"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Then(test.Func)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
