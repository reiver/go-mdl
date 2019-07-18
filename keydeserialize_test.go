package mdl_test

import (
	"github.com/reiver/go-mdl"

	"reflect"

	"testing"
)

func TestKeyDeserialize(t *testing.T) {

	tests := []struct{
		Serialized string
		Expected []string
	}{
		{
			Serialized:        "",
			Expected: []string(nil),
		},


		{
			Serialized:        "apple",
			Expected: []string{"apple"},
		},
		{
			Serialized:        "banana",
			Expected: []string{"banana"},
		},
		{
			Serialized:        "cherry",
			Expected: []string{"cherry"},
		},



		{
			Serialized:        "apple/banana",
			Expected: []string{"apple", "banana"},
		},
		{
			Serialized:        "apple/banana/cherry",
			Expected: []string{"apple", "banana", "cherry"},
		},



		{
			Serialized:        `i\/o`,
			Expected: []string{"i/o"},
		},



		{
			Serialized:        `🙂/slightly\ smiling\ face/😀😃😄😁😆/ac\/dc/if\ nil\ !=\ err\ \{\ return\ }`,
			Expected: []string{"🙂", "slightly smiling face", "😀😃😄😁😆", "ac/dc", "if nil != err { return }"},
		},
	}

	for testNumber, test := range tests {
		actual, err := mdl.KeyDeserialize(test.Serialized)
		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error, but actually got one: (%T) %q.", testNumber, err, err)
			t.Logf("SERIALZED: «%s»", test.Serialized)
			continue
		}

		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual encoded key, is not what was expected.", testNumber)
			t.Logf("\tEXPECTED: %#v", expected)
			t.Logf("\tACTUAL:   %#v", actual)
			t.Logf("SERIALZED: «%s»", test.Serialized)
			continue
		}
	}
}
