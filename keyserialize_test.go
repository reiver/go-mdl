package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestKeySerialize(t *testing.T) {

	tests := []struct{
		Key    []string
		Expected string
	}{
		{
			Key: []string(nil),
			Expected :    "",
		},
		{
			Key: []string{},
			Expected :    "",
		},
		{
			Key: []string{""},
			Expected :    "",
		},



		{
			Key: []string{"apple"},
			Expected :    "apple",
		},
		{
			Key: []string{"banana"},
			Expected :    "banana",
		},
		{
			Key: []string{"cherry"},
			Expected :    "cherry",
		},



		{
			Key: []string{"apple", "banana"},
			Expected :    "apple/banana",
		},
		{
			Key: []string{"apple", "banana", "cherry"},
			Expected :    "apple/banana/cherry",
		},



		{
			Key: []string{"i/o"},
			Expected :    `i\/o`,
		},



		{
			Key: []string{"🙂", "slightly smiling face", "😀😃😄😁😆", "ac/dc", "if nil != err { return }"},
			Expected :    `🙂/slightly\ smiling\ face/😀😃😄😁😆/ac\/dc/if\ nil\ !=\ err\ \{\ return\ }`,
		},
	}

	for testNumber, test := range tests {
		actual := mdl.KeySerialize(test.Key...)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual encoded key, is not what was expected.", testNumber)
			t.Logf("\tEXPECTED: %q", expected)
			t.Logf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
