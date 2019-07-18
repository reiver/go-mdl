package mdl_test

import (
	"github.com/reiver/go-mdl"

	"testing"
)

func TestKeyValuesDoubleStoreError(t *testing.T) {

	var keyvalues mdl.KeyValues

	var expectedKey   mdl.Key = mdl.SomeKey("when")
	var expectedValue string  = "now"

	{
		var key   mdl.Key = expectedKey
		var value string  = expectedValue

		if err := keyvalues.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			t.Log("\tKEY-VALUES:...")
			keyvalues.For(func(key mdl.Key, value string){
				t.Logf("\t%#v -> %q", key, value)
			})
			return
		}
	}

	{
		var key   mdl.Key = mdl.SomeKey("n")
		var value string  = "5"

		if err := keyvalues.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", err, err)
			t.Log("\tKEY-VALUES:...")
			keyvalues.For(func(key mdl.Key, value string){
				t.Logf("\t%#v -> %q", key, value)
			})
			return
		}
	}

	{
		var key   mdl.Key = mdl.SomeKey("when")
		var value string  = "later"

		if err := keyvalues.Store(key, value); nil == err {
			t.Errorf("Expected an error, but did not actually get one: #%v", err)
			t.Log("\tKEY-VALUES:...")
			keyvalues.For(func(key mdl.Key, value string){
				t.Logf("\t%#v -> %q", key, value)
			})
			return
		}
	}

	{
		if expected, actual := 2, keyvalues.Len(); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			t.Log("\tKEY-VALUES:...")
			keyvalues.For(func(key mdl.Key, value string){
				t.Logf("%#v -> %q", key, value)
			})
			return
		}

		{
			actualValue := keyvalues.Load(expectedKey)

			if expected, actual := mdl.SomeString(expectedValue), actualValue; expected != actual {
				t.Errorf("Expected value %#v, but actually got value %#v.", expected, actual)
				t.Log("\tKEY-VALUES:...")
				keyvalues.For(func(key mdl.Key, value string){
					t.Logf("%#v -> %q", key, value)
				})
				return
			}
		}
	}

	{
		var key   mdl.Key = mdl.SomeKey("one","two","three")
		var value string  = "l"

		if err := keyvalues.Store(key, value); nil != err {
			t.Errorf("Did not expect an error, but did not actually get one: (%T) #%v", err, err)
			t.Log("\tKEY-VALUES:...")
			keyvalues.For(func(key mdl.Key, value string){
				t.Logf("\t%#v -> %q", key, value)
			})
			return
		}
	}
}
