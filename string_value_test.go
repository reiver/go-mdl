package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestStringValue(t *testing.T) {

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

		actualInterface, err := test.Value.Value()

		if !test.ExpectedLoaded && nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %#v", testNumber, err)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tEXPECTED: %#v", test.Expected)
			t.Errorf("\tACTUAL: %T %#v", actualInterface, actualInterface)
			continue
		}

		if test.ExpectedLoaded && nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tEXPECTED: %#v", test.Expected)
			t.Errorf("\tACTUAL: %T %#v", actualInterface, actualInterface)
			continue
		}

		if !test.ExpectedLoaded {
			continue
		}
		actual, casted := actualInterface.(string)
		if !casted {
			t.Errorf("For test #%d, expected type ‘string’, but actually got type ‘%T’.", testNumber, actualInterface)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tEXPECTED: %#v", test.Expected)
			t.Errorf("\tACTUAL: %T %#v", actualInterface, actualInterface)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tEXPECTED loaded: %t", test.ExpectedLoaded)
			t.Errorf("\tACTUAL   err:   (%T) %q", err, err)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
