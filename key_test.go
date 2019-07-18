package mdl_test

import (
	"github.com/reiver/go-mdl"

	"reflect"
	"strings"

	"testing"
)

func TestNoKey(t *testing.T) {

	var key mdl.Key

	if expected, actual := mdl.NoKey(), key; expected != actual {
		t.Errorf("Expected an uninitialized mdl.Key to have a value of mdl.NoKey(), but actually didn't.")
		return
	}
}

func TestSomeKey(t *testing.T) {

	var this mdl.Key = mdl.SomeKey("apple", "banana", "cherry")
	var that mdl.Key = mdl.SomeKey("apple", "banana", "cherry")

	if this != that {
		t.Errorf("Expected two mdl.Key assigned the same value with mdl.SomeKey() to be equal, but actually aren't.")
		return
	}
}

func TestSomeKeyNoKey(t *testing.T) {

	var key mdl.Key = mdl.SomeKey()

	if expected, actual := mdl.NoKey(), key; expected != actual {
		t.Errorf("Expected mdl.SomeKey() to equal mdl.NoKey(), but actually wasn't.")
		return
	}
}

func TestKeyElse(t *testing.T) {

	tests := []struct{
		Key     mdl.Key
		Else  []string
		Expected mdl.Key
	}{
		{
			Key: mdl.NoKey(),
			Else: []string{"apple", "banana", "cherry"},
			Expected: mdl.SomeKey("apple", "banana", "cherry"),
		},
		{
			Key: mdl.SomeKey("one", "two"),
			Else: []string{"apple", "banana", "cherry"},
			Expected: mdl.SomeKey("one", "two"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Key.Else(test.Else...)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the key which was actually gotten from .Else(), was not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestKeyElseUnwrap(t *testing.T) {

	tests := []struct{
		Key        mdl.Key
		Else     []string
		Expected []string
	}{
		{
			Key:     mdl.NoKey(),
			Else:     []string{"apple", "banana", "cherry"},
			Expected: []string{"apple", "banana", "cherry"},
		},
		{
			Key:   mdl.SomeKey("one", "two"),
			Else:     []string{"apple", "banana", "cherry"},
			Expected: []string{"one", "two"},
		},
	}

	for testNumber, test := range tests {

		actual := test.Key.ElseUnwrap(test.Else...)

		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the key which was actually gotten from .ElseUnwrap(), was not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestKeyGoString(t *testing.T) {

	tests := []struct{
		Key      mdl.Key
		Expected string
	}{
		{
			Key:       mdl.NoKey(),
			Expected: "mdl.NoKey()",
		},
		{
			Key:       mdl.SomeKey(),
			Expected: "mdl.NoKey()",
		},



		{
			Key:       mdl.SomeKey("apple"),
			Expected: `mdl.SomeKey("apple")`,
		},
		{
			Key:       mdl.SomeKey("apple", "banana"),
			Expected: `mdl.SomeKey("apple", "banana")`,
		},
		{
			Key:       mdl.SomeKey("apple", "banana", "cherry"),
			Expected: `mdl.SomeKey("apple", "banana", "cherry")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.Key.GoString()

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, what was expected from .GoString() is not what was actually received.", testNumber)
			t.Logf("EXPECTED: «%s»", expected)
			t.Logf("ACTUAL:   «%s»", actual)
			continue
		}
	}
}

func TestKeyMap(t *testing.T) {

	tests := []struct{
		FN func(...string)[]string
		Key mdl.Key
		Expected mdl.Key
	}{
		{
			FN: func(...string)[]string{
				return []string{"apple","banana", "cherry"}
			},
			Key:      mdl.NoKey(),
			Expected: mdl.NoKey(),
		},
		{
			FN: func(...string)[]string{
				return []string{"apple","banana", "cherry"}
			},
			Key:      mdl.SomeKey(),
			Expected: mdl.NoKey(),
		},



		{
			FN: func(...string)[]string{
				return []string{"apple","banana", "cherry"}
			},
			Key:      mdl.SomeKey("one", "two"),
			Expected: mdl.SomeKey("apple","banana", "cherry"),
		},



		{
			FN: func(a ...string)[]string{
				var b []string

				for _, aa := range a {
					var bb string = strings.ToUpper(aa)

					b = append(b, bb)
				}
				return b
			},
			Key:      mdl.SomeKey("One", "TWO", "three"),
			Expected: mdl.SomeKey("ONE","TWO", "THREE"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Key.Map(test.FN)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, what was expected from .Map() is not what was actually received.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
