package mdl

import (
	"bufio"
	"net/http"
	"strings"

	"testing"
)

func TestInferID(t *testing.T) {

	tests := []struct{
		Request    string
		ExpectedID string
		ExpectedOK bool
	}{
		{
			Request:
`APPLE /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`BANANA /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`CHERRY /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},



		{
			Request:
`KICK /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`PUNCH /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`SCREAM /path/to/script HTTP/1.1
Host: api.example.com

`,
			ExpectedID: "",
			ExpectedOK: false,
		},



		{
			Request:
`PROVIDE_ADDRESS /apple/banana/cherry.html HTTP/1.1
Host: api.example.com

country_name=Canada&region=BC&locality=Vancouver
`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`PROVIDE_EMAIL_ADDRESS /apple/banana/cherry.html HTTP/1.1
Host: api.example.com

email_address=joeblow@example.com
`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`PROVIDE_NAME /apple/banana/cherry.html HTTP/1.1
Host: api.example.com

given_name=Joe&family_name=Blow
`,
			ExpectedID: "",
			ExpectedOK: false,
		},



		{
			Request:
`PATCH /apple/banana/cherry.html HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three
`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`POST /apple/banana/cherry.html HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three
`,
			ExpectedID: "",
			ExpectedOK: false,
		},



		{
			Request:
`GET /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`GET /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`GET /path/to/api HTTP/1.1
Host: api.example.com
X-Idempotent-ID: abcd-1234

apple=one&banana=two&cherry=three`,
			ExpectedID: "abcd-1234",
			ExpectedOK: true,
		},
		{
			Request:
`GET /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL
X-Idempotent-ID: abc-123

apple=one&banana=two&cherry=three`,
			ExpectedID: "abc-123",
			ExpectedOK: true,
		},



		{
			Request:
`PATCH /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`PATCH /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`PATCH /path/to/api HTTP/1.1
Host: api.example.com
X-Idempotent-ID: abcd-1234

apple=one&banana=two&cherry=three`,
			ExpectedID: "abcd-1234",
			ExpectedOK: true,
		},
		{
			Request:
`PATCH /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL
X-Idempotent-ID: abc-123

apple=one&banana=two&cherry=three`,
			ExpectedID: "abc-123",
			ExpectedOK: true,
		},



		{
			Request:
`POST /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`POST /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`,
			ExpectedID: "",
			ExpectedOK: false,
		},
		{
			Request:
`POST /path/to/api HTTP/1.1
Host: api.example.com
X-Idempotent-ID: abcd-1234

apple=one&banana=two&cherry=three`,
			ExpectedID: "abcd-1234",
			ExpectedOK: true,
		},
		{
			Request:
`POST /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL
X-Idempotent-ID: abc-123

apple=one&banana=two&cherry=three`,
			ExpectedID: "abc-123",
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		var request *http.Request
		{
			r, err := http.ReadRequest(bufio.NewReader(strings.NewReader(test.Request)))
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				t.Log("REQUEST:...")
				t.Log(test.Request)
				continue
			}
			request = r
		}

		actualID, actualOK := inferID(request)

		if expected, actual := test.ExpectedOK, actualOK; expected != actual {
			t.Errorf("For test #%d, did not get expected value for ‘ok’.", testNumber)
			t.Logf("\tEXPECTED ok: %t", expected)
			t.Logf("\tACTUAL   ok: %t", actual)
			t.Log("\tREQUEST:...")
			t.Log(test.Request)
			continue
		}

		if expected, actual := test.ExpectedID, actualID; expected != actual {
			t.Errorf("For test #%d, did not get expected value for ‘verb’.", testNumber)
			t.Logf("\tEXPECTED verb: %q", expected)
			t.Logf("\tACTUAL   verb: %q", actual)
			t.Log("\tREQUEST:...")
			t.Log(test.Request)
			continue
		}
	}
}
