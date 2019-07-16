package mdl

import (
	"strings"

	"testing"
)

func TestEncodeToken(t *testing.T) {

	tests := []struct{
		Token    string
		Expected string
	}{
		{
			Token:    "",
			Expected: "",
		},



		{
			Token:    "apple",
			Expected: "apple",
		},
		{
			Token:    "banana",
			Expected: "banana",
		},
		{
			Token:    "cherry",
			Expected: "cherry",
		},



		{
			Token:    "Hello world!",
			Expected: `Hello\ world!`,
		},
		{
			Token:    "Hello\tworld!",
			Expected: `Hello\	world!`,
		},



		{
			Token:    "i/o",
			Expected: `i\/o`,
		},
		{
			Token:    "ac/dc",
			Expected: `ac\/dc`,
		},



		{
			Token:    `C:\Documents\Photo.jpg`,
			Expected: `C:\\Documents\\Photo.jpg`,
		},



		{
			Token:    "database/username{",
			Expected: `database\/username\{`,
		},



		{
			Token:    "if x/2 < 1 {",
			Expected: `if\ x\/2\ <\ 1\ \{`,
		},
	}

	for testNumber, test := range tests {

		var builder strings.Builder


		if err := encodeToken(&builder, test.Token); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("\tTOKEN: %q", test.Token)
			continue
		}

		if expected, actual := test.Expected, builder.String(); expected != actual {
			t.Errorf("For test #%d, the actual encoding is not was expected.", testNumber)
			t.Logf("\tEXPECTED encoding: %q", expected)
			t.Logf("\tACTUAL   encoding: %q", actual)
			t.Logf("\tTOKEN: %q", test.Token)
			continue
		}
	}
}
