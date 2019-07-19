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

func TestKeyFormat(t *testing.T) {

	tests := []struct{
		Format  string
		Key      mdl.Key
		Expected string
	}{
		{
			Format: "%q",
			Key: mdl.NoKey(),
			Expected:     `«no-key»`,
		},
		{
			Format: "%q",
			Key: mdl.SomeKey(),
			Expected:       `«no-key»`,
		},
		{
			Format: "%q",
			Key: mdl.SomeKey("apple"),
			Expected:       `"apple"`,
		},
		{
			Format: "%q",
			Key: mdl.SomeKey("apple", "banana"),
			Expected:       `"apple/banana"`,
		},
		{
			Format: "%q",
			Key: mdl.SomeKey("apple", "banana", "cherry"),
			Expected:       `"apple/banana/cherry"`,
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf(test.Format, test.Key)
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, for print verb, did not actually get what was expected.", testNumber)
			t.Logf("\tEXPECTED: %q", expected)
			t.Logf("\tACTUAL:   %q", actual)
			t.Logf("\tFormat: %q", test.Format)
			t.Logf("\tKey: %#v", test.Key)
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

func TestKeyString(t *testing.T) {

	tests := []struct{
		Key mdl.Key
		Expected string
		ExpectedErr string
	}{
		{
			Key: mdl.NoKey(),
			Expected: "",
			ExpectedErr: "mdl: Not Loaded",
		},
		{
			Key: mdl.SomeKey(),
			Expected: "",
			ExpectedErr: "mdl: Not Loaded",
		},



		{
			Key: mdl.SomeKey("apple", "banana", "cherry"),
			Expected: "apple/banana/cherry",
			ExpectedErr: "",
		},



		{
			Key: mdl.SomeKey("ONE", "Two", "three"),
			Expected: "ONE/Two/three",
			ExpectedErr: "",
		},
	}

	for testNumber, test := range tests {

		actual, actualErr := test.Key.String()

		switch actualErr {
		case nil:
			if expected := test.ExpectedErr; "" != expected {
				t.Errorf("For test #%d, expected an error, but did not actually get one.", testNumber)
				t.Logf("EXPECTED error message: %q", expected)
				t.Logf("ACTUAL   error:         %#v", actualErr)
				continue
			}
		default:
			if expected, actual := test.ExpectedErr, actualErr.Error(); expected != actual {
				t.Errorf("For test #%d, expected an error, but did not actually get one.", testNumber)
				t.Logf("EXPECTED error message: %q", expected)
				t.Logf("ACTUAL   error message: %q", actual)
			}
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, what was expected from .String() is not what was actually received.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestKeyThen(t *testing.T) {

	tests := []struct{
		FN func(...string)mdl.Key
		Key mdl.Key
		Expected mdl.Key
	}{
		{
			FN: func(...string)mdl.Key {
				return mdl.NoKey()
			},
			Key:           mdl.NoKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(...string)mdl.Key {
				return mdl.NoKey()
			},
			Key:           mdl.SomeKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(...string)mdl.Key {
				return mdl.NoKey()
			},
			Key:           mdl.SomeKey("apple", "banana", "cherry"),
			Expected:      mdl.NoKey(),
		},



		{
			FN: func(...string)mdl.Key {
				return mdl.SomeKey("one", "two")
			},
			Key:           mdl.NoKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(...string)mdl.Key {
				return mdl.SomeKey("one", "two")
			},
			Key:           mdl.SomeKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(...string)mdl.Key {
				return mdl.SomeKey("one", "two")
			},
			Key:           mdl.SomeKey("apple", "banana", "cherry"),
			Expected:      mdl.SomeKey("one", "two"),
		},



		{
			FN: func(a ...string)mdl.Key {
				var b []string

				for _, aa := range a {
					bb := strings.ToUpper(aa)
					b = append(b, bb)
				}

				return mdl.SomeKey(b...)
			},
			Key:           mdl.NoKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(a ...string)mdl.Key {
				var b []string

				for _, aa := range a {
					bb := strings.ToUpper(aa)
					b = append(b, bb)
				}

				return mdl.SomeKey(b...)
			},
			Key:           mdl.SomeKey(),
			Expected:      mdl.NoKey(),
		},
		{
			FN: func(a ...string)mdl.Key {
				var b []string

				for _, aa := range a {
					bb := strings.ToUpper(aa)
					b = append(b, bb)
				}

				return mdl.SomeKey(b...)
			},
			Key:           mdl.SomeKey("apple", "banana", "cherry"),
			Expected:      mdl.SomeKey("APPLE", "BANANA", "CHERRY"),
		},
	}

	for testNumber, test := range tests {

		actual := test.Key.Then(test.FN)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, what was expected from .Then() is not what was actually received.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestKeyValue(t *testing.T) {

	tests := []struct{
		Key mdl.Key
		Expected interface{}
		ExpectedErr string
	}{
		{
			Key: mdl.NoKey(),
			Expected: nil,
			ExpectedErr: "mdl: Not Loaded",
		},
		{
			Key: mdl.SomeKey(),
			Expected: nil,
			ExpectedErr: "mdl: Not Loaded",
		},



		{
			Key: mdl.SomeKey("apple", "banana", "cherry"),
			Expected: "apple/banana/cherry",
			ExpectedErr: "",
		},
	}

	for testNumber, test := range tests {

		actual, actualErr := test.Key.Value()

		switch actualErr {
		case nil:
			if expected := test.ExpectedErr; "" != expected {
				t.Errorf("For test #%d, expected an error, but did not actually get one.", testNumber)
				t.Logf("EXPECTED error message: %q", expected)
				t.Logf("ACTUAL   error:         %#v", actualErr)
				continue
			}
		default:
			if expected, actual := test.ExpectedErr, actualErr.Error(); expected != actual {
				t.Errorf("For test #%d, expected an error message, but did not actually get it.", testNumber)
				t.Logf("EXPECTED error message: %q", expected)
				t.Logf("ACTUAL   error message: %q", actual)
				continue

			}
		}

		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, did not actually get what was expected from .Value()", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("KEY:      %#v", test.Key)
			continue
		}
	}
}
