package mdl

import (
	"net/http"
	"strings"

	"testing"
)

func TestInferData(t *testing.T) {

	tests := []struct{
		HttpRequest *http.Request
		Expected *KeyValues
	}{
		{
			HttpRequest: func()*http.Request{
				var body string = "email_address=joeblow@example.com"

				r, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(body))
				if nil != err {
					panic(err)
				}
				r.Header.Add("X-Idempotent-ID", "z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m")
				r.Header.Add("X-HTTP-Method-Override", "RECORD_EMAIL")
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				return r
			}(),
			Expected: func() *KeyValues {
				var keyvalues KeyValues

				if err := keyvalues.ShallowStore("email_address", "joeblow@example.com"); nil != err {
					panic(err)
				}

				return &keyvalues
			}(),
		},



		{
			HttpRequest: func()*http.Request{
				var body string = "given_name=Joe&family_name=Blow"

				r, err := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(body))
				if nil != err {
					panic(err)
				}
				r.Header.Add("X-Idempotent-ID", "z-2015-05-07T10:27:24Z_cjaLIpvS2MmsLqzPOJIFHwPxqgbdCmi749u9rfAiUu0wZHQ4Z714zSgjumgj")
				r.Header.Add("X-HTTP-Method-Override", "RECORD_EMAIL")
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				return r
			}(),
			Expected: func() *KeyValues {
				var keyvalues KeyValues

				if err := keyvalues.ShallowStore("given_name", "Joe"); nil != err {
					panic(err)
				}
				if err := keyvalues.ShallowStore("family_name", "Blow"); nil != err {
					panic(err)
				}


				return &keyvalues
			}(),
		},
	}

	for testNumber, test := range tests {

		var keyvalues KeyValues

		if err := inferData(&keyvalues, test.HttpRequest); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if  expected, actual := test.Expected.CanonicalForm(), keyvalues.CanonicalForm(); expected != actual {
			t.Errorf("For test #%d, the key-values that were actually gotten were not what was expected.", testNumber)
			t.Logf("EXPECTED:...")
			t.Log(expected)
			t.Logf("ACTUAL:...")
			t.Log(actual)
			continue
		}
	}
}
