package mdl

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

func TestInferVerb(t *testing.T) {

	tests := []struct{
		Request     *http.Request
		ExpectedVerb string
		ExpectedOK   bool
	}{
		{
			Request: nil,
			ExpectedVerb: "",
			ExpectedOK: false,
		},



		{
			Request: httptest.NewRequest("APPLE", "/path/to/script", nil),
			ExpectedVerb: "APPLE",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("BANANA", "/path/to/script", nil),
			ExpectedVerb: "BANANA",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("CHERRY", "/path/to/script", nil),
			ExpectedVerb: "CHERRY",
			ExpectedOK: true,
		},



		{
			Request: httptest.NewRequest("KICK", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "KICK",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("PUNCH", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "PUNCH",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("SCREAM", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "SCREAM",
			ExpectedOK: true,
		},



		{
			Request: httptest.NewRequest("PROVIDE_ADDRESS", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "PROVIDE_ADDRESS",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("PROVIDE_EMAIL_ADDRESS", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "PROVIDE_EMAIL_ADDRESS",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("PROVIDE_NAME", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "PROVIDE_NAME",
			ExpectedOK: true,
		},



		{
			Request: httptest.NewRequest("PATCH", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "PATCH",
			ExpectedOK: true,
		},
		{
			Request: httptest.NewRequest("POST", "/apple/banana/cherry.html", nil),
			ExpectedVerb: "POST",
			ExpectedOK: true,
		},



		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`GET /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "GET",
			ExpectedOK: true,
		},
		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`GET /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "GET",
			ExpectedOK: true,
		},



		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`PATCH /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "PATCH",
			ExpectedOK: true,
		},
		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`PATCH /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "YELL",
			ExpectedOK: true,
		},



		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`POST /path/to/api HTTP/1.1
Host: api.example.com

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "POST",
			ExpectedOK: true,
		},
		{
			Request: func()*http.Request{r,e:=http.ReadRequest(bufio.NewReader(strings.NewReader(
`POST /path/to/api HTTP/1.1
Host: api.example.com
X-HTTP-Method-Override: YELL

apple=one&banana=two&cherry=three`))) ; if nil != e { panic(e) } ; return r}(),
			ExpectedVerb: "YELL",
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualVerb, actualOK := inferVerb(test.Request)

		if expected, actual := test.ExpectedOK, actualOK; expected != actual {
			t.Errorf("For test #%d, did not get expected value for ‘ok’.", testNumber)
			t.Logf("\tEXPECTED ok: %t", expected)
			t.Logf("\tACTUAL   ok: %t", actual)
			continue
		}

		if expected, actual := test.ExpectedVerb, actualVerb; expected != actual {
			t.Errorf("For test #%d, did not get expected value for ‘verb’.", testNumber)
			t.Logf("\tEXPECTED verb: %q", expected)
			t.Logf("\tACTUAL   verb: %q", actual)
			continue
		}
	}
}
