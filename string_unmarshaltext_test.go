package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringTextUnmarshaler(t *testing.T) {

	tests := []struct{
		Datum   []byte
		Expected  mdl.String
	}{
		{
			Datum:            []byte(nil),
			Expected: mdl.NoString(),
		},



		{
			Datum:            []byte("Hi!"),
			Expected: mdl.SomeString("Hi!"),
		},
		{
			Datum:            []byte("Hello world!"),
			Expected: mdl.SomeString("Hello world!"),
		},
		{
			Datum:            []byte("Apple banana CHERRY 🙂 “👾”."),
			Expected: mdl.SomeString("Apple banana CHERRY 🙂 “👾”."),
		},
	}

	for testNumber, test := range tests {

		var actual mdl.String

		err := actual.UnmarshalText(test.Datum)

		if expected := test.Expected; nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tDATUM: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tDATUM: (%T) %#v", test.Datum, test.Datum)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
