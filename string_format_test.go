package mdl_test

import (
	"github.com/reiver/go-mdl"

	"fmt"

	"testing"
)

func TestStringFormat(t *testing.T) {

	tests := []struct{
		Format   string
		String   mdl.String
		Expected string
	}{
		{
			Format: "%q",
			String: mdl.NoString(),
			Expected:        `«no-string»`,
		},
		{
			Format: "%q",
			String: mdl.SomeString(""),
			Expected:             `""`,
		},
		{
			Format: "%q",
			String: mdl.SomeString("Hello world!"),
			Expected:             `"Hello world!"`,
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf(test.Format, test.String)
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, for print verb, did not actually get what was expected.", testNumber)
			t.Logf("\tEXPECTED: %q", expected)
			t.Logf("\tACTUAL:   %q", actual)
			t.Logf("\tFormat: %q", test.Format)
			t.Logf("\tString: %#v", test.String)
			continue
		}

	}
}
